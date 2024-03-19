package components

import (
	"context"
	"fmt"
	"path"
	"strings"

	corev1 "k8s.io/api/core/v1"

	ytv1 "github.com/ytsaurus/yt-k8s-operator/api/v1"
	"github.com/ytsaurus/yt-k8s-operator/pkg/apiproxy"
	"github.com/ytsaurus/yt-k8s-operator/pkg/consts"
	"github.com/ytsaurus/yt-k8s-operator/pkg/labeller"
	"github.com/ytsaurus/yt-k8s-operator/pkg/resources"
	"github.com/ytsaurus/yt-k8s-operator/pkg/ytconfig"
)

type StrawberryController struct {
	localComponent
	cfgen              *ytconfig.Generator
	microservice       microservice
	initUserAndUrlJob  *InitJob
	initChytClusterJob *InitJob
	secret             *resources.StringSecret

	master    Component
	scheduler Component
	dataNodes []Component

	name string
}

func getControllerConfigFileName(name string) string {
	if name == "chyt" {
		return "chyt-controller.yson"
	} else {
		return "strawberry-controller.yson"
	}
}

const ChytInitClusterJobConfigFileName = "chyt-init-cluster.yson"

func NewStrawberryController(
	cfgen *ytconfig.Generator,
	ytsaurus *apiproxy.Ytsaurus,
	master Component,
	scheduler Component,
	dataNodes []Component) *StrawberryController {
	resource := ytsaurus.GetResource()

	image := resource.Spec.CoreImage
	name := "chyt"
	componentName := "ChytController"
	if resource.Spec.DeprecatedChytController != nil {
		if resource.Spec.DeprecatedChytController.Image != nil {
			image = *resource.Spec.DeprecatedChytController.Image
		}
	}
	if resource.Spec.StrawberryController != nil {
		name = "strawberry"
		componentName = "StrawberryController"
		if resource.Spec.StrawberryController.Image != nil {
			image = *resource.Spec.StrawberryController.Image
		}
	}

	l := labeller.Labeller{
		ObjectMeta:     &resource.ObjectMeta,
		APIProxy:       ytsaurus.APIProxy(),
		ComponentLabel: fmt.Sprintf("yt-%s-controller", name),
		ComponentName:  componentName,
		Annotations:    resource.Spec.ExtraPodAnnotations,
	}

	microservice := newMicroservice(
		&l,
		ytsaurus,
		image,
		1,
		map[string]ytconfig.GeneratorDescriptor{
			getControllerConfigFileName(name): {
				F:   cfgen.GetStrawberryControllerConfig,
				Fmt: ytconfig.ConfigFormatYson,
			},
		},
		fmt.Sprintf("%s-controller", name),
		name)

	return &StrawberryController{
		localComponent: newLocalComponent(&l, ytsaurus),
		cfgen:          cfgen,
		microservice:   microservice,
		initUserAndUrlJob: NewInitJob(
			&l,
			ytsaurus.APIProxy(),
			ytsaurus,
			ytsaurus.GetResource().Spec.ImagePullSecrets,
			"user",
			consts.ClientConfigFileName,
			resource.Spec.CoreImage,
			cfgen.GetNativeClientConfig),
		initChytClusterJob: NewInitJob(
			&l,
			ytsaurus.APIProxy(),
			ytsaurus,
			resource.Spec.ImagePullSecrets,
			"cluster",
			ChytInitClusterJobConfigFileName,
			image,
			cfgen.GetChytInitClusterConfig),
		secret: resources.NewStringSecret(
			l.GetSecretName(),
			&l,
			ytsaurus.APIProxy()),
		name:      name,
		master:    master,
		scheduler: scheduler,
		dataNodes: dataNodes,
	}
}

func (c *StrawberryController) IsUpdatable() bool {
	return true
}

func (c *StrawberryController) Fetch(ctx context.Context) error {
	return resources.Fetch(ctx,
		c.microservice,
		c.initUserAndUrlJob,
		c.initChytClusterJob,
		c.secret,
	)
}

func (c *StrawberryController) initUsers() string {
	token, _ := c.secret.GetValue(consts.TokenSecretKey)
	commands := createUserCommand(consts.StrawberryControllerUserName, "", token, true)
	return strings.Join(commands, "\n")
}

func (c *StrawberryController) createInitUserAndUrlScript() string {
	script := []string{
		initJobWithNativeDriverPrologue(),
		c.initUsers(),
		RunIfNonexistent("//sys/@ui_config", "yt set //sys/@ui_config '{}'"),
		fmt.Sprintf("yt set //sys/@ui_config/chyt_controller_base_url '\"http://%v:%v\"'",
			c.microservice.getHttpService().Name(), consts.StrawberryHTTPAPIPort),
	}

	return strings.Join(script, "\n")
}

func (c *StrawberryController) createInitChytClusterScript() string {
	script := []string{
		initJobPrologue,
		fmt.Sprintf("/usr/bin/chyt-controller --config-path %s init-cluster",
			path.Join(consts.ConfigMountPoint, ChytInitClusterJobConfigFileName)),
	}

	return strings.Join(script, "\n")
}

func (c *StrawberryController) getEnvSource() []corev1.EnvFromSource {
	return []corev1.EnvFromSource{
		{
			SecretRef: &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: c.secret.Name(),
				},
			},
		},
	}
}

func (c *StrawberryController) prepareInitChytClusterJob() {
	c.initChytClusterJob.SetInitScript(c.createInitChytClusterScript())

	job := c.initChytClusterJob.Build()
	container := &job.Spec.Template.Spec.Containers[0]
	container.EnvFrom = []corev1.EnvFromSource{c.secret.GetEnvSource()}
}

func (c *StrawberryController) syncComponents(ctx context.Context) (err error) {
	service := c.microservice.buildService()
	service.Spec.Type = "ClusterIP"

	deployment := c.microservice.buildDeployment()
	volumeMounts := []corev1.VolumeMount{
		createConfigVolumeMount(),
	}

	deployment.Spec.Template.Spec.Containers = []corev1.Container{
		{
			Image:   c.microservice.getImage(),
			Name:    consts.UIContainerName,
			EnvFrom: c.getEnvSource(),
			Command: []string{
				"/usr/bin/chyt-controller",
				"--config-path",
				path.Join(consts.ConfigMountPoint, getControllerConfigFileName(c.name)),
				"run",
			},
			VolumeMounts: volumeMounts,
		},
	}

	deployment.Spec.Template.Spec.Volumes = []corev1.Volume{
		createConfigVolume(consts.ConfigVolumeName, c.labeller.GetMainConfigMapName(), nil),
	}

	return c.microservice.Sync(ctx)
}

func (c *StrawberryController) Status(ctx context.Context) ComponentStatus {
	if ytv1.IsReadyToUpdateClusterState(c.ytsaurus.GetClusterState()) && c.microservice.needUpdate() {
		return SimpleStatus(SyncStatusNeedLocalUpdate)
	}

	if c.ytsaurus.GetClusterState() == ytv1.ClusterStateUpdating {
		if IsUpdatingComponent(c.ytsaurus, c) {
			if c.ytsaurus.GetUpdateState() == ytv1.UpdateStateWaitingForPodsRemoval {
				return WaitingStatus(SyncStatusUpdating, "pods removal")
			}

			if c.ytsaurus.GetUpdateState() != ytv1.UpdateStateWaitingForPodsCreation {
				return NewComponentStatus(SyncStatusReady, "Nothing to do now")
			}
		} else {
			return NewComponentStatus(SyncStatusReady, "Not updating component")
		}
	}

	if !IsRunningStatus(c.master.Status(ctx).SyncStatus) {
		return WaitingStatus(SyncStatusBlocked, c.master.GetName())
	}

	if !IsRunningStatus(c.scheduler.Status(ctx).SyncStatus) {
		return WaitingStatus(SyncStatusBlocked, c.scheduler.GetName())
	}

	for _, dataNode := range c.dataNodes {
		if !IsRunningStatus(dataNode.Status(ctx).SyncStatus) {
			return WaitingStatus(SyncStatusBlocked, dataNode.GetName())
		}
	}

	if c.secret.NeedSync(consts.TokenSecretKey, "") {
		return WaitingStatus(SyncStatusPending, c.secret.Name())
	}

	status, err := c.initUserAndUrlJob.Sync(ctx, true)
	if err != nil {
		panic(err)
	}
	if status.SyncStatus != SyncStatusReady {
		return status
	}

	status, err = c.initChytClusterJob.Sync(ctx, true)
	if err != nil {
		panic(err)
	}
	if status.SyncStatus != SyncStatusReady {
		return status
	}

	if c.microservice.needSync() {
		return WaitingStatus(SyncStatusPending, "components")
	}

	return SimpleStatus(SyncStatusReady)
}

func (c *StrawberryController) Sync(ctx context.Context) error {
	if ytv1.IsReadyToUpdateClusterState(c.ytsaurus.GetClusterState()) && c.microservice.needUpdate() {
		return nil
	}

	if c.ytsaurus.GetClusterState() == ytv1.ClusterStateUpdating {
		if IsUpdatingComponent(c.ytsaurus, c) {
			if c.ytsaurus.GetUpdateState() == ytv1.UpdateStateWaitingForPodsRemoval {
				return removePods(ctx, c.microservice, &c.localComponent)
			}

			if c.ytsaurus.GetUpdateState() != ytv1.UpdateStateWaitingForPodsCreation {
				return nil
			}
		} else {
			return nil
		}
	}

	if !IsRunningStatus(c.master.Status(ctx).SyncStatus) {
		return nil
	}

	if !IsRunningStatus(c.scheduler.Status(ctx).SyncStatus) {
		return nil
	}

	for _, dataNode := range c.dataNodes {
		if !IsRunningStatus(dataNode.Status(ctx).SyncStatus) {
			return nil
		}
	}

	if c.secret.NeedSync(consts.TokenSecretKey, "") {
		s := c.secret.Build()
		s.StringData = map[string]string{
			consts.TokenSecretKey: ytconfig.RandString(30),
		}
		return c.secret.Sync(ctx)
	}

	c.initUserAndUrlJob.SetInitScript(c.createInitUserAndUrlScript())
	status, err := c.initUserAndUrlJob.Sync(ctx, false)
	if err != nil || status.SyncStatus != SyncStatusReady {
		return err
	}

	c.prepareInitChytClusterJob()
	status, err = c.initChytClusterJob.Sync(ctx, false)
	if err != nil || status.SyncStatus != SyncStatusReady {
		return err
	}

	if c.microservice.needSync() {
		return c.syncComponents(ctx)
	}

	return nil
}

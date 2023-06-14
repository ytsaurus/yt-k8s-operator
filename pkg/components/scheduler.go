package components

import (
	"context"
	"fmt"
	"strings"

	v1 "github.com/ytsaurus/yt-k8s-operator/api/v1"

	"github.com/ytsaurus/yt-k8s-operator/pkg/apiproxy"
	"github.com/ytsaurus/yt-k8s-operator/pkg/consts"
	"github.com/ytsaurus/yt-k8s-operator/pkg/labeller"
	"github.com/ytsaurus/yt-k8s-operator/pkg/resources"
	"github.com/ytsaurus/yt-k8s-operator/pkg/ytconfig"
	corev1 "k8s.io/api/core/v1"
)

type scheduler struct {
	ServerComponentBase
	master        Component
	execNodes     Component
	tabletNodes   Component
	initUser      *InitJob
	initOpArchive *InitJob
	secret        *resources.StringSecret
}

func NewScheduler(cfgen *ytconfig.Generator, apiProxy *apiproxy.APIProxy, master, execNodes, tabletNodes Component) Component {
	ytsaurus := apiProxy.Ytsaurus()
	labeller := labeller.Labeller{
		Ytsaurus:       ytsaurus,
		APIProxy:       apiProxy,
		ComponentLabel: consts.YTComponentLabelScheduler,
		ComponentName:  "Scheduler",
		MonitoringPort: consts.SchedulerMonitoringPort,
	}

	server := NewServer(
		&labeller,
		apiProxy,
		&ytsaurus.Spec.Schedulers.InstanceSpec,
		"/usr/bin/ytserver-scheduler",
		"ytserver-scheduler.yson",
		cfgen.GetSchedulerStatefulSetName(),
		cfgen.GetSchedulerServiceName(),
		cfgen.GetSchedulerConfig,
	)

	return &scheduler{
		ServerComponentBase: ServerComponentBase{
			ComponentBase: ComponentBase{
				labeller: &labeller,
				apiProxy: apiProxy,
				cfgen:    cfgen,
			},
			server: server,
		},
		master:      master,
		execNodes:   execNodes,
		tabletNodes: tabletNodes,
		initUser: NewInitJob(
			&labeller,
			apiProxy,
			"user",
			consts.ClientConfigFileName,
			cfgen.GetNativeClientConfig),
		initOpArchive: NewInitJob(
			&labeller,
			apiProxy,
			"op-archive",
			consts.ClientConfigFileName,
			cfgen.GetNativeClientConfig),
		secret: resources.NewStringSecret(
			labeller.GetSecretName(),
			&labeller,
			apiProxy),
	}
}

func (s *scheduler) createInitScript() string {
	token, _ := s.secret.GetValue(consts.TokenSecretKey)
	commands := createUserCommand("operation_archivarius", "", token, true)
	script := []string{
		initJobWithNativeDriverPrologue(),
	}
	script = append(script, commands...)

	return strings.Join(script, "\n")
}

func (s *scheduler) prepareInitOperationArchive() {
	script := []string{
		initJobWithNativeDriverPrologue(),
		fmt.Sprintf("/usr/bin/init_operation_archive --force --latest --proxy %s",
			s.cfgen.GetHTTPProxiesServiceAddress(consts.DefaultHTTPProxyRole)),
		"/usr/bin/yt set //sys/cluster_nodes/@config '{\"%true\" = {job_agent={enable_job_reporter=%true}}}'",
	}

	s.initOpArchive.SetInitScript(strings.Join(script, "\n"))
	job := s.initOpArchive.Build()
	container := &job.Spec.Template.Spec.Containers[0]
	container.EnvFrom = []corev1.EnvFromSource{s.secret.GetEnvSource()}
}

func (s *scheduler) Fetch(ctx context.Context) error {
	return resources.Fetch(ctx, []resources.Fetchable{
		s.server,
		s.initOpArchive,
		s.initUser,
		s.secret,
	})
}

func (s *scheduler) doSync(ctx context.Context, dry bool) (SyncStatus, error) {
	var err error

	if s.apiProxy.GetClusterState() == v1.ClusterStateUpdating {
		if s.apiProxy.GetUpdateState() == v1.UpdateStateWaitingForPodsRemoval {
			return SyncStatusUpdating, s.removePods(ctx, dry)
		}
	}

	if s.master.Status(ctx) != SyncStatusReady {
		return SyncStatusBlocked, err
	}

	if s.execNodes == nil || s.execNodes.Status(ctx) != SyncStatusReady {
		// It makes no sense to start scheduler without exec nodes.
		return SyncStatusBlocked, err
	}

	if s.secret.NeedSync(consts.TokenSecretKey, "") {
		if !dry {
			secretSpec := s.secret.Build()
			secretSpec.StringData = map[string]string{
				consts.TokenSecretKey: ytconfig.RandString(30),
			}
			err = s.secret.Sync(ctx)
		}
		return SyncStatusPending, err
	}

	if !s.server.IsInSync() {
		if !dry {
			// TODO(psushin): there should be me more sophisticated logic for version updates.
			err = s.server.Sync(ctx)
		}
		return SyncStatusPending, err
	}

	if !s.server.ArePodsReady(ctx) {
		return SyncStatusBlocked, err
	}

	if s.tabletNodes == nil {
		// Don't initialize operations archive.
		return SyncStatusReady, err
	}

	if !dry {
		s.initUser.SetInitScript(s.createInitScript())
	}

	status, err := s.initUser.Sync(ctx, dry)
	if status != SyncStatusReady {
		return status, err
	}

	if s.tabletNodes.Status(ctx) != SyncStatusReady {
		// Wait for tablet nodes to proceed with operations archive init.
		return SyncStatusBlocked, err
	}

	if !dry {
		s.prepareInitOperationArchive()
	}
	return s.initOpArchive.Sync(ctx, dry)
}

func (s *scheduler) Status(ctx context.Context) SyncStatus {
	status, err := s.doSync(ctx, true)
	if err != nil {
		panic(err)
	}

	return status
}

func (s *scheduler) Sync(ctx context.Context) error {
	_, err := s.doSync(ctx, false)
	return err
}

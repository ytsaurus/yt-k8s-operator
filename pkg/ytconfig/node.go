package ytconfig

import (
	"fmt"
	"math"
	"strings"
	"time"

	"go.ytsaurus.tech/yt/go/yson"

	ptr "k8s.io/utils/pointer"

	"k8s.io/apimachinery/pkg/api/resource"

	ytv1 "github.com/ytsaurus/yt-k8s-operator/api/v1"
	"github.com/ytsaurus/yt-k8s-operator/pkg/consts"
	corev1 "k8s.io/api/core/v1"
)

type NodeFlavor string

const (
	NodeFlavorData   NodeFlavor = "data"
	NodeFlavorExec   NodeFlavor = "exec"
	NodeFlavorTablet NodeFlavor = "tablet"
)

type StoreLocation struct {
	Path                   string `yson:"path"`
	MediumName             string `yson:"medium_name"`
	Quota                  int64  `yson:"quota,omitempty"`
	HighWatermark          int64  `yson:"high_watermark,omitempty"`
	LowWatermark           int64  `yson:"low_watermark,omitempty"`
	DisableWritesWatermark int64  `yson:"disable_writes_watermark,omitempty"`
}

type ResourceLimits struct {
	TotalMemory      int64    `yson:"total_memory,omitempty"`
	TotalCpu         *float32 `yson:"total_cpu,omitempty"`
	NodeDedicatedCpu *float32 `yson:"node_dedicated_cpu,omitempty"`
}

type DiskLocation struct {
	Path string `yson:"path"`
}

type SlotLocation struct {
	Path               string `yson:"path"`
	DiskQuota          *int64 `yson:"disk_quota,omitempty"`
	DiskUsageWatermark int64  `yson:"disk_usage_watermark,omitempty"`
	MediumName         string `yson:"medium_name"`
}

type DataNode struct {
	StoreLocations []StoreLocation `yson:"store_locations"`
	CacheLocations []DiskLocation  `yson:"cache_locations"`
	BlockCache     BlockCache      `yson:"block_cache"`
	BlocksExtCache Cache           `yson:"blocks_ext_cache"`
	ChunkMetaCache Cache           `yson:"chunk_meta_cache"`
	BlockMetaCache Cache           `yson:"block_meta_cache"`
}

type JobEnvironmentType string

const (
	JobEnvironmentTypeSimple JobEnvironmentType = "simple"
	JobEnvironmentTypePorto  JobEnvironmentType = "porto"
	JobEnvironmentTypeCRI    JobEnvironmentType = "cri"
)

type GpuInfoSourceType string

const (
	GpuInfoSourceTypeNvGpuManager GpuInfoSourceType = "nv_gpu_manager"
	GpuInfoSourceTypeNvidiaSmi    GpuInfoSourceType = "nvidia_smi"
)

type CriExecutor struct {
	RetryingChannel

	RuntimeEndpoint string        `yson:"runtime_endpoint,omitempty"`
	ImageEndpoint   string        `yson:"image_endpoint,omitempty"`
	Namespace       string        `yson:"namespace"`
	BaseCgroup      string        `yson:"base_cgroup"`
	RuntimeHandler  string        `yson:"runtime_handler,omitempty"`
	CpuPeriod       yson.Duration `yson:"cpu_period,omitempty"`
}

type CriJobEnvironment struct {
	CriExecutor          *CriExecutor `yson:"cri_executor,omitempty"`
	JobProxyImage        string       `yson:"job_proxy_image,omitempty"`
	JobProxyBindMounts   []BindMount  `yson:"job_proxy_bind_mounts,omitempty"`
	UseJobProxyFromImage *bool        `yson:"use_job_proxy_from_image,omitempty"`
}

type JobEnvironment struct {
	Type     JobEnvironmentType `yson:"type,omitempty"`
	StartUID int                `yson:"start_uid,omitempty"`

	// FIXME(khlebnikov): Add "inline" tag into yson or remove polymorphism in config.
	CriJobEnvironment
}

type SlotManager struct {
	Locations      []SlotLocation `yson:"locations"`
	JobEnvironment JobEnvironment `yson:"job_environment"`

	DoNotSetUserId      *bool `yson:"do_not_set_user_id,omitempty"`
	EnableTmpfs         *bool `yson:"enable_tmpfs,omitempty"`
	DetachedTmpfsUmount *bool `yson:"detached_tmpfs_umount,omitempty"`
}

type JobResourceLimits struct {
	UserSlots *int `yson:"user_slots,omitempty"`
}

type GpuInfoSource struct {
	Type GpuInfoSourceType `yson:"type"`
}

type GpuManager struct {
	GpuInfoSource GpuInfoSource `yson:"gpu_info_source"`
}

type JobController struct {
	ResourceLimitsLegacy *JobResourceLimits `yson:"resource_limits,omitempty"`
	GpuManagerLegacy     *GpuManager        `yson:"gpu_manager,omitempty"`
}

type JobResourceManager struct {
	ResourceLimits JobResourceLimits `yson:"resource_limits"`
}

type JobProxy struct {
	JobProxyLogging Logging `yson:"job_proxy_logging"`
}

type ExecNode struct {
	SlotManager   SlotManager   `yson:"slot_manager"`
	GpuManager    GpuManager    `yson:"gpu_manager"`
	JobController JobController `yson:"job_controller"`
	JobProxy      JobProxy      `yson:"job_proxy"`

	JobProxyLoggingLegacy *Logging `yson:"job_proxy_logging,omitempty"`
	DoNotSetUserIdLegacy  *bool    `yson:"do_not_set_user_id,omitempty"`

	// NOTE: Non-legacy "use_artifact_binds" moved into dynamic config.
	UseArtifactBindsLegacy *bool `yson:"use_artifact_binds,omitempty"`
}

type Cache struct {
	Capacity int64 `yson:"capacity"`
}

type BlockCache struct {
	Compressed   Cache `yson:"compressed_data"`
	Uncompressed Cache `yson:"uncompressed_data"`
}

type TabletNode struct {
	VersionedChunkMetaCache Cache `yson:"versioned_chunk_meta_cache"`
}

type NodeServer struct {
	CommonServer
	Flavors        []NodeFlavor   `yson:"flavors"`
	ResourceLimits ResourceLimits `yson:"resource_limits,omitempty"`
	Tags           []string       `yson:"tags,omitempty"`
	Rack           string         `yson:"rack,omitempty"`
	SkynetHttpPort int32          `yson:"skynet_http_port"`
}

type DataNodeServer struct {
	NodeServer
	DataNode DataNode `yson:"data_node"`
}

type ExecNodeServer struct {
	NodeServer
	JobResourceManager   JobResourceManager `yson:"job_resource_manager"`
	ExecNode             ExecNode           `yson:"exec_node"`
	DataNode             DataNode           `yson:"data_node"`
	TabletNode           TabletNode         `yson:"tablet_node"`
	CachingObjectService Cache              `yson:"caching_object_service"`
}

type TabletNodeServer struct {
	NodeServer
	// TabletNode TabletNode `yson:"tablet_node"`
	CachingObjectService Cache `yson:"caching_object_service"`
}

func findVolumeMountForPath(locationPath string, spec ytv1.InstanceSpec) *corev1.VolumeMount {
	for _, mount := range spec.VolumeMounts {
		if strings.HasPrefix(locationPath, mount.MountPath) {
			return &mount
		}
	}
	return nil
}

func findVolumeClaimTemplate(volumeName string, spec ytv1.InstanceSpec) *ytv1.EmbeddedPersistentVolumeClaim {
	for _, claim := range spec.VolumeClaimTemplates {
		if claim.Name == volumeName {
			return &claim
		}
	}
	return nil
}

func findVolume(volumeName string, spec ytv1.InstanceSpec) *corev1.Volume {
	for _, volume := range spec.Volumes {
		if volume.Name == volumeName {
			return &volume
		}
	}
	return nil
}

func findQuotaForPath(locationPath string, spec ytv1.InstanceSpec) *int64 {
	mount := findVolumeMountForPath(locationPath, spec)
	if mount == nil {
		return nil
	}

	if claim := findVolumeClaimTemplate(mount.Name, spec); claim != nil {
		storage := claim.Spec.Resources.Requests.Storage()
		if storage != nil {
			value := storage.Value()
			return &value
		} else {
			return nil
		}
	}

	if volume := findVolume(mount.Name, spec); volume != nil {
		if volume.EmptyDir != nil && volume.EmptyDir.SizeLimit != nil {
			value := volume.EmptyDir.SizeLimit.Value()
			return &value
		}
	}

	return nil
}

func fillClusterNodeServerCarcass(n *NodeServer, flavor NodeFlavor, spec ytv1.ClusterNodesSpec, is *ytv1.InstanceSpec) {
	switch flavor {
	case NodeFlavorData:
		n.RPCPort = consts.DataNodeRPCPort
		n.SkynetHttpPort = consts.DataNodeSkynetPort
	case NodeFlavorExec:
		n.RPCPort = consts.ExecNodeRPCPort
		n.SkynetHttpPort = consts.ExecNodeSkynetPort
	case NodeFlavorTablet:
		n.RPCPort = consts.TabletNodeRPCPort
		n.SkynetHttpPort = consts.TabletNodeSkynetPort
	}

	n.MonitoringPort = *is.MonitoringPort

	n.Flavors = []NodeFlavor{flavor}
	n.Tags = spec.Tags
	n.Rack = spec.Rack
}

func getDataNodeResourceLimits(spec *ytv1.DataNodesSpec) ResourceLimits {
	var resourceLimits ResourceLimits

	var cpu float32 = 0
	resourceLimits.NodeDedicatedCpu = &cpu
	resourceLimits.TotalCpu = &cpu

	memoryRequest := spec.Resources.Requests.Memory()
	memoryLimit := spec.Resources.Limits.Memory()
	if memoryRequest != nil && !memoryRequest.IsZero() {
		resourceLimits.TotalMemory = memoryRequest.Value()
	} else if memoryLimit != nil && !memoryLimit.IsZero() {
		resourceLimits.TotalMemory = memoryLimit.Value()
	}

	return resourceLimits
}

func getDataNodeLogging(spec *ytv1.DataNodesSpec) Logging {
	return createLogging(
		&spec.InstanceSpec,
		"data-node",
		[]ytv1.TextLoggerSpec{defaultInfoLoggerSpec(), defaultStderrLoggerSpec()})
}

func getDataNodeServerCarcass(spec *ytv1.DataNodesSpec) (DataNodeServer, error) {
	var c DataNodeServer
	fillClusterNodeServerCarcass(&c.NodeServer, NodeFlavorData, spec.ClusterNodesSpec, &spec.InstanceSpec)

	c.ResourceLimits = getDataNodeResourceLimits(spec)

	for _, location := range ytv1.FindAllLocations(spec.Locations, ytv1.LocationTypeChunkStore) {
		quota := findQuotaForPath(location.Path, spec.InstanceSpec)
		storeLocation := StoreLocation{
			MediumName: location.Medium,
			Path:       location.Path,
		}
		if quota != nil {
			storeLocation.Quota = *quota

			// These are just simple heuristics.
			gb := float64(1024 * 1024 * 1024)
			storeLocation.LowWatermark = int64(math.Min(0.1*float64(storeLocation.Quota), float64(5)*gb))
			storeLocation.HighWatermark = storeLocation.LowWatermark / 2
			storeLocation.DisableWritesWatermark = storeLocation.HighWatermark / 2
		}
		c.DataNode.StoreLocations = append(c.DataNode.StoreLocations, storeLocation)
	}

	if len(c.DataNode.StoreLocations) == 0 {
		return c, fmt.Errorf("error creating data node config: no storage locations provided")
	}

	c.Logging = getDataNodeLogging(spec)

	return c, nil
}

func getResourceQuantity(resources *corev1.ResourceRequirements, name corev1.ResourceName) resource.Quantity {
	if request, ok := resources.Requests[name]; ok && !request.IsZero() {
		return request
	}
	if limit, ok := resources.Limits[name]; ok && !limit.IsZero() {
		return limit
	}
	return resource.Quantity{}
}

func getExecNodeResourceLimits(spec *ytv1.ExecNodesSpec) ResourceLimits {
	var resourceLimits ResourceLimits

	nodeMemory := getResourceQuantity(&spec.Resources, corev1.ResourceMemory)
	nodeCpu := getResourceQuantity(&spec.Resources, corev1.ResourceCPU)

	totalMemory := nodeMemory
	totalCpu := nodeCpu

	if spec.JobResources != nil {
		totalMemory.Add(getResourceQuantity(spec.JobResources, corev1.ResourceMemory))
		totalCpu.Add(getResourceQuantity(spec.JobResources, corev1.ResourceCPU))

		resourceLimits.NodeDedicatedCpu = ptr.Float32(float32(nodeCpu.AsApproximateFloat64()))
	} else {
		// TODO(khlebnikov): Add better defaults.
		resourceLimits.NodeDedicatedCpu = ptr.Float32(0)
	}

	resourceLimits.TotalMemory = totalMemory.Value()
	if !totalCpu.IsZero() {
		resourceLimits.TotalCpu = ptr.Float32(float32(totalCpu.AsApproximateFloat64()))
	}

	return resourceLimits
}

func getExecNodeLogging(spec *ytv1.ExecNodesSpec) Logging {
	return createLogging(
		&spec.InstanceSpec,
		"exec-node",
		[]ytv1.TextLoggerSpec{defaultInfoLoggerSpec(), defaultStderrLoggerSpec()})
}

func fillJobEnvironment(execNode *ExecNode, spec *ytv1.ExecNodesSpec, commonSpec *ytv1.CommonSpec) error {
	envSpec := spec.JobEnvironment
	jobEnv := &execNode.SlotManager.JobEnvironment

	jobEnv.StartUID = consts.StartUID

	if envSpec != nil && envSpec.CRI != nil {
		jobEnv.Type = JobEnvironmentTypeCRI

		if jobImage := commonSpec.JobImage; jobImage != nil {
			jobEnv.JobProxyImage = *jobImage
			jobEnv.UseJobProxyFromImage = ptr.Bool(false)
		} else {
			jobEnv.JobProxyImage = ptr.StringDeref(spec.Image, commonSpec.CoreImage)
			jobEnv.UseJobProxyFromImage = ptr.Bool(true)
		}

		endpoint := "unix://" + getContainerdSocketPath(spec)

		jobEnv.CriExecutor = &CriExecutor{
			RuntimeEndpoint: endpoint,
			ImageEndpoint:   endpoint,
			Namespace:       ptr.StringDeref(envSpec.CRI.CRINamespace, consts.CRINamespace),
			BaseCgroup:      ptr.StringDeref(envSpec.CRI.BaseCgroup, consts.CRIBaseCgroup),
		}

		if timeout := envSpec.CRI.APIRetryTimeoutSeconds; timeout != nil {
			jobEnv.CriExecutor.RetryingChannel = RetryingChannel{
				RetryBackoffTime: yson.Duration(time.Second),
				RetryAttempts:    *timeout,
				RetryTimeout:     yson.Duration(time.Duration(*timeout) * time.Second),
			}
		}

		// NOTE: Default was "false", now it's "true" and option was moved into dynamic config.
		execNode.UseArtifactBindsLegacy = ptr.Bool(ptr.BoolDeref(envSpec.UseArtifactBinds, true))
		if !*execNode.UseArtifactBindsLegacy {
			// Bind mount chunk cache into job containers if artifact are passed via symlinks.
			for _, location := range ytv1.FindAllLocations(spec.Locations, ytv1.LocationTypeChunkCache) {
				jobEnv.JobProxyBindMounts = append(jobEnv.JobProxyBindMounts, BindMount{
					InternalPath: location.Path,
					ExternalPath: location.Path,
					ReadOnly:     true,
				})
			}
		}

		// FIXME(khlebnikov): For now running jobs as non-root is more likely broken.
		execNode.SlotManager.DoNotSetUserId = ptr.Bool(ptr.BoolDeref(envSpec.UseArtifactBinds, true))

	} else if commonSpec.UsePorto {
		jobEnv.Type = JobEnvironmentTypePorto
		// TODO(psushin): volume locations, root fs binds, etc.
	} else {
		jobEnv.Type = JobEnvironmentTypeSimple
	}

	return nil
}

func getExecNodeServerCarcass(spec *ytv1.ExecNodesSpec, commonSpec *ytv1.CommonSpec) (ExecNodeServer, error) {
	var c ExecNodeServer
	fillClusterNodeServerCarcass(&c.NodeServer, NodeFlavorExec, spec.ClusterNodesSpec, &spec.InstanceSpec)

	c.ResourceLimits = getExecNodeResourceLimits(spec)

	for _, location := range ytv1.FindAllLocations(spec.Locations, ytv1.LocationTypeChunkCache) {
		c.DataNode.CacheLocations = append(c.DataNode.CacheLocations, DiskLocation{
			Path: location.Path,
		})
	}

	if len(c.DataNode.CacheLocations) == 0 {
		return c, fmt.Errorf("error creating exec node config: no cache locations provided")
	}

	for _, location := range ytv1.FindAllLocations(spec.Locations, ytv1.LocationTypeSlots) {
		slotLocation := SlotLocation{
			Path:       location.Path,
			MediumName: location.Medium,
		}
		quota := findQuotaForPath(location.Path, spec.InstanceSpec)
		if quota != nil {
			slotLocation.DiskQuota = quota

			// These are just simple heuristics.
			gb := float64(1024 * 1024 * 1024)
			slotLocation.DiskUsageWatermark = int64(math.Min(0.1*float64(*quota), float64(10)*gb))
		}
		c.ExecNode.SlotManager.Locations = append(c.ExecNode.SlotManager.Locations, slotLocation)
	}

	if len(c.ExecNode.SlotManager.Locations) == 0 {
		return c, fmt.Errorf("error creating exec node config: no slot locations provided")
	}

	if spec.JobEnvironment != nil && spec.JobEnvironment.UserSlots != nil {
		c.JobResourceManager.ResourceLimits.UserSlots = ptr.Int(*spec.JobEnvironment.UserSlots)
	} else {
		// Dummy heuristic.
		jobCpu := ptr.Float32Deref(c.ResourceLimits.TotalCpu, 0) - ptr.Float32Deref(c.ResourceLimits.NodeDedicatedCpu, 0)
		if jobCpu > 0 {
			c.JobResourceManager.ResourceLimits.UserSlots = ptr.Int(int(5 * jobCpu))
		}
	}

	if err := fillJobEnvironment(&c.ExecNode, spec, commonSpec); err != nil {
		return c, err
	}

	c.ExecNode.GpuManager.GpuInfoSource.Type = GpuInfoSourceTypeNvidiaSmi

	c.Logging = getExecNodeLogging(spec)

	jobProxyLoggingBuilder := newJobProxyLoggingBuilder()
	if spec.JobProxyLoggers != nil && len(spec.JobProxyLoggers) > 0 {
		for _, loggerSpec := range spec.JobProxyLoggers {
			jobProxyLoggingBuilder.addLogger(loggerSpec)
		}
	} else {
		for _, defaultLoggerSpec := range []ytv1.TextLoggerSpec{defaultInfoLoggerSpec(), defaultStderrLoggerSpec()} {
			jobProxyLoggingBuilder.addLogger(defaultLoggerSpec)
		}
	}
	jobProxyLoggingBuilder.logging.FlushPeriod = 3000
	c.ExecNode.JobProxy.JobProxyLogging = jobProxyLoggingBuilder.logging

	// TODO(khlebnikov): Drop legacy fields depending on ytsaurus version.
	c.ExecNode.JobController.ResourceLimitsLegacy = &c.JobResourceManager.ResourceLimits
	c.ExecNode.JobController.GpuManagerLegacy = &c.ExecNode.GpuManager
	c.ExecNode.JobProxyLoggingLegacy = &c.ExecNode.JobProxy.JobProxyLogging
	c.ExecNode.DoNotSetUserIdLegacy = c.ExecNode.SlotManager.DoNotSetUserId

	return c, nil
}

func getTabletNodeLogging(spec *ytv1.TabletNodesSpec) Logging {
	return createLogging(
		&spec.InstanceSpec,
		"tablet-node",
		[]ytv1.TextLoggerSpec{defaultInfoLoggerSpec(), defaultStderrLoggerSpec()})
}

func getTabletNodeServerCarcass(spec *ytv1.TabletNodesSpec) (TabletNodeServer, error) {
	var c TabletNodeServer
	fillClusterNodeServerCarcass(&c.NodeServer, NodeFlavorTablet, spec.ClusterNodesSpec, &spec.InstanceSpec)

	var cpu float32 = 0
	c.ResourceLimits.NodeDedicatedCpu = &cpu
	c.ResourceLimits.TotalCpu = &cpu

	memoryRequest := spec.Resources.Requests.Memory()
	memoryLimit := spec.Resources.Limits.Memory()
	if memoryRequest != nil && !memoryRequest.IsZero() {
		c.ResourceLimits.TotalMemory = memoryRequest.Value()
	} else if memoryLimit != nil && !memoryLimit.IsZero() {
		c.ResourceLimits.TotalMemory = memoryLimit.Value()
	}

	c.Logging = getTabletNodeLogging(spec)

	return c, nil
}

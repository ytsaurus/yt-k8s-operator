//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseLoggerSpec) DeepCopyInto(out *BaseLoggerSpec) {
	*out = *in
	if in.RotationPolicy != nil {
		in, out := &in.RotationPolicy, &out.RotationPolicy
		*out = new(LogRotationPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseLoggerSpec.
func (in *BaseLoggerSpec) DeepCopy() *BaseLoggerSpec {
	if in == nil {
		return nil
	}
	out := new(BaseLoggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapSpec) DeepCopyInto(out *BootstrapSpec) {
	*out = *in
	if in.TabletCellBundles != nil {
		in, out := &in.TabletCellBundles, &out.TabletCellBundles
		*out = new(BundlesBootstrapSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapSpec.
func (in *BootstrapSpec) DeepCopy() *BootstrapSpec {
	if in == nil {
		return nil
	}
	out := new(BootstrapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleBootstrapSpec) DeepCopyInto(out *BundleBootstrapSpec) {
	*out = *in
	if in.SnapshotPrimaryMedium != nil {
		in, out := &in.SnapshotPrimaryMedium, &out.SnapshotPrimaryMedium
		*out = new(string)
		**out = **in
	}
	if in.ChangelogPrimaryMedium != nil {
		in, out := &in.ChangelogPrimaryMedium, &out.ChangelogPrimaryMedium
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleBootstrapSpec.
func (in *BundleBootstrapSpec) DeepCopy() *BundleBootstrapSpec {
	if in == nil {
		return nil
	}
	out := new(BundleBootstrapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundlesBootstrapSpec) DeepCopyInto(out *BundlesBootstrapSpec) {
	*out = *in
	if in.Sys != nil {
		in, out := &in.Sys, &out.Sys
		*out = new(BundleBootstrapSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(BundleBootstrapSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundlesBootstrapSpec.
func (in *BundlesBootstrapSpec) DeepCopy() *BundlesBootstrapSpec {
	if in == nil {
		return nil
	}
	out := new(BundlesBootstrapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CategoriesFilter) DeepCopyInto(out *CategoriesFilter) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CategoriesFilter.
func (in *CategoriesFilter) DeepCopy() *CategoriesFilter {
	if in == nil {
		return nil
	}
	out := new(CategoriesFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chyt) DeepCopyInto(out *Chyt) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chyt.
func (in *Chyt) DeepCopy() *Chyt {
	if in == nil {
		return nil
	}
	out := new(Chyt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Chyt) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChytList) DeepCopyInto(out *ChytList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Chyt, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChytList.
func (in *ChytList) DeepCopy() *ChytList {
	if in == nil {
		return nil
	}
	out := new(ChytList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChytList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChytSpec) DeepCopyInto(out *ChytSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Ytsaurus != nil {
		in, out := &in.Ytsaurus, &out.Ytsaurus
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChytSpec.
func (in *ChytSpec) DeepCopy() *ChytSpec {
	if in == nil {
		return nil
	}
	out := new(ChytSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChytStatus) DeepCopyInto(out *ChytStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChytStatus.
func (in *ChytStatus) DeepCopy() *ChytStatus {
	if in == nil {
		return nil
	}
	out := new(ChytStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNodesSpec) DeepCopyInto(out *ClusterNodesSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNodesSpec.
func (in *ClusterNodesSpec) DeepCopy() *ClusterNodesSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterNodesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.NativeTransport != nil {
		in, out := &in.NativeTransport, &out.NativeTransport
		*out = new(RPCTransportSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraPodAnnotations != nil {
		in, out := &in.ExtraPodAnnotations, &out.ExtraPodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerAgentsSpec) DeepCopyInto(out *ControllerAgentsSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerAgentsSpec.
func (in *ControllerAgentsSpec) DeepCopy() *ControllerAgentsSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerAgentsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNodesSpec) DeepCopyInto(out *DataNodesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
	in.ClusterNodesSpec.DeepCopyInto(&out.ClusterNodesSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNodesSpec.
func (in *DataNodesSpec) DeepCopy() *DataNodesSpec {
	if in == nil {
		return nil
	}
	out := new(DataNodesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeprecatedSpytSpec) DeepCopyInto(out *DeprecatedSpytSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeprecatedSpytSpec.
func (in *DeprecatedSpytSpec) DeepCopy() *DeprecatedSpytSpec {
	if in == nil {
		return nil
	}
	out := new(DeprecatedSpytSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoverySpec) DeepCopyInto(out *DiscoverySpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoverySpec.
func (in *DiscoverySpec) DeepCopy() *DiscoverySpec {
	if in == nil {
		return nil
	}
	out := new(DiscoverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedObjectMetadata) DeepCopyInto(out *EmbeddedObjectMetadata) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedObjectMetadata.
func (in *EmbeddedObjectMetadata) DeepCopy() *EmbeddedObjectMetadata {
	if in == nil {
		return nil
	}
	out := new(EmbeddedObjectMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedPersistentVolumeClaim) DeepCopyInto(out *EmbeddedPersistentVolumeClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.EmbeddedObjectMetadata.DeepCopyInto(&out.EmbeddedObjectMetadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedPersistentVolumeClaim.
func (in *EmbeddedPersistentVolumeClaim) DeepCopy() *EmbeddedPersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(EmbeddedPersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecNodesSpec) DeepCopyInto(out *ExecNodesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
	in.ClusterNodesSpec.DeepCopyInto(&out.ClusterNodesSpec)
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.JobProxyLoggers != nil {
		in, out := &in.JobProxyLoggers, &out.JobProxyLoggers
		*out = make([]TextLoggerSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecNodesSpec.
func (in *ExecNodesSpec) DeepCopy() *ExecNodesSpec {
	if in == nil {
		return nil
	}
	out := new(ExecNodesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProxiesSpec) DeepCopyInto(out *HTTPProxiesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
	if in.HttpNodePort != nil {
		in, out := &in.HttpNodePort, &out.HttpNodePort
		*out = new(int32)
		**out = **in
	}
	if in.HttpsNodePort != nil {
		in, out := &in.HttpsNodePort, &out.HttpsNodePort
		*out = new(int32)
		**out = **in
	}
	in.Transport.DeepCopyInto(&out.Transport)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProxiesSpec.
func (in *HTTPProxiesSpec) DeepCopy() *HTTPProxiesSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPProxiesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPTransportSpec) DeepCopyInto(out *HTTPTransportSpec) {
	*out = *in
	if in.HTTPSSecret != nil {
		in, out := &in.HTTPSSecret, &out.HTTPSSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPTransportSpec.
func (in *HTTPTransportSpec) DeepCopy() *HTTPTransportSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPTransportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.MinReadyInstanceCount != nil {
		in, out := &in.MinReadyInstanceCount, &out.MinReadyInstanceCount
		*out = new(int)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]LocationSpec, len(*in))
		copy(*out, *in)
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]EmbeddedPersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableAntiAffinity != nil {
		in, out := &in.EnableAntiAffinity, &out.EnableAntiAffinity
		*out = new(bool)
		**out = **in
	}
	if in.Loggers != nil {
		in, out := &in.Loggers, &out.Loggers
		*out = make([]TextLoggerSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StructuredLoggers != nil {
		in, out := &in.StructuredLoggers, &out.StructuredLoggers
		*out = make([]StructuredLoggerSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NativeTransport != nil {
		in, out := &in.NativeTransport, &out.NativeTransport
		*out = new(RPCTransportSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationSpec) DeepCopyInto(out *LocationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationSpec.
func (in *LocationSpec) DeepCopy() *LocationSpec {
	if in == nil {
		return nil
	}
	out := new(LocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogRotationPolicy) DeepCopyInto(out *LogRotationPolicy) {
	*out = *in
	if in.RotationPeriodMilliseconds != nil {
		in, out := &in.RotationPeriodMilliseconds, &out.RotationPeriodMilliseconds
		*out = new(int64)
		**out = **in
	}
	if in.MaxSegmentSize != nil {
		in, out := &in.MaxSegmentSize, &out.MaxSegmentSize
		*out = new(int64)
		**out = **in
	}
	if in.MaxTotalSizeToKeep != nil {
		in, out := &in.MaxTotalSizeToKeep, &out.MaxTotalSizeToKeep
		*out = new(int64)
		**out = **in
	}
	if in.MaxSegmentCountToKeep != nil {
		in, out := &in.MaxSegmentCountToKeep, &out.MaxSegmentCountToKeep
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogRotationPolicy.
func (in *LogRotationPolicy) DeepCopy() *LogRotationPolicy {
	if in == nil {
		return nil
	}
	out := new(LogRotationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterConnectionSpec) DeepCopyInto(out *MasterConnectionSpec) {
	*out = *in
	if in.HostAddresses != nil {
		in, out := &in.HostAddresses, &out.HostAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterConnectionSpec.
func (in *MasterConnectionSpec) DeepCopy() *MasterConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(MasterConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MastersSpec) DeepCopyInto(out *MastersSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
	in.MasterConnectionSpec.DeepCopyInto(&out.MasterConnectionSpec)
	if in.MaxSnapshotCountToKeep != nil {
		in, out := &in.MaxSnapshotCountToKeep, &out.MaxSnapshotCountToKeep
		*out = new(int)
		**out = **in
	}
	if in.MaxChangelogCountToKeep != nil {
		in, out := &in.MaxChangelogCountToKeep, &out.MaxChangelogCountToKeep
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MastersSpec.
func (in *MastersSpec) DeepCopy() *MastersSpec {
	if in == nil {
		return nil
	}
	out := new(MastersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthServiceSpec) DeepCopyInto(out *OauthServiceSpec) {
	*out = *in
	in.UserInfo.DeepCopyInto(&out.UserInfo)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthServiceSpec.
func (in *OauthServiceSpec) DeepCopy() *OauthServiceSpec {
	if in == nil {
		return nil
	}
	out := new(OauthServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthUserInfoHandlerSpec) DeepCopyInto(out *OauthUserInfoHandlerSpec) {
	*out = *in
	if in.ErrorField != nil {
		in, out := &in.ErrorField, &out.ErrorField
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthUserInfoHandlerSpec.
func (in *OauthUserInfoHandlerSpec) DeepCopy() *OauthUserInfoHandlerSpec {
	if in == nil {
		return nil
	}
	out := new(OauthUserInfoHandlerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryTrackerSpec) DeepCopyInto(out *QueryTrackerSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryTrackerSpec.
func (in *QueryTrackerSpec) DeepCopy() *QueryTrackerSpec {
	if in == nil {
		return nil
	}
	out := new(QueryTrackerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueAgentSpec) DeepCopyInto(out *QueueAgentSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueAgentSpec.
func (in *QueueAgentSpec) DeepCopy() *QueueAgentSpec {
	if in == nil {
		return nil
	}
	out := new(QueueAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RPCProxiesSpec) DeepCopyInto(out *RPCProxiesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
	if in.ServiceType != nil {
		in, out := &in.ServiceType, &out.ServiceType
		*out = new(corev1.ServiceType)
		**out = **in
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(int32)
		**out = **in
	}
	in.Transport.DeepCopyInto(&out.Transport)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RPCProxiesSpec.
func (in *RPCProxiesSpec) DeepCopy() *RPCProxiesSpec {
	if in == nil {
		return nil
	}
	out := new(RPCProxiesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RPCTransportSpec) DeepCopyInto(out *RPCTransportSpec) {
	*out = *in
	if in.TLSSecret != nil {
		in, out := &in.TLSSecret, &out.TLSSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RPCTransportSpec.
func (in *RPCTransportSpec) DeepCopy() *RPCTransportSpec {
	if in == nil {
		return nil
	}
	out := new(RPCTransportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulersSpec) DeepCopyInto(out *SchedulersSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulersSpec.
func (in *SchedulersSpec) DeepCopy() *SchedulersSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spyt) DeepCopyInto(out *Spyt) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spyt.
func (in *Spyt) DeepCopy() *Spyt {
	if in == nil {
		return nil
	}
	out := new(Spyt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Spyt) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpytList) DeepCopyInto(out *SpytList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Spyt, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpytList.
func (in *SpytList) DeepCopy() *SpytList {
	if in == nil {
		return nil
	}
	out := new(SpytList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpytList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpytSpec) DeepCopyInto(out *SpytSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Ytsaurus != nil {
		in, out := &in.Ytsaurus, &out.Ytsaurus
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpytSpec.
func (in *SpytSpec) DeepCopy() *SpytSpec {
	if in == nil {
		return nil
	}
	out := new(SpytSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpytStatus) DeepCopyInto(out *SpytStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpytStatus.
func (in *SpytStatus) DeepCopy() *SpytStatus {
	if in == nil {
		return nil
	}
	out := new(SpytStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrawberryControllerSpec) DeepCopyInto(out *StrawberryControllerSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrawberryControllerSpec.
func (in *StrawberryControllerSpec) DeepCopy() *StrawberryControllerSpec {
	if in == nil {
		return nil
	}
	out := new(StrawberryControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StructuredLoggerSpec) DeepCopyInto(out *StructuredLoggerSpec) {
	*out = *in
	in.BaseLoggerSpec.DeepCopyInto(&out.BaseLoggerSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StructuredLoggerSpec.
func (in *StructuredLoggerSpec) DeepCopy() *StructuredLoggerSpec {
	if in == nil {
		return nil
	}
	out := new(StructuredLoggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPProxiesSpec) DeepCopyInto(out *TCPProxiesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
	if in.ServiceType != nil {
		in, out := &in.ServiceType, &out.ServiceType
		*out = new(corev1.ServiceType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPProxiesSpec.
func (in *TCPProxiesSpec) DeepCopy() *TCPProxiesSpec {
	if in == nil {
		return nil
	}
	out := new(TCPProxiesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TabletCellBundleInfo) DeepCopyInto(out *TabletCellBundleInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TabletCellBundleInfo.
func (in *TabletCellBundleInfo) DeepCopy() *TabletCellBundleInfo {
	if in == nil {
		return nil
	}
	out := new(TabletCellBundleInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TabletNodesSpec) DeepCopyInto(out *TabletNodesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
	in.ClusterNodesSpec.DeepCopyInto(&out.ClusterNodesSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TabletNodesSpec.
func (in *TabletNodesSpec) DeepCopy() *TabletNodesSpec {
	if in == nil {
		return nil
	}
	out := new(TabletNodesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TextLoggerSpec) DeepCopyInto(out *TextLoggerSpec) {
	*out = *in
	in.BaseLoggerSpec.DeepCopyInto(&out.BaseLoggerSpec)
	if in.CategoriesFilter != nil {
		in, out := &in.CategoriesFilter, &out.CategoriesFilter
		*out = new(CategoriesFilter)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TextLoggerSpec.
func (in *TextLoggerSpec) DeepCopy() *TextLoggerSpec {
	if in == nil {
		return nil
	}
	out := new(TextLoggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UISpec) DeepCopyInto(out *UISpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.HttpNodePort != nil {
		in, out := &in.HttpNodePort, &out.HttpNodePort
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.OdinBaseUrl != nil {
		in, out := &in.OdinBaseUrl, &out.OdinBaseUrl
		*out = new(string)
		**out = **in
	}
	if in.ExtraEnvVariables != nil {
		in, out := &in.ExtraEnvVariables, &out.ExtraEnvVariables
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UISpec.
func (in *UISpec) DeepCopy() *UISpec {
	if in == nil {
		return nil
	}
	out := new(UISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateStatus) DeepCopyInto(out *UpdateStatus) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TabletCellBundles != nil {
		in, out := &in.TabletCellBundles, &out.TabletCellBundles
		*out = make([]TabletCellBundleInfo, len(*in))
		copy(*out, *in)
	}
	if in.MasterMonitoringPaths != nil {
		in, out := &in.MasterMonitoringPaths, &out.MasterMonitoringPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateStatus.
func (in *UpdateStatus) DeepCopy() *UpdateStatus {
	if in == nil {
		return nil
	}
	out := new(UpdateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YQLAgentSpec) DeepCopyInto(out *YQLAgentSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YQLAgentSpec.
func (in *YQLAgentSpec) DeepCopy() *YQLAgentSpec {
	if in == nil {
		return nil
	}
	out := new(YQLAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ytsaurus) DeepCopyInto(out *Ytsaurus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ytsaurus.
func (in *Ytsaurus) DeepCopy() *Ytsaurus {
	if in == nil {
		return nil
	}
	out := new(Ytsaurus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ytsaurus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YtsaurusList) DeepCopyInto(out *YtsaurusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ytsaurus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YtsaurusList.
func (in *YtsaurusList) DeepCopy() *YtsaurusList {
	if in == nil {
		return nil
	}
	out := new(YtsaurusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *YtsaurusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YtsaurusSpec) DeepCopyInto(out *YtsaurusSpec) {
	*out = *in
	in.ConfigurationSpec.DeepCopyInto(&out.ConfigurationSpec)
	if in.AdminCredentials != nil {
		in, out := &in.AdminCredentials, &out.AdminCredentials
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.OauthService != nil {
		in, out := &in.OauthService, &out.OauthService
		*out = new(OauthServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Bootstrap != nil {
		in, out := &in.Bootstrap, &out.Bootstrap
		*out = new(BootstrapSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Discovery.DeepCopyInto(&out.Discovery)
	in.PrimaryMasters.DeepCopyInto(&out.PrimaryMasters)
	if in.SecondaryMasters != nil {
		in, out := &in.SecondaryMasters, &out.SecondaryMasters
		*out = make([]MastersSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPProxies != nil {
		in, out := &in.HTTPProxies, &out.HTTPProxies
		*out = make([]HTTPProxiesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RPCProxies != nil {
		in, out := &in.RPCProxies, &out.RPCProxies
		*out = make([]RPCProxiesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TCPProxies != nil {
		in, out := &in.TCPProxies, &out.TCPProxies
		*out = make([]TCPProxiesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataNodes != nil {
		in, out := &in.DataNodes, &out.DataNodes
		*out = make([]DataNodesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExecNodes != nil {
		in, out := &in.ExecNodes, &out.ExecNodes
		*out = make([]ExecNodesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schedulers != nil {
		in, out := &in.Schedulers, &out.Schedulers
		*out = new(SchedulersSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerAgents != nil {
		in, out := &in.ControllerAgents, &out.ControllerAgents
		*out = new(ControllerAgentsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TabletNodes != nil {
		in, out := &in.TabletNodes, &out.TabletNodes
		*out = make([]TabletNodesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StrawberryController != nil {
		in, out := &in.StrawberryController, &out.StrawberryController
		*out = new(StrawberryControllerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DeprecatedChytController != nil {
		in, out := &in.DeprecatedChytController, &out.DeprecatedChytController
		*out = new(StrawberryControllerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.QueryTrackers != nil {
		in, out := &in.QueryTrackers, &out.QueryTrackers
		*out = new(QueryTrackerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Spyt != nil {
		in, out := &in.Spyt, &out.Spyt
		*out = new(DeprecatedSpytSpec)
		**out = **in
	}
	if in.YQLAgents != nil {
		in, out := &in.YQLAgents, &out.YQLAgents
		*out = new(YQLAgentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.QueueAgents != nil {
		in, out := &in.QueueAgents, &out.QueueAgents
		*out = new(QueueAgentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UI != nil {
		in, out := &in.UI, &out.UI
		*out = new(UISpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YtsaurusSpec.
func (in *YtsaurusSpec) DeepCopy() *YtsaurusSpec {
	if in == nil {
		return nil
	}
	out := new(YtsaurusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YtsaurusStatus) DeepCopyInto(out *YtsaurusStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.UpdateStatus.DeepCopyInto(&out.UpdateStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YtsaurusStatus.
func (in *YtsaurusStatus) DeepCopy() *YtsaurusStatus {
	if in == nil {
		return nil
	}
	out := new(YtsaurusStatus)
	in.DeepCopyInto(out)
	return out
}

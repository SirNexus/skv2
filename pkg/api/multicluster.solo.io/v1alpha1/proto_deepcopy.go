// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1alpha1

import (
	proto "github.com/gogo/protobuf/proto"
)

// DeepCopyInto for the KubernetesCluster.Spec
func (in *KubernetesClusterSpec) DeepCopyInto(out *KubernetesClusterSpec) {
	p := proto.Clone(in).(*KubernetesClusterSpec)
	*out = *p
}

// DeepCopyInto for the KubernetesCluster.Status
func (in *KubernetesClusterStatus) DeepCopyInto(out *KubernetesClusterStatus) {
	p := proto.Clone(in).(*KubernetesClusterStatus)
	*out = *p
}
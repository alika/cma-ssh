/*
Copyright 2018 Samsung SDS.
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	"github.com/samsung-cnct/ims-kaas/pkg/apis/cluster/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const MachineFinalizer = "cnctmachine.cluster.cnct.sds.samsung.com"

// MachineSpec defines the desired state of Machine
type MachineSpec struct {
	// The full, authoritative list of taints to apply to the corresponding
	// Node.
	// +optional
	Taints []corev1.Taint `json:"taints,omitempty"`

	Roles []common.MachineRoles `json:"roles,omitempty"`

	// This field will be set by the actuators and consumed by higher level
	// entities like autoscaler that will be interfacing with cluster-api as
	// generic provider.
	// +optional
	ProviderID *string `json:"providerID,omitempty"`

	// InstanceType references the type of machine to provision in maas based on cpu, gpu, memory tags
	InstanceType string `json:"instanceType,omitempty"`
}

// MachineSshConfigInfo defines the ssh configuration for the physical
// node represented by this Machine
type MachineSshConfigInfo struct {
	Username string `json:"username,omitempty"`

	// maas machine ip
	Host string `json:"host,omitempty"`

	Port uint32 `json:"port,omitempty"`
}

// MachineStatus defines the observed state of Machine
type MachineStatus struct {
	// When was this status last observed
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`

	// Machine status
	Phase common.MachineStatusPhase `json:"phase,omitempty"`

	// Kubernetes version of the node, should be equal to corresponding cluster version
	KubernetesVersion string `json:"kubernetesVersion"`

	// SshConfig used to record ssh configuration of physical machine
	SshConfig MachineSshConfigInfo `json:"sshConfig,omitempty"`

	// SystemId references the maas system id.
	SystemId string `json:"systemId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Machine is the Schema for the machines API
// +k8s:openapi-gen=true
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase",description="machine status"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type CnctMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachineSpec   `json:"spec,omitempty"`
	Status MachineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineList contains a list of Machine
type CnctMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CnctMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CnctMachine{}, &CnctMachineList{})
}

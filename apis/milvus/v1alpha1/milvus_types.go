/*
Copyright 2020 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// MilvusParameters are the configurable fields of a Milvus.
type MilvusParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// MilvusObservation are the observable fields of a Milvus.
type MilvusObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A MilvusSpec defines the desired state of a Milvus.
type MilvusSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MilvusParameters `json:"forProvider"`
}

// A MilvusStatus represents the observed state of a Milvus.
type MilvusStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MilvusObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Milvus is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type Milvus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MilvusSpec   `json:"spec"`
	Status MilvusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MilvusList contains a list of Milvus
type MilvusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Milvus `json:"items"`
}

// Milvus type metadata.
var (
	MilvusKind             = reflect.TypeOf(Milvus{}).Name()
	MilvusGroupKind        = schema.GroupKind{Group: Group, Kind: MilvusKind}.String()
	MilvusKindAPIVersion   = MilvusKind + "." + SchemeGroupVersion.String()
	MilvusGroupVersionKind = SchemeGroupVersion.WithKind(MilvusKind)
)

func init() {
	SchemeBuilder.Register(&Milvus{}, &MilvusList{})
}

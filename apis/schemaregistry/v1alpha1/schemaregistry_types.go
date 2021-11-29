package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// SchemaRegistryParameters are the configurable fields of a SchemaRegistry.
type SchemaRegistryParameters struct {
	Subject     string `json:"subject"`
	Schema      string `json:"schema"`
	SchemaType  string `json:"schemaType"`
	Environment string `json:"environment"`
}

// SchemaRegistryObservation are the observable fields of a SchemaRegistry.
type SchemaRegistryObservation struct {
	Subject     string `json:"subject,omitempty"`
	Schema      string `json:"schema,omitempty"`
	SchemaType  string `json:"schemaType,omitempty"`
	Environment string `json:"environment,omitempty"`
}

// A SchemaRegistrySpec defines the desired state of a SchemaRegistry.
type SchemaRegistrySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SchemaRegistryParameters `json:"forProvider"`
}

// A MyTypeStatus represents the observed state of a SchemaRegistry.
type SchemaRegistryStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SchemaRegistryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A MyType is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type SchemaRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SchemaRegistrySpec   `json:"spec"`
	Status SchemaRegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemaRegistryList contains a list of SchemaRegistry
type SchemaRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SchemaRegistry `json:"items"`
}

// MyType type metadata.
var (
	SchemaRegistryKind             = reflect.TypeOf(SchemaRegistry{}).Name()
	SchemaRegistryGroupKind        = schema.GroupKind{Group: Group, Kind: SchemaRegistryKind}.String()
	SchemaRegistryKindAPIVersion   = SchemaRegistryKind + "." + SchemeGroupVersion.String()
	SchemaRegistryGroupVersionKind = SchemeGroupVersion.WithKind(SchemaRegistryKind)
)

func init() {
	SchemeBuilder.Register(&SchemaRegistry{}, &SchemaRegistryList{})
}

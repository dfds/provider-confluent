package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// SchemaParameters are the configurable fields of a Schema.
type SchemaParameters struct {
	Subject       string `json:"subject"`
	Compatibility string `json:"compatibility"`
	Schema        string `json:"schema"`
	SchemaType    string `json:"schemaType"`
	Environment   string `json:"environment"`
}

// SchemaObservation are the observable fields of a Schema.
type SchemaObservation struct {
	Subject       string `json:"subject,omitempty"`
	Compatibility string `json:"compatibility,omitempty"`
	Schema        string `json:"schema,omitempty"`
	SchemaType    string `json:"schemaType,omitempty"`
	Environment   string `json:"environment,omitempty"`
}

// SchemaSpec defines the desired state of a Schema.
type SchemaSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SchemaParameters `json:"forProvider"`
}

// SchemaStatus represents the observed state of a Schema.
type SchemaStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schema is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type Schema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchemaSpec   `json:"spec"`
	Status            SchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemaList contains a list of Schema
type SchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schema `json:"items"`
}

// Schema type metadata.
var (
	SchemaKind             = reflect.TypeOf(Schema{}).Name()
	SchemaGroupKind        = schema.GroupKind{Group: Group, Kind: SchemaKind}.String()
	SchemaKindAPIVersion   = SchemaKind + "." + SchemeGroupVersion.String()
	SchemaGroupVersionKind = SchemeGroupVersion.WithKind(SchemaKind)
)

func init() {
	SchemeBuilder.Register(&Schema{}, &SchemaList{})
}

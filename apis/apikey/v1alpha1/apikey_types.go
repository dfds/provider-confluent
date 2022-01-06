package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ApiKeyParameters are the configurable fields of a ApiKey.
type ApiKeyParameters struct {
	Resource       string `json:"resource"`
	ServiceAccount string `json:"serviceAccount"`
	Environment    string `json:"environment"`
	Description    string `json:"description"`
}

// ApiKeyObservation are the observable fields of a ApiKey.
type ApiKeyObservation struct {
	Key string `json:"key"`
}

// ApiKey Spec defines the desired state of a ApiKey.
type ApiKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiKeyParameters `json:"forProvider"`
}

// ApiKey Status represents the observed state of a ApiKey.
type ApiKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schema is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type ApiKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiKeySpec   `json:"spec"`
	Status            ApiKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiKey List contains a list of ApiKey
type ApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiKey `json:"items"`
}

// ApiKey type metadata.
var (
	ApiKeyKind             = reflect.TypeOf(ApiKey{}).Name()
	ApiKeyGroupKind        = schema.GroupKind{Group: Group, Kind: ApiKeyKind}.String()
	ApiKeyKindAPIVersion   = ApiKeyKind + "." + SchemeGroupVersion.String()
	ApiKeyGroupVersionKind = SchemeGroupVersion.WithKind(ApiKeyKind)
)

func init() {
	SchemeBuilder.Register(&ApiKey{}, &ApiKeyList{})
}

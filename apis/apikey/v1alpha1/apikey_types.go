package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// APIKeyParameters are the configurable fields of a APIKey.
type APIKeyParameters struct {
	Resource       string `json:"resource"`
	ServiceAccount string `json:"serviceAccount"`
	Environment    string `json:"environment"`
	Description    string `json:"description"`
}

// APIKeyObservation are the observable fields of a APIKey.
type APIKeyObservation struct {
	Key            string `json:"key"`
	Resource       string `json:"resource"`
	ServiceAccount string `json:"serviceAccount"`
	Environment    string `json:"environment"`
}

// APIKeySpec defines the desired state of a APIKey.
type APIKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       APIKeyParameters `json:"forProvider"`
}

// APIKeyStatus Status represents the observed state of a APIKey.
type APIKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          APIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIKey is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type APIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIKeySpec   `json:"spec"`
	Status            APIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIKeyList contains a list of APIKey
type APIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIKey `json:"items"`
}

// APIKey type metadata.
var (
	APIKeyKind             = reflect.TypeOf(APIKey{}).Name()
	APIKeyGroupKind        = schema.GroupKind{Group: Group, Kind: APIKeyKind}.String()
	APIKeyKindAPIVersion   = APIKeyKind + "." + SchemeGroupVersion.String()
	APIKeyGroupVersionKind = SchemeGroupVersion.WithKind(APIKeyKind)
)

func init() {
	SchemeBuilder.Register(&APIKey{}, &APIKeyList{})
}

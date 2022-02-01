package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ServiceAccountParameters are the configurable fields of a ServiceAccount.
type ServiceAccountParameters struct {
	Description string `json:"description"`
}

// ServiceAccountObservation are the observable fields of a ServiceAccount.
type ServiceAccountObservation struct {
	ID string `json:"id,omitempty"`
}

// ServiceAccountSpec defines the desired state of a ServiceAccount.
type ServiceAccountSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServiceAccountParameters `json:"forProvider"`
}

// ServiceAccountStatus represents the observed state of a ServiceAccount.
type ServiceAccountStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServiceAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccount is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type ServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountSpec   `json:"spec"`
	Status            ServiceAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountList contains a list of ServiceAccount
type ServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccount `json:"items"`
}

// ServiceAccount type metadata.
var (
	ServiceAccountKind             = reflect.TypeOf(ServiceAccount{}).Name()
	ServiceAccountGroupKind        = schema.GroupKind{Group: Group, Kind: ServiceAccountKind}.String()
	ServiceAccountKindAPIVersion   = ServiceAccountKind + "." + SchemeGroupVersion.String()
	ServiceAccountGroupVersionKind = SchemeGroupVersion.WithKind(ServiceAccountKind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccount{}, &ServiceAccountList{})
}

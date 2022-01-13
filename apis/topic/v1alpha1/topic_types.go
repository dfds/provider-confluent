package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// TopicConfig
type Config struct {
	Retention int64 `json:"retention"`
}

// TopicRule
type TopicConfig struct {
	Name       string `json:"name"`
	Partitions int    `json:"partitions"`
	Config     Config `json:"config"`
}

// TopicParameters are the configurable fields of a Topic.
type TopicParameters struct {
	Topic       TopicConfig `json:"Topic"`
	Environment string      `json:"environment"`
	Cluster     string      `json:"cluster"`
}

// TopicObservation are the observable fields of a Topic.
type TopicObservation struct {
	Environment string `json:"environment"`
	Cluster     string `json:"cluster"`
	Name        string `json:"name"`
}

// Topic Spec defines the desired state of a Topic.
type TopicSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TopicParameters `json:"forProvider"`
}

// Topic Status represents the observed state of a Topic.
type TopicStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schema is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicSpec   `json:"spec"`
	Status            TopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Topic List contains a list of Topic
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

// Topic type metadata.
var (
	TopicKind             = reflect.TypeOf(Topic{}).Name()
	TopicGroupKind        = schema.GroupKind{Group: Group, Kind: TopicKind}.String()
	TopicKindAPIVersion   = TopicKind + "." + SchemeGroupVersion.String()
	TopicGroupVersionKind = SchemeGroupVersion.WithKind(TopicKind)
)

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}

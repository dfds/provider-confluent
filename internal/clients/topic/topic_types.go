package topic

import (
	"github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
)

// IClient interface for service account client
type IClient interface {
	TopicCreate(tp v1alpha1.TopicParameters) (interface{}, error)
	TopicDelete(tp v1alpha1.TopicParameters) error
	TopicDescribe(to v1alpha1.TopicObservation) (interface{}, error)
	TopicUpdate(tp v1alpha1.TopicParameters) error
}

// Config is a configuration element for the service account client
type Config struct {
	APICredentials clients.APICredentials
}

// Client is a struct for service account client
type Client struct {
	Config Config
}

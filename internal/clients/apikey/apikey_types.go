package apikey

import "github.com/dfds/provider-confluent/internal/clients"

// IClient interface for service account client
type IClient interface {
	APIKeyCreate(resource string, description string, serviceAccount string, environment string) (APIKey, error)
	APIKeyDelete(key string) error
	// APIKeyList() (APIKeyList, error)
	GetAPIKeyByKey(key string) (APIKeyMetadata, error)
	APIKeyUpdate(key string, description string) error
}

// Config is a configuration element for the service account client
type Config struct {
	APICredentials clients.APICredentials
}

// Client is a struct for service account client
type Client struct {
	Config Config
}

type APIKey struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type APIKeyMetadata struct {
	Created         string `json:"created"`
	Description     string `json:"description"`
	Key             string `json:"key"`
	OwnerEmail      string `json:"owner_email"`
	OwnerResourceId string `json:"owner_resource_id"`
	ResourceId      string `json:"resource_id"`
	ResourceType    string `json:"resource_type"`
}

type APIKeyList []APIKeyMetadata

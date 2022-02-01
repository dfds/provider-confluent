package apikey

import "github.com/dfds/provider-confluent/internal/clients"

// IClient interface for service account client
type IClient interface {
	APIKeyCreate(resource string, description string, serviceAccount string, environment string) (APIKey, error)
	APIKeyDelete(key string) error
	GetAPIKeyByKey(key string) (Metadata, error)
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

// APIKey response from create method
type APIKey struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// Metadata response from get/list method
type Metadata struct {
	Created         string `json:"created"`
	Description     string `json:"description"`
	Key             string `json:"key"`
	OwnerEmail      string `json:"owner_email"`
	OwnerResourceID string `json:"owner_resource_id"`
	ResourceID      string `json:"resource_id"`
	ResourceType    string `json:"resource_type"`
}

type APIKeyList []Metadata

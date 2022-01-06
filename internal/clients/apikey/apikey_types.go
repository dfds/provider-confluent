package apikey

import "github.com/dfds/provider-confluent/internal/clients"

// IClient interface for service account client
type IClient interface {
	ApiKeyCreate(resource string, description string, serviceAccount string, environment string) (ApiKey, error)
	ApiKeyDelete(key string) error
	// ApiKeyList() (ApiKeyList, error)
	GetApiKeyByKey(key string) (ApiKeyMetadata, error)
	ApiKeyUpdate(key string, description string) error
}

// Config is a configuration element for the service account client
type Config struct {
	APICredentials clients.APICredentials
}

// Client is a struct for service account client
type Client struct {
	Config Config
}

type ApiKey struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type ApiKeyMetadata struct {
	Created         string `json:"created"`
	Description     string `json:"description"`
	Key             string `json:"key"`
	OwnerEmail      string `json:"owner_email"`
	OwnerResourceId string `json:"owner_resource_id"`
	ResourceId      string `json:"resource_id"`
	ResourceType    string `json:"resource_type"`
}

type ApiKeyList []ApiKeyMetadata

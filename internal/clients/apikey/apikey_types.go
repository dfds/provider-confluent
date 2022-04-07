package apikey

// IClient interface for service account client
type IClient interface {
	APIKeyCreate(resource string, description string, serviceAccount string, environment string) (APIKey, error)
	APIKeyDelete(key string) error
	GetAPIKeyByKey(key string) (Metadata, error)
	APIKeyUpdate(key string, description string) error
}

// Client is a struct for service account client
type Client struct {
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

// List response from list method
type List []Metadata

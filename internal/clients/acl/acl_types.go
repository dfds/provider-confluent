package acl

import "github.com/dfds/provider-confluent/internal/clients"

// IClient interface for service account client
type IClient interface {
	ACLCreate() ACLBlockList
	ACLDelete() error
	ACLList() (ACLBlockList, error)
	ACLUpdate() error
}

// Config is a configuration element for the service account client
type Config struct {
	APICredentials clients.APICredentials
}

// Client is a struct for service account client
type Client struct {
	Config Config
}

type ACLBlock struct {
	Operation    string `json:"operation"`
	PatternType  string `json:"pattern_type"`
	Permission   string `json:"permission"`
	Principal    string `json:"principal"`
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

type ACLBlockList []ACLBlock

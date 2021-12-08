package schemaregistry

import "github.com/dfds/provider-confluent/internal/clients"

// IClient interface for schemaregistry client
type IClient interface {
	SchemaCreate(subject string, schema string, schemaType string, environment string) (string, error)
	SchemaDelete(subject string, version string, permanent bool, environment string) (string, error)
	SchemaDescribe(subject string, version string, environment string) (SchemaDescribeResponse, error)
}

// Config is a configuration element for the schema registry client
type Config struct {
	APICredentials clients.APICredentials
	SchemaPath     string
}

// Client is a struct for schemaregistry client
type Client struct {
	Config Config
}

// SchemaDescribeResponse is a struct for a response from the schemaregistry in confluent cloud
type SchemaDescribeResponse struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Doc       string `json:"doc"`
	Fields    []struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Doc  string `json:"doc"`
	} `json:"fields"`
}

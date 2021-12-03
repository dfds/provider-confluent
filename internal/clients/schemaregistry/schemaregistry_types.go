package schemaregistry

import "github.com/dfds/provider-confluent/internal/clients"

type Client interface {
	SchemaCreate(subject string, schema string, schemaType string, environment string) (string, error)
	SchemaDelete(subject string, version string, permanent bool, environment string) (string, error)
	SchemaDescribe(subject string, version string, environment string) (SchemaDescribeResponse, error)
}

type SchemaRegistryConfig struct {
	ApiConfig  clients.ApiConfig
	SchemaPath string
}

type SchemaRegistryClient struct {
	Config SchemaRegistryConfig
}

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

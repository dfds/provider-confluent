package v1alpha1

type SchemaRegistry struct {
	Subject     string `json:"subject"`
	Schema      string `json:"schema"`
	SchemaType  string `json:"schemaType"`
	Environment string `json:"environment"`
}

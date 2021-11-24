package schemaregistry

type Client interface {
	Create(subject string, schema string, schemaType string, environment string)
	Delete(subject string, version string, permanent bool)
	Describe(subject string, version string)
}

func NewClient() Client {
	return &ConfluentCLIClient{}
}

type ConfluentCLIClient struct {
}

func (c *ConfluentCLIClient) Create(subject string, schema string, schemaType string, environment string) {
	//TODO: Implement
}

func (c *ConfluentCLIClient) Delete(subject string, version string, permanent bool) {
	//TODO: Implement
}

func (c *ConfluentCLIClient) Describe(subject string, version string) {
	//TODO: Implement
}

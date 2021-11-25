package schemaregistry

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

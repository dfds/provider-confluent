package clients

// APICredentials is a configuration element for all clients who need API keys to access confluent cloud
type APICredentials struct {
	Identifier string `json:"identifier"`
	Key        string `json:"key"`
	Secret     string `json:"secret"`
}

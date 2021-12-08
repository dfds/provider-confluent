package clients

// APIConfig is a configuration element for all clients who need API keys to access confluent cloud
type APIConfig struct {
	Name      string `json:"name"`
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

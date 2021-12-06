package clients

// APIConfig is a configuration element for all clients who need API keys to access confluent cloud
type APIConfig struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

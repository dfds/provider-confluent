package serviceaccount

// IClient interface for service account client
type IClient interface {
	ServiceAccountCreate(name string, description string) (ServiceAccount, error)
	ServiceAccountDelete(id string) error
	ServiceAccountList() ([]ServiceAccount, error)
	ServiceAccountByName(name string) (ServiceAccount, error)
	ServiceAccountByID(id string) (ServiceAccount, error)
	ServiceAccountUpdate(id string, description string) error
}

// Client is a struct for service account client
type Client struct {
}

// ServiceAccount struct for deserialising Confluent Cloud response
type ServiceAccount struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

// List type for deserialising Confluent Cloud list response
type List []ServiceAccount

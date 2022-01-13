package acl

import (
	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
)

// IClient interface for service account client
type IClient interface {
	ACLCreate(aclP v1alpha1.ACLParameters) ([]v1alpha1.ACLRule, error)
	ACLDelete(aclP v1alpha1.ACLParameters) error
	ACLList(serviceAccount string, environment string, cluster string) ([]v1alpha1.ACLRule, error)
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

func FromACLBlockToACLRule(input ACLBlock) v1alpha1.ACLRule {
	return v1alpha1.ACLRule{
		Operation:    input.Operation,
		PatternType:  input.PatternType,
		Permission:   input.Permission,
		Principal:    input.Principal,
		ResourceName: input.ResourceName,
		ResourceType: input.ResourceType,
	}
}

func FromACLRuleToACLBlock(input v1alpha1.ACLRule) ACLBlock {
	return ACLBlock{
		Operation:    input.Operation,
		PatternType:  input.PatternType,
		Permission:   input.Permission,
		Principal:    input.Principal,
		ResourceName: input.ResourceName,
		ResourceType: input.ResourceType,
	}
}

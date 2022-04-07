package acl

import (
	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"
)

// IClient interface for service account client
type IClient interface {
	ACLCreate(aclP v1alpha1.ACLParameters) ([]v1alpha1.ACLRule, error)
	ACLDelete(aclP v1alpha1.ACLParameters) error
	ACLList(serviceAccount string, environment string, cluster string) ([]v1alpha1.ACLRule, error)
}

// Client is a struct for service account client
type Client struct {
}

// Block response object
type Block struct {
	Operation    string `json:"operation"`
	PatternType  string `json:"pattern_type"`
	Permission   string `json:"permission"`
	Principal    string `json:"principal"`
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

// FromACLBlockToACLRule converter from Block to ACLRule
func FromACLBlockToACLRule(input Block) v1alpha1.ACLRule {
	return v1alpha1.ACLRule{
		Operation:    input.Operation,
		PatternType:  input.PatternType,
		Permission:   input.Permission,
		Principal:    input.Principal,
		ResourceName: input.ResourceName,
		ResourceType: input.ResourceType,
	}
}

// FromACLRuleToACLBlock converter from ACLRule to Block
func FromACLRuleToACLBlock(input v1alpha1.ACLRule) Block {
	return Block{
		Operation:    input.Operation,
		PatternType:  input.PatternType,
		Permission:   input.Permission,
		Principal:    input.Principal,
		ResourceName: input.ResourceName,
		ResourceType: input.ResourceType,
	}
}

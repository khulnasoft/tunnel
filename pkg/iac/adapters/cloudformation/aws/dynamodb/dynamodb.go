package dynamodb

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/dynamodb"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a dynamodb instance
func Adapt(cfFile parser.FileContext) dynamodb.DynamoDB {
	return dynamodb.DynamoDB{
		DAXClusters: getClusters(cfFile),
		Tables:      getTables(cfFile),
	}
}

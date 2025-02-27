package documentdb

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/documentdb"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adaps a documentDB instance
func Adapt(cfFile parser.FileContext) documentdb.DocumentDB {
	return documentdb.DocumentDB{
		Clusters: getClusters(cfFile),
	}
}

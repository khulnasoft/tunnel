package config

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/config"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a configurationaggregator instance
func Adapt(cfFile parser.FileContext) config.Config {
	return config.Config{
		ConfigurationAggregrator: getConfigurationAggregator(cfFile),
	}
}

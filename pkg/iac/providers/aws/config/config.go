package config

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type Config struct {
	ConfigurationAggregrator ConfigurationAggregrator
}

type ConfigurationAggregrator struct {
	Metadata         iacTypes.Metadata
	SourceAllRegions iacTypes.BoolValue
}

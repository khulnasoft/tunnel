package mq

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type MQ struct {
	Brokers []Broker
}

type Broker struct {
	Metadata     iacTypes.Metadata
	PublicAccess iacTypes.BoolValue
	Logging      Logging
}

type Logging struct {
	Metadata iacTypes.Metadata
	General  iacTypes.BoolValue
	Audit    iacTypes.BoolValue
}

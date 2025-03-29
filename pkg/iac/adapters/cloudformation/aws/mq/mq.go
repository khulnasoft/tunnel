package mq

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/mq"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an MQ instance
func Adapt(cfFile parser.FileContext) mq.MQ {
	return mq.MQ{
		Brokers: getBrokers(cfFile),
	}
}

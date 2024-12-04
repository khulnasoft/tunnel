package sqs

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/iam"
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type SQS struct {
	Queues []Queue
}

type Queue struct {
	Metadata   iacTypes.Metadata
	QueueURL   iacTypes.StringValue
	Encryption Encryption
	Policies   []iam.Policy
}

type Encryption struct {
	Metadata          iacTypes.Metadata
	KMSKeyID          iacTypes.StringValue
	ManagedEncryption iacTypes.BoolValue
}

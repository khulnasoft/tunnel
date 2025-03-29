package lambda

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type Lambda struct {
	Functions []Function
}

type Function struct {
	Metadata    iacTypes.Metadata
	Tracing     Tracing
	Permissions []Permission
}

const (
	TracingModePassThrough = "PassThrough"
	TracingModeActive      = "Active"
)

type Tracing struct {
	Metadata iacTypes.Metadata
	Mode     iacTypes.StringValue
}

type Permission struct {
	Metadata  iacTypes.Metadata
	Principal iacTypes.StringValue
	SourceARN iacTypes.StringValue
}

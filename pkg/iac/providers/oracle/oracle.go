package oracle

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type Oracle struct {
	Compute Compute
}

type Compute struct {
	AddressReservations []AddressReservation
}

type AddressReservation struct {
	Metadata iacTypes.Metadata
	Pool     iacTypes.StringValue // e.g. public-pool
}

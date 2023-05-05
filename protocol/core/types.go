package core

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type OperationContents interface {
	OperationKind() string
}

type OperationContentsAndResult interface {
	OperationContentsAndResult()
	OperationContents
}

type InternalOperationResult interface {
	OperationContents
	InternalOperationResult()
}

type BalanceUpdateKind interface {
	BalanceUpdateKind() string
}

type Bytes struct {
	Bytes []byte `tz:"dyn"`
}

type BallotKind uint8

const (
	BallotYay BallotKind = iota
	BallotNay
	BallotPass
)

type ContractID interface {
	tz.Base58Encoder
	ContractID()
}

type OriginatedContract struct {
	*tz.ContractHash
	Padding uint8
}

type ImplicitContract struct {
	tz.PublicKeyHash
}

func (*ImplicitContract) ContractID() {}

type OriginatedContractID interface {
	tz.Base58Encoder
	OriginatedContractID()
}

func (*OriginatedContract) ContractID()           {}
func (*OriginatedContract) OriginatedContractID() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ContractID]{
		Variants: encoding.Variants[ContractID]{
			0: (*ImplicitContract)(nil),
			1: (*OriginatedContract)(nil),
		},
	})
	encoding.RegisterEnum(&encoding.Enum[OriginatedContractID]{
		Variants: encoding.Variants[OriginatedContractID]{
			1: (*OriginatedContract)(nil),
		},
	})
}

type BalanceUpdate interface {
	BalanceUpdate()
}

type ManagerMetadata[T OperationResult, B BalanceUpdate] struct {
	BalanceUpdates           []B `tz:"dyn"`
	OperationResult          T
	InternalOperationResults []InternalOperationResult `tz:"dyn"`
}
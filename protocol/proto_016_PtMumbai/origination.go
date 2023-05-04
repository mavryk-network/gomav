package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
)

type Origination = proto_012_Psithaca.Origination
type Script = proto_012_Psithaca.Script

type OriginationResult interface {
	OriginationResult()
	core.OperationResult
}

type OriginationResultContents[T core.BalanceUpdate] struct {
	BalanceUpdates      []T                         `tz:"dyn"`
	OriginatedContracts []core.OriginatedContractID `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	StorageSize         tz.BigInt
	PaidStorageSizeDiff tz.BigInt
	LazyStorageDiff     tz.Option[LazyStorageDiff]
}

func (OriginationResultContents[T]) SuccessfulManagerOperationResult() {}
func (OriginationResultContents[T]) OperationKind() string             { return "origination" }

type OriginationResultApplied struct {
	core.OperationResultApplied[OriginationResultContents[*BalanceUpdate]]
}

func (*OriginationResultApplied) OriginationResult() {}

type OriginationResultBacktracked struct {
	core.OperationResultBacktracked[OriginationResultContents[*BalanceUpdate]]
}

func (*OriginationResultBacktracked) OriginationResult() {}

type OriginationResultFailed struct{ core.OperationResultFailed }

func (*OriginationResultFailed) OriginationResult() {}

type OriginationResultSkipped struct{ core.OperationResultSkipped }

func (*OriginationResultSkipped) OriginationResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OriginationResult]{
		Variants: encoding.Variants[OriginationResult]{
			0: (*OriginationResultApplied)(nil),
			1: (*OriginationResultFailed)(nil),
			2: (*OriginationResultSkipped)(nil),
			3: (*OriginationResultBacktracked)(nil),
		},
	})
}

type OriginationContentsAndResult struct {
	Origination
	Metadata ManagerMetadata[OriginationResult, *BalanceUpdate]
}

func (*OriginationContentsAndResult) OperationContentsAndResult() {}

type OriginationInternalOperationResult struct {
	Source   TransactionDestination
	Nonce    uint16
	Balance  tz.BigUint
	Delegate tz.Option[tz.PublicKeyHash]
	Script   Script
	Result   OriginationResult
}

func (*OriginationInternalOperationResult) InternalOperationResult() {}
func (*OriginationInternalOperationResult) OperationKind() string    { return "origination" }

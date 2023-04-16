package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/proto_005_PsBABY5H"
	kathma "github.com/ecadlabs/gotez/protocol/proto_014_PtKathma"
)

type Origination = proto_005_PsBABY5H.Origination
type Script = proto_005_PsBABY5H.Script

type OriginationResult interface {
	OriginationResult()
	OperationResult
}

type OriginationResultContents struct {
	BalanceUpdates      []*BalanceUpdate          `tz:"dyn"`
	OriginatedContracts []tz.OriginatedContractID `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	StorageSize         tz.BigInt
	PaidStorageSizeDiff tz.BigInt
	LazyStorageDiff     tz.Option[LazyStorageDiff]
}

type OriginationResultApplied struct {
	kathma.OperationResultApplied[OriginationResultContents]
}

func (*OriginationResultApplied) OriginationResult() {}

type OriginationResultBacktracked struct {
	kathma.OperationResultBacktracked[OriginationResultContents]
}

func (*OriginationResultBacktracked) OriginationResult() {}

type OriginationResultFailed struct{ kathma.OperationResultFailed }

func (*OriginationResultFailed) OriginationResult() {}

type OriginationResultSkipped struct{ kathma.OperationResultSkipped }

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
	Metadata MetadataWithResult[OriginationResult]
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

type OriginationSuccessfulManagerOperationResult OriginationResultContents

func (*OriginationSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*OriginationSuccessfulManagerOperationResult) OperationKind() string             { return "origination" }

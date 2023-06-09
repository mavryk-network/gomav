package proto_013_PtJakart

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca/lazy"
)

type Origination = proto_012_Psithaca.Origination
type Script = proto_012_Psithaca.Script

type OriginationResult interface {
	proto_012_Psithaca.OriginationResult
}

type OriginationResultContents struct {
	BalanceUpdates
	OriginatedContracts []core.OriginatedContractID `tz:"dyn" json:"originated_contracts"`
	ConsumedGas         tz.BigUint                  `json:"consumed_gas"`
	ConsumedMilligas    tz.BigUint                  `json:"consumed_milligas"`
	StorageSize         tz.BigInt                   `json:"storage_size"`
	PaidStorageSizeDiff tz.BigInt                   `json:"paid_storage_size_diff"`
	LazyStorageDiff     tz.Option[lazy.StorageDiff] `json:"lazy_storage_diff"`
}

//json:kind=OperationKind()
type OriginationSuccessfulManagerResult struct{ OriginationResultApplied }

func (*OriginationSuccessfulManagerResult) OperationKind() string { return "origination" }

type OriginationResultApplied struct {
	core.OperationResultApplied[OriginationResultContents]
}

func (*OriginationResultApplied) OriginationResult() {}

type OriginationResultBacktracked struct {
	core.OperationResultBacktracked[OriginationResultContents]
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
	Metadata ManagerMetadata[OriginationResult] `json:"metadata"`
}

func (op *OriginationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}
func (*OriginationContentsAndResult) OperationContentsAndResult() {}

//json:kind=OperationKind()
type OriginationInternalOperationResult struct {
	Source   TransactionDestination      `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Balance  tz.BigUint                  `json:"balance"`
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
	Script   Script                      `json:"script"`
	Result   OriginationResult           `json:"result"`
}

func (r *OriginationInternalOperationResult) GetSource() core.Address { return r.Source }
func (r *OriginationInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*OriginationInternalOperationResult) OperationKind() string { return "origination" }

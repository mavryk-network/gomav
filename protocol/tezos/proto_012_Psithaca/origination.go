package proto_012_Psithaca

import (
	"math/big"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_012_Psithaca/big_map"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_012_Psithaca/lazy"
)

//json:kind=OperationKind()
type Origination struct {
	ManagerOperation
	Balance  mv.BigUint                  `json:"balance"`
	Delegate mv.Option[mv.PublicKeyHash] `json:"delegate"`
	Script   Script                      `json:"script"`
}

func (*Origination) OperationKind() string { return "origination" }

type OriginationResultContents struct {
	BigMapDiff mv.Option[big_map.Diff] `json:"big_map_diff"`
	BalanceUpdates
	OriginatedContracts []core.OriginatedContractID `mv:"dyn" json:"originated_contracts"`
	ConsumedGas         mv.BigUint                  `json:"consumed_gas"`
	ConsumedMilligas    mv.BigUint                  `json:"consumed_milligas"`
	StorageSize         mv.BigInt                   `json:"storage_size"`
	PaidStorageSizeDiff mv.BigInt                   `json:"paid_storage_size_diff"`
	LazyStorageDiff     mv.Option[lazy.StorageDiff] `json:"lazy_storage_diff"`
}

func (r *OriginationResultContents) GetConsumedMilligas() mv.BigUint   { return r.ConsumedMilligas }
func (r *OriginationResultContents) GetStorageSize() mv.BigInt         { return r.StorageSize }
func (r *OriginationResultContents) GetPaidStorageSizeDiff() mv.BigInt { return r.PaidStorageSizeDiff }
func (r *OriginationResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	x := r.PaidStorageSizeDiff.Int()
	x.Add(x, big.NewInt(int64(constants.GetOriginationSize())))
	return x
}

//json:kind=OperationKind()
type OriginationSuccessfulManagerResult struct {
	core.OperationResultApplied[*OriginationResultContents]
}

func (*OriginationSuccessfulManagerResult) OperationKind() string { return "origination" }

type OriginationResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OriginationResult]{
		Variants: encoding.Variants[OriginationResult]{
			0: (*core.OperationResultApplied[*OriginationResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*OriginationResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type OriginationContentsAndResult struct {
	Origination
	Metadata ManagerMetadata[OriginationResult] `json:"metadata"`
}

func (*OriginationContentsAndResult) OperationContentsAndResult() {}
func (op *OriginationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type OriginationInternalOperationResult struct {
	Source   core.ContractID             `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Balance  mv.BigUint                  `json:"balance"`
	Delegate mv.Option[mv.PublicKeyHash] `json:"delegate"`
	Script   Script                      `json:"script"`
	Result   OriginationResult           `json:"result"`
}

var _ core.InternalOperationResult = (*OriginationInternalOperationResult)(nil)

func (r *OriginationInternalOperationResult) GetSource() core.TransactionDestination {
	switch d := r.Source.(type) {
	case core.ImplicitContract:
		return d
	case core.OriginatedContract:
		return d
	default:
		panic("unexpected contract id type")
	}
}
func (r *OriginationInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*OriginationInternalOperationResult) OperationKind() string { return "origination" }

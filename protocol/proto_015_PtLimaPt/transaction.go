package proto_015_PtLimaPt

import (
	"bytes"
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca/lazy"
	"github.com/ecadlabs/gotez/v2/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/v2/protocol/proto_014_PtKathma"
)

type Transaction = proto_012_Psithaca.Transaction
type Parameters = proto_012_Psithaca.Parameters
type TxRollupDestination = proto_013_PtJakart.TxRollupDestination
type ScRollupDestination = proto_014_PtKathma.ScRollupDestination
type ToScRollup = proto_014_PtKathma.ToScRollup

type TransactionResultDestination interface {
	proto_013_PtJakart.TransactionResultDestination
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResultDestination]{
		Variants: encoding.Variants[TransactionResultDestination]{
			0: (*ToContract)(nil),
			1: (*ToTxRollup)(nil),
			2: (*ToScRollup)(nil),
		},
	})
}

type ToContract struct {
	Storage tz.Option[expression.Expression] `json:"storage"`
	BalanceUpdates
	TicketUpdates                []*TicketReceipt            `tz:"dyn" json:"ticket_updates"`
	OriginatedContracts          []core.OriginatedContractID `tz:"dyn" json:"originated_contracts"`
	ConsumedMilligas             tz.BigUint                  `json:"consumed_milligas"`
	StorageSize                  tz.BigInt                   `json:"storage_size"`
	PaidStorageSizeDiff          tz.BigInt                   `json:"paid_storage_size_diff"`
	AllocatedDestinationContract bool                        `json:"allocated_destination_contract"`
	LazyStorageDiff              tz.Option[lazy.StorageDiff] `json:"lazy_storage_diff"`
}

func (*ToContract) TransactionResultDestination()       {}
func (r *ToContract) GetConsumedMilligas() tz.BigUint   { return r.ConsumedMilligas }
func (r *ToContract) GetStorageSize() tz.BigInt         { return r.StorageSize }
func (r *ToContract) GetPaidStorageSizeDiff() tz.BigInt { return r.PaidStorageSizeDiff }
func (r *ToContract) EstimateStorageSize(constants core.Constants) *big.Int {
	x := r.PaidStorageSizeDiff.Int()
	if r.AllocatedDestinationContract {
		x.Add(x, big.NewInt(int64(constants.GetOriginationSize())))
	}
	return x
}

type ToTxRollup struct {
	BalanceUpdates
	ConsumedMilligas    tz.BigUint         `json:"consumed_milligas"`
	TicketHash          *tz.ScriptExprHash `json:"ticket_hash"`
	PaidStorageSizeDiff tz.BigUint         `json:"paid_storage_size_diff"`
}

func (*ToTxRollup) TransactionResultDestination()     {}
func (r *ToTxRollup) GetConsumedMilligas() tz.BigUint { return r.ConsumedMilligas }
func (r *ToTxRollup) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.PaidStorageSizeDiff.Int()
}

type TicketReceipt struct {
	TicketToken TicketToken            `json:"ticket_token"`
	Updates     []*TicketReceiptUpdate `tz:"dyn" json:"updates"`
}

type TicketToken struct {
	Ticketer    core.ContractID       `json:"ticketer"`
	ContentType expression.Expression `json:"content_type"`
	Content     expression.Expression `json:"content"`
}

type TicketReceiptUpdate struct {
	Account TransactionDestination `json:"account"`
	Amount  tz.BigInt              `json:"amount"`
}

type TransactionDestination interface {
	core.TransactionDestination
}

type ZkRollupDestination struct {
	*tz.ZkRollupAddress
	Padding uint8
}

func (ZkRollupDestination) TransactionDestination() {}
func (a ZkRollupDestination) Eq(b core.TransactionDestination) bool {
	if other, ok := b.(ZkRollupDestination); ok {
		return bytes.Equal(a.ZkRollupAddress[:], other.ZkRollupAddress[:])
	}
	return false
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionDestination]{
		Variants: encoding.Variants[TransactionDestination]{
			0: core.ImplicitContract{},
			1: core.OriginatedContract{},
			2: TxRollupDestination{},
			3: ScRollupDestination{},
			4: ZkRollupDestination{},
		},
	})
}

type TransactionResultContents = TransactionResultDestination

//json:kind=OperationKind()
type TransactionSuccessfulManagerResult struct {
	core.OperationResultApplied[TransactionResultContents]
}

func (TransactionSuccessfulManagerResult) OperationKind() string { return "transaction" }

type TransactionContentsAndResult struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult] `json:"metadata"`
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}
func (op *TransactionContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type TransactionResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResult]{
		Variants: encoding.Variants[TransactionResult]{
			0: (*core.OperationResultApplied[TransactionResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[TransactionResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type TransactionInternalOperationResult struct {
	Source      core.ContractID        `json:"source"`
	Nonce       uint16                 `json:"nonce"`
	Amount      tz.BigUint             `json:"amount"`
	Destination TransactionDestination `json:"destination"`
	Parameters  tz.Option[Parameters]  `json:"parameters"`
	Result      TransactionResult      `json:"result"`
}

var _ core.TransactionInternalOperationResult = (*TransactionInternalOperationResult)(nil)

func (r *TransactionInternalOperationResult) GetSource() core.TransactionDestination {
	switch d := r.Source.(type) {
	case core.ImplicitContract:
		return d
	case core.OriginatedContract:
		return d
	default:
		panic("unexpected contract id type")
	}
}
func (r *TransactionInternalOperationResult) GetSourceAddress() core.ContractID { return r.Source }
func (r *TransactionInternalOperationResult) GetNonce() uint16                  { return r.Nonce }
func (t *TransactionInternalOperationResult) GetAmount() tz.BigUint             { return t.Amount }
func (t *TransactionInternalOperationResult) GetDestination() core.TransactionDestination {
	return t.Destination
}
func (t *TransactionInternalOperationResult) GetParameters() tz.Option[core.Parameters] {
	if p, ok := t.Parameters.CheckUnwrapPtr(); ok {
		return tz.Some[core.Parameters](p)
	}
	return tz.None[core.Parameters]()
}
func (r *TransactionInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*TransactionInternalOperationResult) OperationKind() string { return "transaction" }

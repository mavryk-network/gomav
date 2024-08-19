package proto_016_PtMumbai

import (
	"math/big"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/core/expression"
	"github.com/mavryk-network/gomav/v2/protocol/proto_012_Psithaca"
	"github.com/mavryk-network/gomav/v2/protocol/proto_012_Psithaca/lazy"
	"github.com/mavryk-network/gomav/v2/protocol/proto_013_PtJakart"
	"github.com/mavryk-network/gomav/v2/protocol/proto_015_PtLimaPt"
)

type Transaction = proto_015_PtLimaPt.Transaction
type Parameters = proto_015_PtLimaPt.Parameters
type Entrypoint = proto_015_PtLimaPt.Entrypoint
type EpDefault = proto_012_Psithaca.EpDefault
type EpRoot = proto_012_Psithaca.EpRoot
type EpDo = proto_012_Psithaca.EpDo
type EpSetDelegate = proto_012_Psithaca.EpSetDelegate
type EpRemoveDelegate = proto_012_Psithaca.EpRemoveDelegate
type EpDeposit = proto_015_PtLimaPt.EpDeposit
type EpNamed = proto_012_Psithaca.EpNamed
type TicketReceipt = proto_015_PtLimaPt.TicketReceipt

type TransactionResultDestination interface {
	proto_013_PtJakart.TransactionResultDestination
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResultDestination]{
		Variants: encoding.Variants[TransactionResultDestination]{
			0: (*ToContract)(nil),
			1: (*ToTxRollup)(nil),
			2: (*ToSmartRollup)(nil),
		},
	})
}

type TransactionResultContents = TransactionResultDestination

//json:kind=OperationKind()
type TransactionSuccessfulManagerResult struct {
	core.OperationResultApplied[TransactionResultContents]
}

func (TransactionSuccessfulManagerResult) OperationKind() string { return "transaction" }

type ToContract struct {
	Storage mv.Option[expression.Expression] `json:"storage"`
	BalanceUpdates
	TicketUpdates                []*TicketReceipt            `mv:"dyn" json:"ticket_updates"`
	OriginatedContracts          []core.OriginatedContractID `mv:"dyn" json:"originated_contracts"`
	ConsumedMilligas             mv.BigUint                  `json:"consumed_milligas"`
	StorageSize                  mv.BigInt                   `json:"storage_size"`
	PaidStorageSizeDiff          mv.BigInt                   `json:"paid_storage_size_diff"`
	AllocatedDestinationContract bool                        `json:"allocated_destination_contract"`
	LazyStorageDiff              mv.Option[lazy.StorageDiff] `json:"lazy_storage_diff"`
}

func (*ToContract) TransactionResultDestination()       {}
func (r *ToContract) GetConsumedMilligas() mv.BigUint   { return r.ConsumedMilligas }
func (r *ToContract) GetStorageSize() mv.BigInt         { return r.StorageSize }
func (r *ToContract) GetPaidStorageSizeDiff() mv.BigInt { return r.PaidStorageSizeDiff }
func (r *ToContract) EstimateStorageSize(constants core.Constants) *big.Int {
	x := r.PaidStorageSizeDiff.Int()
	if r.AllocatedDestinationContract {
		x.Add(x, big.NewInt(int64(constants.GetOriginationSize())))
	}
	return x
}

type ToTxRollup struct {
	BalanceUpdates
	ConsumedMilligas    mv.BigUint         `json:"consumed_milligas"`
	TicketHash          *mv.ScriptExprHash `json:"ticket_hash"`
	PaidStorageSizeDiff mv.BigUint         `json:"paid_storage_size_diff"`
}

func (*ToTxRollup) TransactionResultDestination()     {}
func (r *ToTxRollup) GetConsumedMilligas() mv.BigUint { return r.ConsumedMilligas }
func (r *ToTxRollup) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.PaidStorageSizeDiff.Int()
}

type ToSmartRollup struct {
	ConsumedMilligas mv.BigUint       `json:"consumed_milligas"`
	TicketUpdates    []*TicketReceipt `mv:"dyn" json:"ticket_updates"`
}

func (*ToSmartRollup) TransactionResultDestination()     {}
func (r *ToSmartRollup) GetConsumedMilligas() mv.BigUint { return r.ConsumedMilligas }

//json:kind=OperationKind()
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
	Source      core.TransactionDestination `json:"source"`
	Nonce       uint16                      `json:"nonce"`
	Amount      mv.BigUint                  `json:"amount"`
	Destination core.TransactionDestination `json:"destination"`
	Parameters  mv.Option[Parameters]       `json:"parameters"`
	Result      TransactionResult           `json:"result"`
}

var _ core.TransactionInternalOperationResult = (*TransactionInternalOperationResult)(nil)

func (r *TransactionInternalOperationResult) GetSource() core.TransactionDestination { return r.Source }
func (r *TransactionInternalOperationResult) GetNonce() uint16                       { return r.Nonce }
func (t *TransactionInternalOperationResult) GetAmount() mv.BigUint                  { return t.Amount }
func (t *TransactionInternalOperationResult) GetDestination() core.TransactionDestination {
	return t.Destination
}
func (t *TransactionInternalOperationResult) GetParameters() mv.Option[core.Parameters] {
	if p, ok := t.Parameters.CheckUnwrapPtr(); ok {
		return mv.Some[core.Parameters](p)
	}
	return mv.None[core.Parameters]()
}
func (r *TransactionInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*TransactionInternalOperationResult) OperationKind() string { return "transaction" }

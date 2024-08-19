package proto_014_PtKathma

import (
	"math/big"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/core/expression"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_012_Psithaca"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_012_Psithaca/lazy"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_013_PtJakart"
)

type Transaction = proto_012_Psithaca.Transaction
type Parameters = proto_012_Psithaca.Parameters

type ToScRollup struct {
	ConsumedMilligas mv.BigUint    `json:"consumed_milligas"`
	InboxAfter       ScRollupInbox `json:"inbox_after"`
}

func (r *ToScRollup) GetConsumedMilligas() mv.BigUint { return r.ConsumedMilligas }

type ScRollupInbox struct {
	Rollup                                 *mv.ScRollupAddress `mv:"dyn" json:"rollup"`
	MessageCounter                         mv.BigInt           `json:"message_counter"`
	NbMessagesInCommitmentPeriod           int64               `json:"nb_messages_in_commitment_period"`
	StartingLevelOfCurrentCommitmentPeriod int32               `json:"starting_level_of_current_commitment_period"`
	Level                                  int32               `json:"level"`
	CurrentLevelHash                       *mv.Bytes32         `json:"current_level_hash"`
	OldLevelsMessages                      OldLevelsMessages   `json:"old_levels_messages"`
}

type OldLevelsMessages struct {
	Index        int32       `json:"index"`
	Content      *mv.Bytes32 `json:"content"`
	BackPointers mv.Bytes    `mv:"dyn" json:"back_pointers"`
}

func (*ToScRollup) TransactionResultDestination() {}

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
	Storage mv.Option[expression.Expression] `json:"storage"`
	BalanceUpdates
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

type TransactionResultContents = TransactionResultDestination

//json:kind=OperationKind()
type TransactionSuccessfulManagerResult struct {
	core.OperationResultApplied[TransactionResultContents]
}

func (*TransactionSuccessfulManagerResult) OperationKind() string { return "transaction" }

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
	Source      core.ContractID             `json:"source"`
	Nonce       uint16                      `json:"nonce"`
	Amount      mv.BigUint                  `json:"amount"`
	Destination core.TransactionDestination `json:"destination"`
	Parameters  mv.Option[Parameters]       `json:"parameters"`
	Result      TransactionResult           `json:"result"`
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
func (r *TransactionInternalOperationResult) GetNonce() uint16      { return r.Nonce }
func (t *TransactionInternalOperationResult) GetAmount() mv.BigUint { return t.Amount }
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

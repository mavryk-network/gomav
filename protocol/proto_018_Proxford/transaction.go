package proto_018_Proxford

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
	"github.com/mavryk-network/gomav/v2/protocol/proto_016_PtMumbai"
)

type TicketReceipt = proto_015_PtLimaPt.TicketReceipt
type ToSmartRollup = proto_016_PtMumbai.ToSmartRollup

//json:kind=OperationKind()
type Transaction struct {
	ManagerOperation
	Amount      mv.BigUint            `json:"amount"`
	Destination core.ContractID       `json:"destination"`
	Parameters  mv.Option[Parameters] `json:"parameters"`
}

func (*Transaction) OperationKind() string          { return "transaction" }
func (t *Transaction) GetAmount() mv.BigUint        { return t.Amount }
func (t *Transaction) GetDestination() core.Address { return t.Destination }
func (t *Transaction) GetParameters() mv.Option[core.Parameters] {
	if p, ok := t.Parameters.CheckUnwrapPtr(); ok {
		return mv.Some[core.Parameters](p)
	}
	return mv.None[core.Parameters]()
}

var _ core.Transaction = (*Transaction)(nil)

type Parameters struct {
	Entrypoint Entrypoint            `json:"entrypoint"`
	Value      expression.Expression `mv:"dyn" json:"value"`
}

func (p *Parameters) GetEntrypoint() string           { return p.Entrypoint.Entrypoint() }
func (p *Parameters) GetValue() expression.Expression { return p.Value }

type Entrypoint interface {
	core.Entrypoint
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[Entrypoint]{
		Variants: encoding.Variants[Entrypoint]{
			0:   EpDefault{},
			1:   EpRoot{},
			2:   EpDo{},
			3:   EpSetDelegate{},
			4:   EpRemoveDelegate{},
			5:   EpDeposit{},
			6:   EpStake{},
			7:   EpUnstake{},
			8:   EpFinalizeUnstake{},
			9:   EpSetDelegateParameters{},
			255: EpNamed{},
		},
	})
}

type EpDefault = proto_012_Psithaca.EpDefault
type EpRoot = proto_012_Psithaca.EpRoot
type EpDo = proto_012_Psithaca.EpDo
type EpSetDelegate = proto_012_Psithaca.EpSetDelegate
type EpRemoveDelegate = proto_012_Psithaca.EpRemoveDelegate
type EpDeposit = proto_015_PtLimaPt.EpDeposit
type EpNamed = proto_012_Psithaca.EpNamed

type EpStake struct{}
type EpUnstake struct{}
type EpFinalizeUnstake struct{}
type EpSetDelegateParameters struct{}

func (EpStake) Entrypoint() string                         { return "stake" }
func (ep EpStake) MarshalText() (text []byte, err error)   { return []byte(ep.Entrypoint()), nil }
func (EpUnstake) Entrypoint() string                       { return "unstake" }
func (ep EpUnstake) MarshalText() (text []byte, err error) { return []byte(ep.Entrypoint()), nil }
func (EpFinalizeUnstake) Entrypoint() string               { return "finalize_unstake" }
func (ep EpFinalizeUnstake) MarshalText() (text []byte, err error) {
	return []byte(ep.Entrypoint()), nil
}
func (EpSetDelegateParameters) Entrypoint() string { return "set_delegate_parameters" }
func (ep EpSetDelegateParameters) MarshalText() (text []byte, err error) {
	return []byte(ep.Entrypoint()), nil
}

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

type TransactionResultDestination interface {
	proto_013_PtJakart.TransactionResultDestination
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResultDestination]{
		Variants: encoding.Variants[TransactionResultDestination]{
			0: (*ToContract)(nil),
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
type TransactionContentsAndResult struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult] `json:"metadata"`
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}
func (op *TransactionContentsAndResult) GetMetadata() any {
	return &op.Metadata
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

package proto_014_PtKathma

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca/lazy"
	"github.com/ecadlabs/gotez/protocol/proto_013_PtJakart"
)

type Transaction = proto_012_Psithaca.Transaction
type Parameters = proto_012_Psithaca.Parameters
type TxRollupDestination = proto_013_PtJakart.TxRollupDestination

type ToScRollup struct {
	ConsumedMilligas tz.BigUint
	InboxAfter       ScRollupInbox
}

type ScRollupInbox struct {
	Rollup                                 *tz.ScRollupAddress `tz:"dyn"`
	MessageCounter                         tz.BigInt
	NbMessagesInCommitmentPeriod           int64
	StartingLevelOfCurrentCommitmentPeriod int32
	Level                                  int32
	CurrentLevelHash                       *[32]byte
	OldLevelsMessages                      OldLevelsMessages
}

type OldLevelsMessages struct {
	Index        int32
	Content      *[32]byte
	BackPointers []byte `tz:"dyn"`
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
	Storage                      tz.Option[expression.Expression]
	BalanceUpdates               []*BalanceUpdate            `tz:"dyn"`
	OriginatedContracts          []core.OriginatedContractID `tz:"dyn"`
	ConsumedMilligas             tz.BigUint
	StorageSize                  tz.BigInt
	PaidStorageSizeDiff          tz.BigInt
	AllocatedDestinationContract bool
	LazyStorageDiff              tz.Option[lazy.StorageDiff]
}

func (*ToContract) TransactionResultDestination() {}

type ToTxRollup struct {
	BalanceUpdates      []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	TicketHash          *tz.ScriptExprHash
	PaidStorageSizeDiff tz.BigUint
}

func (*ToTxRollup) TransactionResultDestination() {}

type ScRollupDestination struct {
	*tz.ScRollupAddress
	Padding uint8
}

func (*ScRollupDestination) TransactionDestination() {}

type TransactionDestination interface {
	core.TransactionDestination
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionDestination]{
		Variants: encoding.Variants[TransactionDestination]{
			0: (*core.ImplicitContract)(nil),
			1: (*core.OriginatedContract)(nil),
			2: (*TxRollupDestination)(nil),
			3: (*ScRollupDestination)(nil),
		},
	})
}

type TransactionResultContents struct {
	Result TransactionResultDestination
}

func (TransactionResultContents) SuccessfulManagerOperationResult() {}
func (TransactionResultContents) OperationKind() string             { return "transaction" }

type TransactionContentsAndResult struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult]
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}
func (op *TransactionContentsAndResult) OperationContents() core.OperationContents {
	return &op.Transaction
}

type TransactionResultApplied struct {
	core.OperationResultApplied[TransactionResultContents]
}

func (*TransactionResultApplied) TransactionResult() {}

type TransactionResultBacktracked struct {
	core.OperationResultBacktracked[TransactionResultContents]
}

func (*TransactionResultBacktracked) TransactionResult() {}

type TransactionResultFailed struct{ core.OperationResultFailed }

func (*TransactionResultFailed) TransactionResult() {}

type TransactionResultSkipped struct{ core.OperationResultSkipped }

func (*TransactionResultSkipped) TransactionResult() {}

type TransactionResult interface {
	proto_012_Psithaca.TransactionResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResult]{
		Variants: encoding.Variants[TransactionResult]{
			0: (*TransactionResultApplied)(nil),
			1: (*TransactionResultFailed)(nil),
			2: (*TransactionResultSkipped)(nil),
			3: (*TransactionResultBacktracked)(nil),
		},
	})
}

type TransactionInternalOperationResult struct {
	Source      core.ContractID
	Nonce       uint16
	Amount      tz.BigUint
	Destination TransactionDestination
	Parameters  tz.Option[Parameters]
	Result      TransactionResult
}

func (*TransactionInternalOperationResult) InternalOperationResult() {}
func (*TransactionInternalOperationResult) OperationKind() string    { return "transaction" }
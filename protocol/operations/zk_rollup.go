package operations

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
)

type ZkRollupOrigination struct {
	ManagerOperation
	PublicParameters []byte              `tz:"dyn"`
	CircuitsInfo     []*CircuitsInfoElem `tz:"dyn"`
	InitState        []byte              `tz:"dyn"`
	NbOps            int32
}

func (*ZkRollupOrigination) OperationKind() string { return "zk_rollup_origination" }

type CircuitsInfoElem struct {
	Value string `tz:"dyn"`
	Tag   CircuitsInfoTag
}

type CircuitsInfoTag uint8

const (
	CircuitsInfoPublic CircuitsInfoTag = iota
	CircuitsInfoPrivate
	CircuitsInfoFee
)

type ZkRollupOriginationContentsAndResult struct {
	ZkRollupOrigination
	Metadata MetadataWithResult[ZkRollupPublishResult]
}

func (*ZkRollupOriginationContentsAndResult) OperationContentsAndResult() {}

type ZkRollupPublish struct {
	ManagerOperation
	ZkRollup *tz.ZkRollupAddress
	Op       []*ZkRollupOpElem `tz:"dyn"`
}

func (*ZkRollupPublish) OperationKind() string { return "zk_rollup_publish" }

type ZkRollupOpElem struct {
	Op     ZkRollupOp
	Ticket tz.Option1[ZkRollupTicket]
}

type ZkRollupOp struct {
	OpCode   int32
	Price    ZkRollupPrice
	L1Dst    tz.PublicKeyHash
	RollupID *tz.ZkRollupAddress
	Payload  []byte `tz:"dyn"`
}

type ZkRollupPrice struct {
	ID     *tz.ScriptExprHash
	Amount tz.BigInt
}

type ZkRollupTicket struct {
	Contents expression.Expression
	Ty       expression.Expression
	Ticketer tz.ContractID
}

type ZkRollupPublishResultContents struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas tz.BigUint
	Size             tz.BigInt
}

type ZkRollupPublishResult interface {
	ZkRollupPublishResult()
	OperationResult
}

type ZkRollupPublishResultApplied struct {
	OperationResultApplied[ZkRollupPublishResultContents]
}

func (*ZkRollupPublishResultApplied) ZkRollupPublishResult() {}

type ZkRollupPublishResultBacktracked struct {
	OperationResultBacktracked[ZkRollupPublishResultContents]
}

func (*ZkRollupPublishResultBacktracked) ZkRollupPublishResult() {}

type ZkRollupPublishResultFailed struct{ OperationResultFailed }

func (*ZkRollupPublishResultFailed) ZkRollupPublishResult() {}

type ZkRollupPublishResultSkipped struct{ OperationResultSkipped }

func (*ZkRollupPublishResultSkipped) ZkRollupPublishResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ZkRollupPublishResult]{
		Variants: encoding.Variants[ZkRollupPublishResult]{
			0: (*ZkRollupPublishResultApplied)(nil),
			1: (*ZkRollupPublishResultFailed)(nil),
			2: (*ZkRollupPublishResultSkipped)(nil),
			3: (*ZkRollupPublishResultBacktracked)(nil),
		},
	})
}

type ZkRollupPublishContentsAndResult struct {
	ZkRollupPublish
	Metadata MetadataWithResult[ZkRollupPublishResult]
}

func (*ZkRollupPublishContentsAndResult) OperationContentsAndResult() {}

type ZkRollupUpdate struct {
	ManagerOperation
	ZkRollup *tz.ZkRollupAddress
	Update   ZkRollupUpdateContents
}

func (*ZkRollupUpdate) OperationKind() string { return "zk_rollup_update" }

type ZkRollupUpdateContents struct {
	PendingPis []*PendingPiElem `tz:"dyn"`
	PrivatePis []*PrivatePiElem `tz:"dyn"`
	FeePi      FeePi
	Proof      []byte `tz:"dyn"`
}

type PendingPiElem struct {
	Key string `tz:"dyn"`
	Pi  PendingPi
}

type ZkRollupScalar [32]byte

type PendingPi struct {
	NewState     []byte `tz:"dyn"`
	Fee          ZkRollupScalar
	ExitValidity bool
}

type PrivatePiElem struct {
	Key string `tz:"dyn"`
	Pi  PrivatePi
}

type PrivatePi struct {
	NewState []byte `tz:"dyn"`
	Fee      ZkRollupScalar
}

type FeePi struct {
	NewState []byte `tz:"dyn"`
}

type ZkRollupUpdateResultContents struct {
	BalanceUpdates      []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	PaidStorageSizeDiff tz.BigInt
}

type ZkRollupUpdateResult interface {
	ZkRollupUpdateResult()
	OperationResult
}

type ZkRollupUpdateResultApplied struct {
	OperationResultApplied[ZkRollupUpdateResultContents]
}

func (*ZkRollupUpdateResultApplied) ZkRollupUpdateResult() {}

type ZkRollupUpdateResultBacktracked struct {
	OperationResultBacktracked[ZkRollupUpdateResultContents]
}

func (*ZkRollupUpdateResultBacktracked) ZkRollupUpdateResult() {}

type ZkRollupUpdateResultFailed struct{ OperationResultFailed }

func (*ZkRollupUpdateResultFailed) ZkRollupUpdateResult() {}

type ZkRollupUpdateResultSkipped struct{ OperationResultSkipped }

func (*ZkRollupUpdateResultSkipped) ZkRollupUpdateResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ZkRollupUpdateResult]{
		Variants: encoding.Variants[ZkRollupUpdateResult]{
			0: (*ZkRollupUpdateResultApplied)(nil),
			1: (*ZkRollupUpdateResultFailed)(nil),
			2: (*ZkRollupUpdateResultSkipped)(nil),
			3: (*ZkRollupUpdateResultBacktracked)(nil),
		},
	})
}

type ZkRollupUpdateContentsAndResult struct {
	ZkRollupUpdate
	Metadata MetadataWithResult[ZkRollupUpdateResult]
}

func (*ZkRollupUpdateContentsAndResult) OperationContentsAndResult() {}

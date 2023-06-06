package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_015_PtLimaPt"
)

type ZkRollupOrigination = proto_015_PtLimaPt.ZkRollupOrigination
type ZkRollupPublish = proto_015_PtLimaPt.ZkRollupPublish

type ZkRollupOriginationContentsAndResult struct {
	ZkRollupOrigination
	Metadata ManagerMetadata[ZkRollupPublishResult]
}

func (*ZkRollupOriginationContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupOriginationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type ZkRollupPublishResultContents struct {
	BalanceUpdates
	ConsumedMilligas tz.BigUint
	Size             tz.BigInt
}

type ZkRollupPublishResult interface {
	ZkRollupPublishResult()
	core.ManagerOperationResult
}

type ZkRollupPublishResultApplied struct {
	core.OperationResultApplied[ZkRollupPublishResultContents]
}

func (*ZkRollupPublishResultApplied) ZkRollupPublishResult() {}

type ZkRollupPublishResultBacktracked struct {
	core.OperationResultBacktracked[ZkRollupPublishResultContents]
}

func (*ZkRollupPublishResultBacktracked) ZkRollupPublishResult() {}

type ZkRollupPublishResultFailed struct{ core.OperationResultFailed }

func (*ZkRollupPublishResultFailed) ZkRollupPublishResult() {}

type ZkRollupPublishResultSkipped struct{ core.OperationResultSkipped }

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
	Metadata ManagerMetadata[ZkRollupPublishResult]
}

func (*ZkRollupPublishContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupPublishContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

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
	BalanceUpdates
	ConsumedMilligas    tz.BigUint
	PaidStorageSizeDiff tz.BigInt
}

type ZkRollupUpdateResult interface {
	ZkRollupUpdateResult()
	core.ManagerOperationResult
}

type ZkRollupUpdateResultApplied struct {
	core.OperationResultApplied[ZkRollupUpdateResultContents]
}

func (*ZkRollupUpdateResultApplied) ZkRollupUpdateResult() {}

type ZkRollupUpdateResultBacktracked struct {
	core.OperationResultBacktracked[ZkRollupUpdateResultContents]
}

func (*ZkRollupUpdateResultBacktracked) ZkRollupUpdateResult() {}

type ZkRollupUpdateResultFailed struct{ core.OperationResultFailed }

func (*ZkRollupUpdateResultFailed) ZkRollupUpdateResult() {}

type ZkRollupUpdateResultSkipped struct{ core.OperationResultSkipped }

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
	Metadata ManagerMetadata[ZkRollupUpdateResult]
}

func (*ZkRollupUpdateContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupUpdateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

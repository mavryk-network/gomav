package proto_002_PtBoreas

import (
	"math/big"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/proto_001_PtAtLas"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_016_PtMumbai"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_017_PtNairob"
)

type SmartRollupAddMessages = proto_016_PtMumbai.SmartRollupAddMessages
type SmartRollupPublish = proto_016_PtMumbai.SmartRollupPublish
type SmartRollupRefute = proto_001_PtAtLas.SmartRollupRefute
type SmartRollupTimeout = proto_016_PtMumbai.SmartRollupTimeout
type SmartRollupExecuteOutboxMessage = proto_016_PtMumbai.SmartRollupExecuteOutboxMessage
type SmartRollupRecoverBond = proto_016_PtMumbai.SmartRollupRecoverBond
type SmartRollupOriginate = proto_001_PtAtLas.SmartRollupOriginate
type SmartRollupCement = proto_001_PtAtLas.SmartRollupCement
type SmartRollupCementResult = proto_017_PtNairob.SmartRollupCementResult
type GameStatus = proto_016_PtMumbai.GameStatus

type SmartRollupOriginateResultContents struct {
	BalanceUpdates
	Address               *mv.SmartRollupAddress        `json:"address"`
	GenesisCommitmentHash *mv.SmartRollupCommitmentHash `json:"genesis_commitment_hash"`
	ConsumedMilligas      mv.BigUint                    `json:"consumed_milligas"`
	Size                  mv.BigInt                     `json:"size"`
}

func (r *SmartRollupOriginateResultContents) GetConsumedMilligas() mv.BigUint {
	return r.ConsumedMilligas
}
func (r *SmartRollupOriginateResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.Size.Int()
}

type SmartRollupOriginateResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupOriginateResult]{
		Variants: encoding.Variants[SmartRollupOriginateResult]{
			0: (*core.OperationResultApplied[*SmartRollupOriginateResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupOriginateResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type SmartRollupOriginateContentsAndResult struct {
	SmartRollupOriginate
	Metadata ManagerMetadata[SmartRollupOriginateResult] `json:"metadata"`
}

func (*SmartRollupOriginateContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupOriginateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupAddMessagesContentsAndResult struct {
	SmartRollupAddMessages
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*SmartRollupAddMessagesContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupAddMessagesContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupCementContentsAndResult struct {
	SmartRollupCement
	Metadata ManagerMetadata[SmartRollupCementResult] `json:"metadata"`
}

func (*SmartRollupCementContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupCementContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type SmartRollupPublishResultContents struct {
	ConsumedMilligas mv.BigUint                    `json:"consumed_milligas"`
	StakedHash       *mv.SmartRollupCommitmentHash `json:"staked_hash"`
	PublishedAtLevel int32                         `json:"published_at_level"`
	BalanceUpdates
}

func (r *SmartRollupPublishResultContents) GetConsumedMilligas() mv.BigUint {
	return r.ConsumedMilligas
}

type SmartRollupPublishResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupPublishResult]{
		Variants: encoding.Variants[SmartRollupPublishResult]{
			0: (*core.OperationResultApplied[*SmartRollupPublishResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupPublishResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type SmartRollupPublishContentsAndResult struct {
	SmartRollupPublish
	Metadata ManagerMetadata[SmartRollupPublishResult] `json:"metadata"`
}

func (*SmartRollupPublishContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupPublishContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type SmartRollupTimeoutResultContents struct {
	ConsumedMilligas mv.BigUint `json:"consumed_milligas"`
	GameStatus       GameStatus `json:"game_status"`
	BalanceUpdates
}

func (r *SmartRollupTimeoutResultContents) GetConsumedMilligas() mv.BigUint {
	return r.ConsumedMilligas
}

type SmartRollupTimeoutResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupTimeoutResult]{
		Variants: encoding.Variants[SmartRollupTimeoutResult]{
			0: (*core.OperationResultApplied[*SmartRollupTimeoutResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupTimeoutResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type SmartRollupTimeoutContentsAndResult struct {
	SmartRollupTimeout
	Metadata ManagerMetadata[SmartRollupTimeoutResult] `json:"metadata"`
}

func (*SmartRollupTimeoutContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupTimeoutContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupRefuteContentsAndResult struct {
	SmartRollupRefute
	Metadata ManagerMetadata[SmartRollupTimeoutResult] `json:"metadata"`
}

func (*SmartRollupRefuteContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupRefuteContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type SmartRollupExecuteOutboxMessageResultContents struct {
	BalanceUpdates
	TicketUpdates       []*TicketReceipt `mv:"dyn" json:"ticket_updates"`
	ConsumedMilligas    mv.BigUint       `json:"consumed_milligas"`
	PaidStorageSizeDiff mv.BigInt        `json:"paid_storage_size_diff"`
}

func (r *SmartRollupExecuteOutboxMessageResultContents) GetConsumedMilligas() mv.BigUint {
	return r.ConsumedMilligas
}

func (r *SmartRollupExecuteOutboxMessageResultContents) GetPaidStorageSizeDiff() mv.BigInt {
	return r.PaidStorageSizeDiff
}

func (r *SmartRollupExecuteOutboxMessageResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.PaidStorageSizeDiff.Int()
}

type SmartRollupExecuteOutboxMessageResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupExecuteOutboxMessageResult]{
		Variants: encoding.Variants[SmartRollupExecuteOutboxMessageResult]{
			0: (*core.OperationResultApplied[*SmartRollupExecuteOutboxMessageResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupExecuteOutboxMessageResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type SmartRollupExecuteOutboxMessageContentsAndResult struct {
	SmartRollupExecuteOutboxMessage
	Metadata ManagerMetadata[SmartRollupExecuteOutboxMessageResult] `json:"metadata"`
}

func (*SmartRollupExecuteOutboxMessageContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupExecuteOutboxMessageContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type SmartRollupRecoverBondResultContents struct {
	BalanceUpdates
	ConsumedMilligas mv.BigUint `json:"consumed_milligas"`
}

func (r *SmartRollupRecoverBondResultContents) GetConsumedMilligas() mv.BigUint {
	return r.ConsumedMilligas
}

type SmartRollupRecoverBondResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupRecoverBondResult]{
		Variants: encoding.Variants[SmartRollupRecoverBondResult]{
			0: (*core.OperationResultApplied[*SmartRollupRecoverBondResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupRecoverBondResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type SmartRollupRecoverBondContentsAndResult struct {
	SmartRollupRecoverBond
	Metadata ManagerMetadata[SmartRollupRecoverBondResult] `json:"metadata"`
}

func (*SmartRollupRecoverBondContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupRecoverBondContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupOriginateSuccessfulManagerResult struct {
	core.OperationResultApplied[*SmartRollupOriginateResultContents]
}

func (*SmartRollupOriginateSuccessfulManagerResult) OperationKind() string {
	return "smart_rollup_originate"
}

package proto_016_PtMumbai

import (
	"math/big"
	"strconv"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/core/expression"
)

type PVMKind uint8

const (
	PVMArith PVMKind = iota
	PVM_WASM_2_0_0
)

func (k PVMKind) String() string {
	switch k {
	case PVMArith:
		return "arith"
	case PVM_WASM_2_0_0:
		return "wasm_2_0_0"
	default:
		return strconv.FormatInt(int64(k), 10)
	}
}

func (k PVMKind) MarshalText() (text []byte, err error) {
	return []byte(k.String()), nil
}

//json:kind=OperationKind()
type SmartRollupOriginate struct {
	ManagerOperation
	PVMKind          PVMKind               `json:"pvm_kind"`
	Kernel           mv.Bytes              `mv:"dyn" json:"kernel"`
	OriginationProof mv.Bytes              `mv:"dyn" json:"origination_proof"`
	ParametersTy     expression.Expression `mv:"dyn" json:"parameters_ty"`
}

func (*SmartRollupOriginate) OperationKind() string { return "smart_rollup_originate" }

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

//json:kind=OperationKind()
type SmartRollupOriginateSuccessfulManagerResult struct {
	core.OperationResultApplied[*SmartRollupOriginateResultContents]
}

func (*SmartRollupOriginateSuccessfulManagerResult) OperationKind() string {
	return "smart_rollup_originate"
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
type SmartRollupAddMessages struct {
	ManagerOperation
	Message []core.Bytes `mv:"dyn" json:"message"`
}

func (*SmartRollupAddMessages) OperationKind() string { return "smart_rollup_add_messages" }

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
type SmartRollupCement struct {
	ManagerOperation
	Rollup     *mv.SmartRollupAddress        `json:"rollup"`
	Commitment *mv.SmartRollupCommitmentHash `json:"commitment"`
}

func (*SmartRollupCement) OperationKind() string { return "smart_rollup_cement" }

type SmartRollupCementResultContents struct {
	ConsumedMilligas mv.BigUint `json:"consumed_milligas"`
	InboxLevel       int32      `json:"inbox_level"`
}

func (r *SmartRollupCementResultContents) GetConsumedMilligas() mv.BigUint { return r.ConsumedMilligas }

type SmartRollupCementResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupCementResult]{
		Variants: encoding.Variants[SmartRollupCementResult]{
			0: (*core.OperationResultApplied[*SmartRollupCementResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupCementResultContents])(nil),
		},
	})
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

//json:kind=OperationKind()
type SmartRollupPublish struct {
	ManagerOperation
	Rollup     *mv.SmartRollupAddress `json:"rollup"`
	Commitment SmartRollupCommitment  `json:"commitment"`
}

func (*SmartRollupPublish) OperationKind() string { return "smart_rollup_publish" }

type SmartRollupCommitment struct {
	CompressedState *mv.SmartRollupStateHash      `json:"compressed_state"`
	InboxLevel      int32                         `json:"inbox_level"`
	Predecessor     *mv.SmartRollupCommitmentHash `json:"predecessor"`
	NumberOfTicks   int64                         `json:"number_of_ticks"`
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

//json:kind=OperationKind()
type SmartRollupRefute struct {
	ManagerOperation
	Rollup     *mv.SmartRollupAddress `json:"rollup"`
	Opponent   mv.PublicKeyHash       `json:"opponent"`
	Refutation SmartRollupRefutation  `json:"refutation"`
}

func (*SmartRollupRefute) OperationKind() string { return "smart_rollup_refute" }

type RefutationStart struct {
	PlayerCommitmentHash   *mv.SmartRollupCommitmentHash `json:"player_commitment_hash"`
	OpponentCommitmentHash *mv.SmartRollupCommitmentHash `json:"opponent_commitment_hash"`
}

func (*RefutationStart) RefutationKind() string { return "start" }

type RefutationMove struct {
	Choice mv.BigUint     `json:"choice"`
	Step   RefutationStep `json:"step"`
}

func (*RefutationMove) RefutationKind() string { return "move" }

type SmartRollupRefutation interface {
	RefutationKind() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupRefutation]{
		Variants: encoding.Variants[SmartRollupRefutation]{
			0: (*RefutationStart)(nil),
			1: (*RefutationMove)(nil),
		},
	})
}

type RefutationStepDissection struct {
	Contents []RefutationStepDissectionElem `mv:"dyn" json:"contents"`
}

func (*RefutationStepDissection) RefutationStepKind() string { return "dissection" }

type RefutationStepDissectionElem struct {
	State mv.Option[*mv.SmartRollupStateHash] `json:"state"`
	Tick  mv.BigUint                          `json:"tick"`
}

type RefutationStepProof struct {
	PVMStep    mv.Bytes                   `mv:"dyn" json:"pvm_step"`
	InputProof mv.Option[RefutationProof] `json:"input_proof"`
}

func (*RefutationStepProof) RefutationStepKind() string { return "proof" }

type RefutationStep interface {
	RefutationStepKind() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RefutationStep]{
		Variants: encoding.Variants[RefutationStep]{
			0: (*RefutationStepDissection)(nil),
			1: (*RefutationStepProof)(nil),
		},
	})
}

type RefutationProof interface {
	RefutationProof()
}

type RefutationProofInbox struct {
	Level           int32      `json:"level"`
	MessageCounter  mv.BigUint `json:"message_counter"`
	SerializedProof mv.Bytes   `mv:"dyn" json:"serialized_proof"`
}

func (*RefutationProofInbox) RefutationProof() {}

type RefutationProofReveal struct {
	RevealProof RevealProof `json:"reveal_proof"`
}

func (*RefutationProofReveal) RefutationProof() {}

type RefutationProofFirstInput struct{}

func (RefutationProofFirstInput) RefutationProof() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RefutationProof]{
		Variants: encoding.Variants[RefutationProof]{
			0: (*RefutationProofInbox)(nil),
			1: (*RefutationProofReveal)(nil),
			2: RefutationProofFirstInput{},
		},
	})
}

type RevealProof interface {
	RevealProof()
}

type RevealProofRawData struct {
	RawData mv.Bytes `mv:"dyn" json:"raw_data"`
}

func (RevealProofRawData) RevealProof() {}

type RevealProofMetadata struct{}

func (RevealProofMetadata) RevealProof() {}

type RevealProofDALPage struct {
	DALPageID DALPageID `json:"dal_page_id"`
	DALProof  mv.Bytes  `mv:"dyn" json:"dal_proof"`
}

func (*RevealProofDALPage) RevealProof() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RevealProof]{
		Variants: encoding.Variants[RevealProof]{
			0: RevealProofRawData{},
			1: RevealProofMetadata{},
			2: (*RevealProofDALPage)(nil),
		},
	})
}

type DALPageID struct {
	PublishedLevel int32 `json:"published_level"`
	SlotIndex      uint8 `json:"slot_index"`
	PageIndex      int16 `json:"page_index"`
}

type SmartRollupTimeoutResultContents struct {
	ConsumedMilligas mv.BigUint `json:"consumed_milligas"`
	GameStatus       GameStatus `json:"game_status"`
	BalanceUpdates
}

func (r *SmartRollupTimeoutResultContents) GetConsumedMilligas() mv.BigUint {
	return r.ConsumedMilligas
}

type GameStatus interface {
	GameStatusKind() string
}

type GameStatusOngoing struct{}

func (GameStatusOngoing) GameStatusKind() string { return "ongoing" }

type GameStatusEnded struct {
	Result GameStatusResult `json:"result"`
}

func (GameStatusEnded) GameStatusKind() string { return "ended" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[GameStatus]{
		Variants: encoding.Variants[GameStatus]{
			0: GameStatusOngoing{},
			1: GameStatusEnded{},
		},
	})
}

type GameStatusResult interface {
	GameStatusResultKind() string
}

type GameStatusResultLoser struct {
	Reason LooseReason      `json:"reason"`
	Player mv.PublicKeyHash `json:"player"`
}

func (*GameStatusResultLoser) GameStatusResultKind() string { return "loser" }

type GameStatusResultDraw struct{}

func (GameStatusResultDraw) GameStatusResultKind() string { return "draw" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[GameStatusResult]{
		Variants: encoding.Variants[GameStatusResult]{
			0: (*GameStatusResultLoser)(nil),
			1: GameStatusResultDraw{},
		},
	})
}

type LooseReason uint8

const (
	LooseReasonConflictResolved LooseReason = iota
	LooseReasonTimeout
)

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
type SmartRollupRefuteContentsAndResult struct {
	SmartRollupRefute
	Metadata ManagerMetadata[SmartRollupTimeoutResult] `json:"metadata"`
}

func (*SmartRollupRefuteContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupRefuteContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupTimeout struct {
	ManagerOperation
	Rollup  *mv.SmartRollupAddress `json:"rollup"`
	Stakers SmartRollupStakers     `json:"stakers"`
}

type SmartRollupStakers struct {
	Alice mv.PublicKeyHash `json:"alice"`
	Bob   mv.PublicKeyHash `json:"bob"`
}

func (*SmartRollupTimeout) OperationKind() string { return "smart_rollup_timeout" }

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
type SmartRollupExecuteOutboxMessage struct {
	ManagerOperation
	Rollup             *mv.SmartRollupAddress        `json:"rollup"`
	CementedCommitment *mv.SmartRollupCommitmentHash `json:"cemented_commitment"`
	OutputProof        mv.Bytes                      `mv:"dyn" json:"output_proof"`
}

func (*SmartRollupExecuteOutboxMessage) OperationKind() string {
	return "smart_rollup_execute_outbox_message"
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

//json:kind=OperationKind()
type SmartRollupRecoverBond struct {
	ManagerOperation
	Rollup *mv.SmartRollupAddress `json:"rollup"`
	Staker mv.PublicKeyHash       `json:"staker"`
}

func (*SmartRollupRecoverBond) OperationKind() string { return "smart_rollup_recover_bond" }

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

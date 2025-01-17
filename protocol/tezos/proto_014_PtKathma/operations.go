package proto_014_PtKathma

//go:generate go run ../../cmd/genmarshaller.go

import (
	"math/big"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/core/expression"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_012_Psithaca"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_013_PtJakart"
)

type ManagerOperation = proto_012_Psithaca.ManagerOperation
type SeedNonceRevelation = proto_012_Psithaca.SeedNonceRevelation
type Preendorsement = proto_012_Psithaca.Preendorsement
type InlinedPreendorsement = proto_012_Psithaca.InlinedPreendorsement
type Endorsement = proto_012_Psithaca.Endorsement
type InlinedEndorsement = proto_012_Psithaca.InlinedEndorsement
type DoublePreendorsementEvidence = proto_012_Psithaca.DoublePreendorsementEvidence
type DoubleEndorsementEvidence = proto_012_Psithaca.DoubleEndorsementEvidence
type Reveal = proto_012_Psithaca.Reveal
type Delegation = proto_012_Psithaca.Delegation
type RegisterGlobalConstant = proto_012_Psithaca.RegisterGlobalConstant
type SetDepositsLimit = proto_012_Psithaca.SetDepositsLimit
type ActivateAccount = proto_012_Psithaca.ActivateAccount
type Proposals = proto_012_Psithaca.Proposals
type Ballot = proto_012_Psithaca.Ballot
type FailingNoop = proto_012_Psithaca.FailingNoop
type TransferTicket = proto_013_PtJakart.TransferTicket
type Entrypoint = proto_012_Psithaca.Entrypoint
type DoubleBakingEvidence = proto_013_PtJakart.DoubleBakingEvidence

type ConsumedGasResultContents struct {
	ConsumedMilligas mv.BigUint `json:"consumed_milligas"`
}

func (r *ConsumedGasResultContents) GetConsumedMilligas() mv.BigUint { return r.ConsumedMilligas }

type RevealResultContents = ConsumedGasResultContents

//json:kind=OperationKind()
type RevealSuccessfulManagerResult struct {
	core.OperationResultApplied[*ConsumedGasResultContents]
}

func (*RevealSuccessfulManagerResult) OperationKind() string { return "reveal" }

type DelegationResultContents = ConsumedGasResultContents

//json:kind=OperationKind()
type DelegationSuccessfulManagerResult struct {
	core.OperationResultApplied[*ConsumedGasResultContents]
}

func (*DelegationSuccessfulManagerResult) OperationKind() string { return "delegation" }

//json:kind=OperationKind()
type DelegationInternalOperationResult struct {
	Source   core.TransactionDestination `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Delegate mv.Option[mv.PublicKeyHash] `json:"delegate"`
	Result   ConsumedGasResult           `json:"result"`
}

func (r *DelegationInternalOperationResult) GetSource() core.TransactionDestination { return r.Source }
func (*DelegationInternalOperationResult) OperationKind() string                    { return "delegation" }
func (r *DelegationInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}

type SetDepositsLimitResultContents = ConsumedGasResultContents

//json:kind=OperationKind()
type SetDepositsLimitSuccessfulManagerResult struct {
	core.OperationResultApplied[*ConsumedGasResultContents]
}

func (*SetDepositsLimitSuccessfulManagerResult) OperationKind() string { return "set_deposits_limit" }

type ConsumedGasResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ConsumedGasResult]{
		Variants: encoding.Variants[ConsumedGasResult]{
			0: (*core.OperationResultApplied[*ConsumedGasResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*ConsumedGasResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type IncreasePaidStorage struct {
	ManagerOperation
	Amount      mv.BigInt                 `json:"amount"`
	Destination core.OriginatedContractID `json:"destination"`
}

func (*IncreasePaidStorage) OperationKind() string { return "increase_paid_storage" }

type IncreasePaidStorageResultContents struct {
	BalanceUpdates
	ConsumedMilligas mv.BigUint `json:"consumed_milligas"`
}

func (r *IncreasePaidStorageResultContents) GetConsumedMilligas() mv.BigUint {
	return r.ConsumedMilligas
}

//json:kind=OperationKind()
type IncreasePaidStorageSuccessfulManagerResult struct {
	core.OperationResultApplied[*IncreasePaidStorageResultContents]
}

func (*IncreasePaidStorageSuccessfulManagerResult) OperationKind() string {
	return "increase_paid_storage"
}

type IncreasePaidStorageResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[IncreasePaidStorageResult]{
		Variants: encoding.Variants[IncreasePaidStorageResult]{
			0: (*core.OperationResultApplied[*IncreasePaidStorageResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*IncreasePaidStorageResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type IncreasePaidStorageContentsAndResult struct {
	IncreasePaidStorage
	Metadata ManagerMetadata[IncreasePaidStorageResult] `json:"metadata"`
}

func (*IncreasePaidStorageContentsAndResult) OperationContentsAndResult() {}
func (op *IncreasePaidStorageContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type VDFRevelation struct {
	Solution [2]*[100]byte `json:"solution"`
}

func (*VDFRevelation) OperationKind() string { return "vdf_revelation" }

//json:kind=OperationKind()
type DALSlotAvailability struct {
	Endorser    mv.PublicKeyHash `json:"endorser"`
	Endorsement mv.BigUint       `json:"endorsement"`
}

func (*DALSlotAvailability) OperationKind() string { return "dal_slot_availability" }

//json:kind=OperationKind()
type DALSlotAvailabilityContentsAndResult struct {
	DALSlotAvailability
	Metadata DALSlotAvailabilityMetadata `json:"metadata"`
}

func (*DALSlotAvailabilityContentsAndResult) OperationContentsAndResult() {}
func (op *DALSlotAvailabilityContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type DALSlotAvailabilityMetadata struct {
	Delegate mv.PublicKeyHash `json:"delegate"`
}

//json:kind=OperationKind()
type DALPublishSlotHeader struct {
	ManagerOperation
	Slot DALSlot `json:"slot"`
}

func (*DALPublishSlotHeader) OperationKind() string {
	return "dal_publish_slot_header"
}

type DALSlot struct {
	Level  int32 `json:"level"`
	Index  uint8 `json:"index"`
	Header int32 `json:"header"`
}

//json:kind=OperationKind()
type SeedNonceRevelationContentsAndResult struct {
	SeedNonceRevelation
	Metadata BalanceUpdates `json:"metadata"`
}

func (*SeedNonceRevelationContentsAndResult) OperationContentsAndResult() {}
func (op *SeedNonceRevelationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type DoubleEndorsementEvidenceContentsAndResult struct {
	DoubleEndorsementEvidence
	Metadata BalanceUpdates `json:"metadata"`
}

func (*DoubleEndorsementEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleEndorsementEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type DoubleBakingEvidenceContentsAndResult struct {
	DoubleBakingEvidence
	Metadata BalanceUpdates `json:"metadata"`
}

func (*DoubleBakingEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleBakingEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type ActivateAccountContentsAndResult struct {
	ActivateAccount
	Metadata BalanceUpdates `json:"metadata"`
}

func (*ActivateAccountContentsAndResult) OperationContentsAndResult() {}
func (op *ActivateAccountContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type DoublePreendorsementEvidenceContentsAndResult struct {
	DoublePreendorsementEvidence
	Metadata BalanceUpdates `json:"metadata"`
}

func (*DoublePreendorsementEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoublePreendorsementEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type VDFRevelationContentsAndResult struct {
	VDFRevelation
	Metadata BalanceUpdates `json:"metadata"`
}

func (*VDFRevelationContentsAndResult) OperationContentsAndResult() {}
func (op *VDFRevelationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type EndorsementMetadata struct {
	BalanceUpdates
	Delegate         mv.PublicKeyHash `json:"delegate"`
	EndorsementPower int32            `json:"endorsement_power"`
}

//json:kind=OperationKind()
type EndorsementContentsAndResult struct {
	Endorsement
	Metadata EndorsementMetadata `json:"metadata"`
}

func (*EndorsementContentsAndResult) OperationContentsAndResult() {}
func (op *EndorsementContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type PreendorsementMetadata = EndorsementMetadata

//json:kind=OperationKind()
type PreendorsementContentsAndResult struct {
	Preendorsement
	Metadata PreendorsementMetadata `json:"metadata"`
}

func (*PreendorsementContentsAndResult) OperationContentsAndResult() {}
func (op *PreendorsementContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type RevealContentsAndResult struct {
	Reveal
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}
func (op *RevealContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type DelegationContentsAndResult struct {
	Delegation
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*DelegationContentsAndResult) OperationContentsAndResult() {}
func (op *DelegationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type RegisterGlobalConstantResultContents struct {
	BalanceUpdates
	ConsumedMilligas mv.BigUint         `json:"consumed_milligas"`
	StorageSize      mv.BigInt          `json:"storage_size"`
	GlobalAddress    *mv.ScriptExprHash `json:"global_address"`
}

func (r *RegisterGlobalConstantResultContents) GetConsumedMilligas() mv.BigUint {
	return r.ConsumedMilligas
}
func (r *RegisterGlobalConstantResultContents) GetStorageSize() mv.BigInt { return r.StorageSize }
func (r *RegisterGlobalConstantResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.StorageSize.Int()
}

type RegisterGlobalConstantResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RegisterGlobalConstantResult]{
		Variants: encoding.Variants[RegisterGlobalConstantResult]{
			0: (*core.OperationResultApplied[*RegisterGlobalConstantResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*RegisterGlobalConstantResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type RegisterGlobalConstantContentsAndResult struct {
	RegisterGlobalConstant
	Metadata ManagerMetadata[RegisterGlobalConstantResult] `json:"metadata"`
}

func (*RegisterGlobalConstantContentsAndResult) OperationContentsAndResult() {}
func (op *RegisterGlobalConstantContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SetDepositsLimitContentsAndResult struct {
	SetDepositsLimit
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}
func (op *SetDepositsLimitContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type DALPublishSlotHeaderContentsAndResult struct {
	DALPublishSlotHeader
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*DALPublishSlotHeaderContentsAndResult) OperationContentsAndResult() {}
func (op *DALPublishSlotHeaderContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type TransferTicketContentsAndResult struct {
	TransferTicket
	Metadata ManagerMetadata[TransferTicketResult] `json:"metadata"`
}

func (*TransferTicketContentsAndResult) OperationContentsAndResult() {}
func (op *TransferTicketContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type TransferTicketResultContents struct {
	BalanceUpdates
	ConsumedMilligas    mv.BigUint `json:"consumed_milligas"`
	PaidStorageSizeDiff mv.BigInt  `json:"paid_storage_size_diff"`
}

func (r *TransferTicketResultContents) GetConsumedMilligas() mv.BigUint { return r.ConsumedMilligas }
func (r *TransferTicketResultContents) GetPaidStorageSizeDiff() mv.BigInt {
	return r.PaidStorageSizeDiff
}
func (r *TransferTicketResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.PaidStorageSizeDiff.Int()
}

type TransferTicketResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransferTicketResult]{
		Variants: encoding.Variants[TransferTicketResult]{
			0: (*core.OperationResultApplied[*TransferTicketResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*TransferTicketResultContents])(nil),
		},
	})
}

type OperationContents interface {
	core.OperationContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationContents]{
		Variants: encoding.Variants[OperationContents]{
			1:   (*SeedNonceRevelation)(nil),
			2:   (*DoubleEndorsementEvidence)(nil),
			3:   (*DoubleBakingEvidence)(nil),
			4:   (*ActivateAccount)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			7:   (*DoublePreendorsementEvidence)(nil),
			8:   (*VDFRevelation)(nil),
			17:  (*FailingNoop)(nil),
			20:  (*Preendorsement)(nil),
			21:  (*Endorsement)(nil),
			22:  (*DALSlotAvailability)(nil),
			107: (*Reveal)(nil),
			108: (*Transaction)(nil),
			109: (*Origination)(nil),
			110: (*Delegation)(nil),
			111: (*RegisterGlobalConstant)(nil),
			112: (*SetDepositsLimit)(nil),
			113: (*IncreasePaidStorage)(nil),
			// Never used and deprecated
			// 150 Tx_rollup_origination
			// 151 Tx_rollup_submit_batch
			// 152 Tx_rollup_commit
			// 153 Tx_rollup_return_bond
			// 154 Tx_rollup_finalize_commitment
			// 155 Tx_rollup_remove_commitment
			// 156 Tx_rollup_rejection
			// 157 Tx_rollup_dispatch_tickets
			158: (*TransferTicket)(nil),
			// Never used in this revision
			// 200: (*ScRollupOriginate)(nil),
			// 201: (*ScRollupAddMessages)(nil),
			// 202: (*ScRollupCement)(nil),
			// 203: (*ScRollupPublish)(nil),
			// 204: (*ScRollupRefute)(nil),
			// 205: (*ScRollupTimeout)(nil),
			// 206: (*ScRollupExecuteOutboxMessage)(nil),
			// 207: (*ScRollupRecoverBond)(nil),
			// 208: (*ScRollupDALSlotSubscribe)(nil),
			230: (*DALPublishSlotHeader)(nil),
		},
	})
}

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationContentsAndResult]{
		Variants: encoding.Variants[OperationContentsAndResult]{
			1:   (*SeedNonceRevelationContentsAndResult)(nil),
			2:   (*DoubleEndorsementEvidenceContentsAndResult)(nil),
			3:   (*DoubleBakingEvidenceContentsAndResult)(nil),
			4:   (*ActivateAccountContentsAndResult)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			7:   (*DoublePreendorsementEvidenceContentsAndResult)(nil),
			8:   (*VDFRevelationContentsAndResult)(nil),
			20:  (*PreendorsementContentsAndResult)(nil),
			21:  (*EndorsementContentsAndResult)(nil),
			22:  (*DALSlotAvailabilityContentsAndResult)(nil),
			107: (*RevealContentsAndResult)(nil),
			108: (*TransactionContentsAndResult)(nil),
			109: (*OriginationContentsAndResult)(nil),
			110: (*DelegationContentsAndResult)(nil),
			111: (*RegisterGlobalConstantContentsAndResult)(nil),
			112: (*SetDepositsLimitContentsAndResult)(nil),
			113: (*IncreasePaidStorageContentsAndResult)(nil),
			158: (*TransferTicketContentsAndResult)(nil),
			230: (*DALPublishSlotHeaderContentsAndResult)(nil),
		},
	})
}

type ManagerMetadata[T core.ManagerOperationResult] struct {
	BalanceUpdates
	OperationResult          T                         `json:"operation_result"`
	InternalOperationResults []InternalOperationResult `mv:"dyn" json:"internal_operation_results"`
}

func (m *ManagerMetadata[T]) GetResult() core.ManagerOperationResult {
	return m.OperationResult
}
func (m *ManagerMetadata[T]) GetInternalOperationResults() []core.InternalOperationResult {
	out := make([]core.InternalOperationResult, len(m.InternalOperationResults))
	for i, r := range m.InternalOperationResults {
		out[i] = r
	}
	return out
}

//json:kind=OperationKind()
type EventInternalOperationResult struct {
	Source  core.TransactionDestination      `json:"source"`
	Nonce   uint16                           `json:"nonce"`
	Type    expression.Expression            `json:"type"`
	Tag     mv.Option[Entrypoint]            `json:"tag"`
	Payload mv.Option[expression.Expression] `json:"payload"`
	Result  ConsumedGasResult                `json:"result"`
}

var _ core.InternalOperationResult = (*EventInternalOperationResult)(nil)

func (r *EventInternalOperationResult) GetSource() core.TransactionDestination { return r.Source }
func (r *EventInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*EventInternalOperationResult) OperationKind() string { return "event" }

type InternalOperationResult interface {
	core.InternalOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InternalOperationResult]{
		Variants: encoding.Variants[InternalOperationResult]{
			1: (*TransactionInternalOperationResult)(nil),
			2: (*OriginationInternalOperationResult)(nil),
			3: (*DelegationInternalOperationResult)(nil),
			4: (*EventInternalOperationResult)(nil),
		},
	})
}

type SuccessfulManagerOperationResult interface {
	core.SuccessfulManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SuccessfulManagerOperationResult]{
		Variants: encoding.Variants[SuccessfulManagerOperationResult]{
			0: (*RevealSuccessfulManagerResult)(nil),
			1: (*TransactionSuccessfulManagerResult)(nil),
			2: (*OriginationSuccessfulManagerResult)(nil),
			3: (*DelegationSuccessfulManagerResult)(nil),
			5: (*SetDepositsLimitSuccessfulManagerResult)(nil),
			9: (*IncreasePaidStorageSuccessfulManagerResult)(nil),
			// 200: (*ScRollupOriginateResultContents)(nil),
		},
	})
}

func ListOperations() []OperationContents {
	return encoding.ListVariants[OperationContents]()
}

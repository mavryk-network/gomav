package proto_012_Psithaca

//go:generate go run ../../cmd/genmarshaller.go

import (
	"math/big"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/core/expression"
)

//json:kind=OperationKind()
type ActivateAccount struct {
	PKH    *mv.Ed25519PublicKeyHash `json:"pkh"`
	Secret *mv.Bytes20              `json:"secret"`
}

func (*ActivateAccount) OperationKind() string { return "activate_account" }

//json:kind=OperationKind()
type Proposals struct {
	Source    mv.PublicKeyHash   `json:"source"`
	Period    int32              `json:"period"`
	Proposals []*mv.ProtocolHash `mv:"dyn" json:"proposals"`
}

var _ core.OperationWithSource = (*Proposals)(nil)

func (*Proposals) OperationKind() string { return "proposals" }
func (p *Proposals) GetSource() core.TransactionDestination {
	return core.ImplicitContract{PublicKeyHash: p.Source}
}
func (*Proposals) OperationContentsAndResult() {}
func (op *Proposals) GetMetadata() any         { return op }

//json:kind=OperationKind()
type Ballot struct {
	Source   mv.PublicKeyHash `json:"source"`
	Period   int32            `json:"period"`
	Proposal *mv.ProtocolHash `json:"proposal"`
	Ballot   core.BallotKind  `json:"ballot"`
}

var _ core.OperationWithSource = (*Ballot)(nil)

func (op *Ballot) GetSource() core.TransactionDestination {
	return core.ImplicitContract{PublicKeyHash: op.Source}
}
func (*Ballot) OperationKind() string       { return "ballot" }
func (*Ballot) OperationContentsAndResult() {}
func (op *Ballot) GetMetadata() any         { return op }

type ManagerOperation struct {
	Source       mv.PublicKeyHash `json:"source"`
	Fee          mv.BigUint       `json:"fee"`
	Counter      mv.BigUint       `json:"counter"`
	GasLimit     mv.BigUint       `json:"gas_limit"`
	StorageLimit mv.BigUint       `json:"storage_limit"`
}

func (m *ManagerOperation) GetSource() core.TransactionDestination {
	return core.ImplicitContract{PublicKeyHash: m.Source}
}
func (m *ManagerOperation) GetSourceAddress() mv.PublicKeyHash { return m.Source }
func (m *ManagerOperation) GetFee() mv.BigUint                 { return m.Fee }
func (m *ManagerOperation) GetCounter() mv.BigUint             { return m.Counter }
func (m *ManagerOperation) GetGasLimit() mv.BigUint            { return m.GasLimit }
func (m *ManagerOperation) GetStorageLimit() mv.BigUint        { return m.StorageLimit }

func (m *ManagerOperation) SetFee(v mv.BigUint)          { m.Fee = v }
func (m *ManagerOperation) SetCounter(v mv.BigUint)      { m.Counter = v }
func (m *ManagerOperation) SetGasLimit(v mv.BigUint)     { m.GasLimit = v }
func (m *ManagerOperation) SetStorageLimit(v mv.BigUint) { m.StorageLimit = v }

type Script struct {
	Code    expression.Expression `mv:"dyn" json:"code"`
	Storage expression.Expression `mv:"dyn" json:"storage"`
}

//json:kind=OperationKind()
type Delegation struct {
	ManagerOperation
	Delegate mv.Option[mv.PublicKeyHash] `json:"delegate"`
}

func (*Delegation) OperationKind() string { return "delegation" }

//json:kind=OperationKind()
type Reveal struct {
	ManagerOperation
	PublicKey mv.PublicKey `json:"public_key"`
}

func (*Reveal) OperationKind() string { return "reveal" }

//json:kind=OperationKind()
type SeedNonceRevelation struct {
	Level int32         `json:"level"`
	Nonce *mv.SeedNonce `json:"nonce"`
}

func (*SeedNonceRevelation) OperationKind() string { return "seed_nonce_revelation" }

//json:kind=OperationKind()
type FailingNoop struct {
	Arbitrary mv.Bytes `mv:"dyn" json:"arbitrary"`
}

func (*FailingNoop) OperationKind() string { return "failing_noop" }

//json:kind=OperationKind()
type RegisterGlobalConstant struct {
	ManagerOperation
	Value expression.Expression `mv:"dyn" json:"value"`
}

func (*RegisterGlobalConstant) OperationKind() string { return "register_global_constant" }

//json:kind=OperationKind()
type SetDepositsLimit struct {
	ManagerOperation
	Limit mv.Option[mv.BigUint] `json:"limit"`
}

func (*SetDepositsLimit) OperationKind() string { return "set_deposits_limit" }

//json:kind=OperationKind()
type Endorsement struct {
	Slot             uint16               `json:"slot"`
	Level            int32                `json:"level"`
	Round            int32                `json:"round"`
	BlockPayloadHash *mv.BlockPayloadHash `json:"block_payload_hash"`
}

func (*Endorsement) InlinedEndorsementContents() {}
func (*Endorsement) OperationKind() string       { return "endorsement" }

type InlinedEndorsement struct {
	Branch    *mv.BlockHash              `json:"branch"`
	Contents  InlinedEndorsementContents `json:"contents"`
	Signature mv.AnySignature            `json:"signature"`
}

type InlinedEndorsementContents interface {
	InlinedEndorsementContents()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InlinedEndorsementContents]{
		Variants: encoding.Variants[InlinedEndorsementContents]{
			21: (*Endorsement)(nil),
		},
	})
}

//json:kind=OperationKind()
type DoubleEndorsementEvidence struct {
	Op1 InlinedEndorsement `mv:"dyn" json:"op1"`
	Op2 InlinedEndorsement `mv:"dyn" json:"op2"`
}

func (*DoubleEndorsementEvidence) OperationKind() string { return "double_endorsement_evidence" }

//json:kind=OperationKind()
type DoublePreendorsementEvidence struct {
	Op1 InlinedPreendorsement `mv:"dyn" json:"op1"`
	Op2 InlinedPreendorsement `mv:"dyn" json:"op2"`
}

func (*DoublePreendorsementEvidence) OperationKind() string { return "double_preendorsement_evidence" }

type InlinedPreendorsement struct {
	Branch    *mv.BlockHash                 `json:"branch"`
	Contents  InlinedPreendorsementContents `json:"contents"`
	Signature *mv.GenericSignature          `json:"signature"`
}

type InlinedPreendorsementContents interface {
	InlinedPreendorsementContents()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InlinedPreendorsementContents]{
		Variants: encoding.Variants[InlinedPreendorsementContents]{
			20: (*Preendorsement)(nil),
		},
	})
}

//json:kind=OperationKind()
type Preendorsement Endorsement

func (*Preendorsement) InlinedPreendorsementContents() {}
func (*Preendorsement) OperationKind() string          { return "preendorsement" }

//json:kind=OperationKind()
type DoubleBakingEvidence struct {
	Block1 BlockHeader `mv:"dyn" json:"block1"`
	Block2 BlockHeader `mv:"dyn" json:"block2"`
}

func (*DoubleBakingEvidence) OperationKind() string { return "double_baking_evidence" }

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
			17:  (*FailingNoop)(nil),
			20:  (*Preendorsement)(nil),
			21:  (*Endorsement)(nil),
			107: (*Reveal)(nil),
			108: (*Transaction)(nil),
			109: (*Origination)(nil),
			110: (*Delegation)(nil),
			111: (*RegisterGlobalConstant)(nil),
			112: (*SetDepositsLimit)(nil),
		},
	})
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

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

type ConsumedGasResultContents struct {
	ConsumedGas      mv.BigUint `json:"consumed_gas"`
	ConsumedMilligas mv.BigUint `json:"consumed_milligas"`
}

func (r *ConsumedGasResultContents) GetConsumedMilligas() mv.BigUint { return r.ConsumedMilligas }

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

type RevealResultContents = ConsumedGasResultContents

//json:kind=OperationKind()
type RevealSuccessfulManagerResult struct {
	core.OperationResultApplied[*ConsumedGasResultContents]
}

func (*RevealSuccessfulManagerResult) OperationKind() string { return "reveal" }

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

type DelegationResultContents = ConsumedGasResultContents

//json:kind=OperationKind()
type DelegationSuccessfulManagerResult struct {
	core.OperationResultApplied[*ConsumedGasResultContents]
}

func (*DelegationSuccessfulManagerResult) OperationKind() string { return "delegation" }

type SetDepositsLimitResultContents = ConsumedGasResultContents

//json:kind=OperationKind()
type SetDepositsLimitSuccessfulManagerResult struct {
	core.OperationResultApplied[*ConsumedGasResultContents]
}

func (*SetDepositsLimitSuccessfulManagerResult) OperationKind() string { return "set_deposits_limit" }

type RegisterGlobalConstantResultContents struct {
	BalanceUpdates
	ConsumedGas   mv.BigUint         `json:"consumed_gas"`
	StorageSize   mv.BigInt          `json:"storage_size"`
	GlobalAddress *mv.ScriptExprHash `json:"global_address"`
}

func (r *RegisterGlobalConstantResultContents) GetConsumedMilligas() mv.BigUint {
	x := r.ConsumedGas.Int()
	x.Mul(x, big.NewInt(1000))
	v, _ := mv.NewBigUint(x)
	return v
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
			20:  (*PreendorsementContentsAndResult)(nil),
			21:  (*EndorsementContentsAndResult)(nil),
			107: (*RevealContentsAndResult)(nil),
			108: (*TransactionContentsAndResult)(nil),
			109: (*OriginationContentsAndResult)(nil),
			110: (*DelegationContentsAndResult)(nil),
			111: (*RegisterGlobalConstantContentsAndResult)(nil),
			112: (*SetDepositsLimitContentsAndResult)(nil),
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
type DelegationInternalOperationResult struct {
	Source   core.ContractID             `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Delegate mv.Option[mv.PublicKeyHash] `json:"delegate"`
	Result   ConsumedGasResult           `json:"result"`
}

var _ core.InternalOperationResult = (*DelegationInternalOperationResult)(nil)

func (r *DelegationInternalOperationResult) GetSource() core.TransactionDestination {
	switch d := r.Source.(type) {
	case core.ImplicitContract:
		return d
	case core.OriginatedContract:
		return d
	default:
		panic("unexpected contract id type")
	}
}
func (r *DelegationInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*DelegationInternalOperationResult) OperationKind() string { return "delegation" }

//json:kind=OperationKind()
type RevealInternalOperationResult struct {
	Source    core.ContractID   `json:"source"`
	Nonce     uint16            `json:"nonce"`
	PublicKey mv.PublicKey      `json:"public_key"`
	Result    ConsumedGasResult `json:"result"`
}

var _ core.InternalOperationResult = (*RevealInternalOperationResult)(nil)

func (r *RevealInternalOperationResult) GetSource() core.TransactionDestination {
	switch d := r.Source.(type) {
	case core.ImplicitContract:
		return d
	case core.OriginatedContract:
		return d
	default:
		panic("unexpected contract id type")
	}
}
func (r *RevealInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*RevealInternalOperationResult) OperationKind() string { return "reveal" }

//json:kind=OperationKind()
type RegisterGlobalConstantInternalOperationResult struct {
	Source core.ContractID              `json:"source"`
	Nonce  uint16                       `json:"nonce"`
	Value  expression.Expression        `mv:"dyn" json:"value"`
	Result RegisterGlobalConstantResult `json:"result"`
}

var _ core.InternalOperationResult = (*RegisterGlobalConstantInternalOperationResult)(nil)

func (r *RegisterGlobalConstantInternalOperationResult) GetSource() core.TransactionDestination {
	switch d := r.Source.(type) {
	case core.ImplicitContract:
		return d
	case core.OriginatedContract:
		return d
	default:
		panic("unexpected contract id type")
	}
}

func (*RegisterGlobalConstantInternalOperationResult) OperationKind() string {
	return "register_global_constant"
}
func (r *RegisterGlobalConstantInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}

//json:kind=OperationKind()
type SetDepositsLimitInternalOperationResult struct {
	Source core.ContractID       `json:"source"`
	Nonce  uint16                `json:"nonce"`
	Limit  mv.Option[mv.BigUint] `json:"limit"`
	Result ConsumedGasResult     `json:"result"`
}

var _ core.InternalOperationResult = (*SetDepositsLimitInternalOperationResult)(nil)

func (r *SetDepositsLimitInternalOperationResult) GetSource() core.TransactionDestination {
	switch d := r.Source.(type) {
	case core.ImplicitContract:
		return d
	case core.OriginatedContract:
		return d
	default:
		panic("unexpected contract id type")
	}
}
func (*SetDepositsLimitInternalOperationResult) OperationKind() string {
	return "set_deposits_limit"
}
func (r *SetDepositsLimitInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}

type InternalOperationResult interface {
	core.InternalOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InternalOperationResult]{
		Variants: encoding.Variants[InternalOperationResult]{
			0: (*RevealInternalOperationResult)(nil),
			1: (*TransactionInternalOperationResult)(nil),
			2: (*OriginationInternalOperationResult)(nil),
			3: (*DelegationInternalOperationResult)(nil),
			4: (*RegisterGlobalConstantInternalOperationResult)(nil),
			5: (*SetDepositsLimitInternalOperationResult)(nil),
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
		},
	})
}

func ListOperations() []OperationContents {
	return encoding.ListVariants[OperationContents]()
}

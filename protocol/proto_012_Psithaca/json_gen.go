package proto_012_Psithaca

import "encoding/json"

// Code generated by genmarshaller.go DO NOT EDIT.

func (self *ActivateAccount) MarshalJSON() ([]byte, error) {
	type ActivateAccount_no_json_marshaller ActivateAccount

	type json_ActivateAccount struct {
		Marker0 any `json:"kind"`
		ActivateAccount_no_json_marshaller
	}

	tmp := json_ActivateAccount {
		Marker0: self.OperationKind(),
		ActivateAccount_no_json_marshaller: ActivateAccount_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *Proposals) MarshalJSON() ([]byte, error) {
	type Proposals_no_json_marshaller Proposals

	type json_Proposals struct {
		Marker0 any `json:"kind"`
		Proposals_no_json_marshaller
	}

	tmp := json_Proposals {
		Marker0: self.OperationKind(),
		Proposals_no_json_marshaller: Proposals_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *Ballot) MarshalJSON() ([]byte, error) {
	type Ballot_no_json_marshaller Ballot

	type json_Ballot struct {
		Marker0 any `json:"kind"`
		Ballot_no_json_marshaller
	}

	tmp := json_Ballot {
		Marker0: self.OperationKind(),
		Ballot_no_json_marshaller: Ballot_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *Delegation) MarshalJSON() ([]byte, error) {
	type Delegation_no_json_marshaller Delegation

	type json_Delegation struct {
		Marker0 any `json:"kind"`
		Delegation_no_json_marshaller
	}

	tmp := json_Delegation {
		Marker0: self.OperationKind(),
		Delegation_no_json_marshaller: Delegation_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *Reveal) MarshalJSON() ([]byte, error) {
	type Reveal_no_json_marshaller Reveal

	type json_Reveal struct {
		Marker0 any `json:"kind"`
		Reveal_no_json_marshaller
	}

	tmp := json_Reveal {
		Marker0: self.OperationKind(),
		Reveal_no_json_marshaller: Reveal_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SeedNonceRevelation) MarshalJSON() ([]byte, error) {
	type SeedNonceRevelation_no_json_marshaller SeedNonceRevelation

	type json_SeedNonceRevelation struct {
		Marker0 any `json:"kind"`
		SeedNonceRevelation_no_json_marshaller
	}

	tmp := json_SeedNonceRevelation {
		Marker0: self.OperationKind(),
		SeedNonceRevelation_no_json_marshaller: SeedNonceRevelation_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *FailingNoop) MarshalJSON() ([]byte, error) {
	type FailingNoop_no_json_marshaller FailingNoop

	type json_FailingNoop struct {
		Marker0 any `json:"kind"`
		FailingNoop_no_json_marshaller
	}

	tmp := json_FailingNoop {
		Marker0: self.OperationKind(),
		FailingNoop_no_json_marshaller: FailingNoop_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *RegisterGlobalConstant) MarshalJSON() ([]byte, error) {
	type RegisterGlobalConstant_no_json_marshaller RegisterGlobalConstant

	type json_RegisterGlobalConstant struct {
		Marker0 any `json:"kind"`
		RegisterGlobalConstant_no_json_marshaller
	}

	tmp := json_RegisterGlobalConstant {
		Marker0: self.OperationKind(),
		RegisterGlobalConstant_no_json_marshaller: RegisterGlobalConstant_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SetDepositsLimit) MarshalJSON() ([]byte, error) {
	type SetDepositsLimit_no_json_marshaller SetDepositsLimit

	type json_SetDepositsLimit struct {
		Marker0 any `json:"kind"`
		SetDepositsLimit_no_json_marshaller
	}

	tmp := json_SetDepositsLimit {
		Marker0: self.OperationKind(),
		SetDepositsLimit_no_json_marshaller: SetDepositsLimit_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *Endorsement) MarshalJSON() ([]byte, error) {
	type Endorsement_no_json_marshaller Endorsement

	type json_Endorsement struct {
		Marker0 any `json:"kind"`
		Endorsement_no_json_marshaller
	}

	tmp := json_Endorsement {
		Marker0: self.OperationKind(),
		Endorsement_no_json_marshaller: Endorsement_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DoubleEndorsementEvidence) MarshalJSON() ([]byte, error) {
	type DoubleEndorsementEvidence_no_json_marshaller DoubleEndorsementEvidence

	type json_DoubleEndorsementEvidence struct {
		Marker0 any `json:"kind"`
		DoubleEndorsementEvidence_no_json_marshaller
	}

	tmp := json_DoubleEndorsementEvidence {
		Marker0: self.OperationKind(),
		DoubleEndorsementEvidence_no_json_marshaller: DoubleEndorsementEvidence_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DoublePreendorsementEvidence) MarshalJSON() ([]byte, error) {
	type DoublePreendorsementEvidence_no_json_marshaller DoublePreendorsementEvidence

	type json_DoublePreendorsementEvidence struct {
		Marker0 any `json:"kind"`
		DoublePreendorsementEvidence_no_json_marshaller
	}

	tmp := json_DoublePreendorsementEvidence {
		Marker0: self.OperationKind(),
		DoublePreendorsementEvidence_no_json_marshaller: DoublePreendorsementEvidence_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *Preendorsement) MarshalJSON() ([]byte, error) {
	type Preendorsement_no_json_marshaller Preendorsement

	type json_Preendorsement struct {
		Marker0 any `json:"kind"`
		Preendorsement_no_json_marshaller
	}

	tmp := json_Preendorsement {
		Marker0: self.OperationKind(),
		Preendorsement_no_json_marshaller: Preendorsement_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DoubleBakingEvidence) MarshalJSON() ([]byte, error) {
	type DoubleBakingEvidence_no_json_marshaller DoubleBakingEvidence

	type json_DoubleBakingEvidence struct {
		Marker0 any `json:"kind"`
		DoubleBakingEvidence_no_json_marshaller
	}

	tmp := json_DoubleBakingEvidence {
		Marker0: self.OperationKind(),
		DoubleBakingEvidence_no_json_marshaller: DoubleBakingEvidence_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SeedNonceRevelationContentsAndResult) MarshalJSON() ([]byte, error) {
	type SeedNonceRevelationContentsAndResult_no_json_marshaller SeedNonceRevelationContentsAndResult

	type json_SeedNonceRevelationContentsAndResult struct {
		Marker0 any `json:"kind"`
		SeedNonceRevelationContentsAndResult_no_json_marshaller
	}

	tmp := json_SeedNonceRevelationContentsAndResult {
		Marker0: self.OperationKind(),
		SeedNonceRevelationContentsAndResult_no_json_marshaller: SeedNonceRevelationContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DoubleEndorsementEvidenceContentsAndResult) MarshalJSON() ([]byte, error) {
	type DoubleEndorsementEvidenceContentsAndResult_no_json_marshaller DoubleEndorsementEvidenceContentsAndResult

	type json_DoubleEndorsementEvidenceContentsAndResult struct {
		Marker0 any `json:"kind"`
		DoubleEndorsementEvidenceContentsAndResult_no_json_marshaller
	}

	tmp := json_DoubleEndorsementEvidenceContentsAndResult {
		Marker0: self.OperationKind(),
		DoubleEndorsementEvidenceContentsAndResult_no_json_marshaller: DoubleEndorsementEvidenceContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DoubleBakingEvidenceContentsAndResult) MarshalJSON() ([]byte, error) {
	type DoubleBakingEvidenceContentsAndResult_no_json_marshaller DoubleBakingEvidenceContentsAndResult

	type json_DoubleBakingEvidenceContentsAndResult struct {
		Marker0 any `json:"kind"`
		DoubleBakingEvidenceContentsAndResult_no_json_marshaller
	}

	tmp := json_DoubleBakingEvidenceContentsAndResult {
		Marker0: self.OperationKind(),
		DoubleBakingEvidenceContentsAndResult_no_json_marshaller: DoubleBakingEvidenceContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *ActivateAccountContentsAndResult) MarshalJSON() ([]byte, error) {
	type ActivateAccountContentsAndResult_no_json_marshaller ActivateAccountContentsAndResult

	type json_ActivateAccountContentsAndResult struct {
		Marker0 any `json:"kind"`
		ActivateAccountContentsAndResult_no_json_marshaller
	}

	tmp := json_ActivateAccountContentsAndResult {
		Marker0: self.OperationKind(),
		ActivateAccountContentsAndResult_no_json_marshaller: ActivateAccountContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DoublePreendorsementEvidenceContentsAndResult) MarshalJSON() ([]byte, error) {
	type DoublePreendorsementEvidenceContentsAndResult_no_json_marshaller DoublePreendorsementEvidenceContentsAndResult

	type json_DoublePreendorsementEvidenceContentsAndResult struct {
		Marker0 any `json:"kind"`
		DoublePreendorsementEvidenceContentsAndResult_no_json_marshaller
	}

	tmp := json_DoublePreendorsementEvidenceContentsAndResult {
		Marker0: self.OperationKind(),
		DoublePreendorsementEvidenceContentsAndResult_no_json_marshaller: DoublePreendorsementEvidenceContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *EndorsementContentsAndResult) MarshalJSON() ([]byte, error) {
	type EndorsementContentsAndResult_no_json_marshaller EndorsementContentsAndResult

	type json_EndorsementContentsAndResult struct {
		Marker0 any `json:"kind"`
		EndorsementContentsAndResult_no_json_marshaller
	}

	tmp := json_EndorsementContentsAndResult {
		Marker0: self.OperationKind(),
		EndorsementContentsAndResult_no_json_marshaller: EndorsementContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *PreendorsementContentsAndResult) MarshalJSON() ([]byte, error) {
	type PreendorsementContentsAndResult_no_json_marshaller PreendorsementContentsAndResult

	type json_PreendorsementContentsAndResult struct {
		Marker0 any `json:"kind"`
		PreendorsementContentsAndResult_no_json_marshaller
	}

	tmp := json_PreendorsementContentsAndResult {
		Marker0: self.OperationKind(),
		PreendorsementContentsAndResult_no_json_marshaller: PreendorsementContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *RevealSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type RevealSuccessfulManagerResult_no_json_marshaller RevealSuccessfulManagerResult

	type json_RevealSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		RevealSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_RevealSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		RevealSuccessfulManagerResult_no_json_marshaller: RevealSuccessfulManagerResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *RevealContentsAndResult) MarshalJSON() ([]byte, error) {
	type RevealContentsAndResult_no_json_marshaller RevealContentsAndResult

	type json_RevealContentsAndResult struct {
		Marker0 any `json:"kind"`
		RevealContentsAndResult_no_json_marshaller
	}

	tmp := json_RevealContentsAndResult {
		Marker0: self.OperationKind(),
		RevealContentsAndResult_no_json_marshaller: RevealContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DelegationContentsAndResult) MarshalJSON() ([]byte, error) {
	type DelegationContentsAndResult_no_json_marshaller DelegationContentsAndResult

	type json_DelegationContentsAndResult struct {
		Marker0 any `json:"kind"`
		DelegationContentsAndResult_no_json_marshaller
	}

	tmp := json_DelegationContentsAndResult {
		Marker0: self.OperationKind(),
		DelegationContentsAndResult_no_json_marshaller: DelegationContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DelegationSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type DelegationSuccessfulManagerResult_no_json_marshaller DelegationSuccessfulManagerResult

	type json_DelegationSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		DelegationSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_DelegationSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		DelegationSuccessfulManagerResult_no_json_marshaller: DelegationSuccessfulManagerResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SetDepositsLimitSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type SetDepositsLimitSuccessfulManagerResult_no_json_marshaller SetDepositsLimitSuccessfulManagerResult

	type json_SetDepositsLimitSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		SetDepositsLimitSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_SetDepositsLimitSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		SetDepositsLimitSuccessfulManagerResult_no_json_marshaller: SetDepositsLimitSuccessfulManagerResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *RegisterGlobalConstantContentsAndResult) MarshalJSON() ([]byte, error) {
	type RegisterGlobalConstantContentsAndResult_no_json_marshaller RegisterGlobalConstantContentsAndResult

	type json_RegisterGlobalConstantContentsAndResult struct {
		Marker0 any `json:"kind"`
		RegisterGlobalConstantContentsAndResult_no_json_marshaller
	}

	tmp := json_RegisterGlobalConstantContentsAndResult {
		Marker0: self.OperationKind(),
		RegisterGlobalConstantContentsAndResult_no_json_marshaller: RegisterGlobalConstantContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SetDepositsLimitContentsAndResult) MarshalJSON() ([]byte, error) {
	type SetDepositsLimitContentsAndResult_no_json_marshaller SetDepositsLimitContentsAndResult

	type json_SetDepositsLimitContentsAndResult struct {
		Marker0 any `json:"kind"`
		SetDepositsLimitContentsAndResult_no_json_marshaller
	}

	tmp := json_SetDepositsLimitContentsAndResult {
		Marker0: self.OperationKind(),
		SetDepositsLimitContentsAndResult_no_json_marshaller: SetDepositsLimitContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DelegationInternalOperationResult) MarshalJSON() ([]byte, error) {
	type DelegationInternalOperationResult_no_json_marshaller DelegationInternalOperationResult

	type json_DelegationInternalOperationResult struct {
		Marker0 any `json:"kind"`
		DelegationInternalOperationResult_no_json_marshaller
	}

	tmp := json_DelegationInternalOperationResult {
		Marker0: self.OperationKind(),
		DelegationInternalOperationResult_no_json_marshaller: DelegationInternalOperationResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *RevealInternalOperationResult) MarshalJSON() ([]byte, error) {
	type RevealInternalOperationResult_no_json_marshaller RevealInternalOperationResult

	type json_RevealInternalOperationResult struct {
		Marker0 any `json:"kind"`
		RevealInternalOperationResult_no_json_marshaller
	}

	tmp := json_RevealInternalOperationResult {
		Marker0: self.OperationKind(),
		RevealInternalOperationResult_no_json_marshaller: RevealInternalOperationResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *RegisterGlobalConstantInternalOperationResult) MarshalJSON() ([]byte, error) {
	type RegisterGlobalConstantInternalOperationResult_no_json_marshaller RegisterGlobalConstantInternalOperationResult

	type json_RegisterGlobalConstantInternalOperationResult struct {
		Marker0 any `json:"kind"`
		RegisterGlobalConstantInternalOperationResult_no_json_marshaller
	}

	tmp := json_RegisterGlobalConstantInternalOperationResult {
		Marker0: self.OperationKind(),
		RegisterGlobalConstantInternalOperationResult_no_json_marshaller: RegisterGlobalConstantInternalOperationResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SetDepositsLimitInternalOperationResult) MarshalJSON() ([]byte, error) {
	type SetDepositsLimitInternalOperationResult_no_json_marshaller SetDepositsLimitInternalOperationResult

	type json_SetDepositsLimitInternalOperationResult struct {
		Marker0 any `json:"kind"`
		SetDepositsLimitInternalOperationResult_no_json_marshaller
	}

	tmp := json_SetDepositsLimitInternalOperationResult {
		Marker0: self.OperationKind(),
		SetDepositsLimitInternalOperationResult_no_json_marshaller: SetDepositsLimitInternalOperationResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *Origination) MarshalJSON() ([]byte, error) {
	type Origination_no_json_marshaller Origination

	type json_Origination struct {
		Marker0 any `json:"kind"`
		Origination_no_json_marshaller
	}

	tmp := json_Origination {
		Marker0: self.OperationKind(),
		Origination_no_json_marshaller: Origination_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *OriginationSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type OriginationSuccessfulManagerResult_no_json_marshaller OriginationSuccessfulManagerResult

	type json_OriginationSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		OriginationSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_OriginationSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		OriginationSuccessfulManagerResult_no_json_marshaller: OriginationSuccessfulManagerResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *OriginationContentsAndResult) MarshalJSON() ([]byte, error) {
	type OriginationContentsAndResult_no_json_marshaller OriginationContentsAndResult

	type json_OriginationContentsAndResult struct {
		Marker0 any `json:"kind"`
		OriginationContentsAndResult_no_json_marshaller
	}

	tmp := json_OriginationContentsAndResult {
		Marker0: self.OperationKind(),
		OriginationContentsAndResult_no_json_marshaller: OriginationContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *OriginationInternalOperationResult) MarshalJSON() ([]byte, error) {
	type OriginationInternalOperationResult_no_json_marshaller OriginationInternalOperationResult

	type json_OriginationInternalOperationResult struct {
		Marker0 any `json:"kind"`
		OriginationInternalOperationResult_no_json_marshaller
	}

	tmp := json_OriginationInternalOperationResult {
		Marker0: self.OperationKind(),
		OriginationInternalOperationResult_no_json_marshaller: OriginationInternalOperationResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateContract) MarshalJSON() ([]byte, error) {
	type BalanceUpdateContract_no_json_marshaller BalanceUpdateContract

	type json_BalanceUpdateContract struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateContract_no_json_marshaller
	}

	tmp := json_BalanceUpdateContract {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateContract_no_json_marshaller: BalanceUpdateContract_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateBlockFees) MarshalJSON() ([]byte, error) {
	type BalanceUpdateBlockFees_no_json_marshaller BalanceUpdateBlockFees

	type json_BalanceUpdateBlockFees struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateBlockFees_no_json_marshaller
	}

	tmp := json_BalanceUpdateBlockFees {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateBlockFees_no_json_marshaller: BalanceUpdateBlockFees_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateDeposits) MarshalJSON() ([]byte, error) {
	type BalanceUpdateDeposits_no_json_marshaller BalanceUpdateDeposits

	type json_BalanceUpdateDeposits struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateDeposits_no_json_marshaller
	}

	tmp := json_BalanceUpdateDeposits {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateDeposits_no_json_marshaller: BalanceUpdateDeposits_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateNonceRevelationRewards) MarshalJSON() ([]byte, error) {
	type BalanceUpdateNonceRevelationRewards_no_json_marshaller BalanceUpdateNonceRevelationRewards

	type json_BalanceUpdateNonceRevelationRewards struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateNonceRevelationRewards_no_json_marshaller
	}

	tmp := json_BalanceUpdateNonceRevelationRewards {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateNonceRevelationRewards_no_json_marshaller: BalanceUpdateNonceRevelationRewards_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateDoubleSigningEvidenceRewards) MarshalJSON() ([]byte, error) {
	type BalanceUpdateDoubleSigningEvidenceRewards_no_json_marshaller BalanceUpdateDoubleSigningEvidenceRewards

	type json_BalanceUpdateDoubleSigningEvidenceRewards struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateDoubleSigningEvidenceRewards_no_json_marshaller
	}

	tmp := json_BalanceUpdateDoubleSigningEvidenceRewards {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateDoubleSigningEvidenceRewards_no_json_marshaller: BalanceUpdateDoubleSigningEvidenceRewards_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateEndorsingRewards) MarshalJSON() ([]byte, error) {
	type BalanceUpdateEndorsingRewards_no_json_marshaller BalanceUpdateEndorsingRewards

	type json_BalanceUpdateEndorsingRewards struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateEndorsingRewards_no_json_marshaller
	}

	tmp := json_BalanceUpdateEndorsingRewards {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateEndorsingRewards_no_json_marshaller: BalanceUpdateEndorsingRewards_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateBakingRewards) MarshalJSON() ([]byte, error) {
	type BalanceUpdateBakingRewards_no_json_marshaller BalanceUpdateBakingRewards

	type json_BalanceUpdateBakingRewards struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateBakingRewards_no_json_marshaller
	}

	tmp := json_BalanceUpdateBakingRewards {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateBakingRewards_no_json_marshaller: BalanceUpdateBakingRewards_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateBakingBonuses) MarshalJSON() ([]byte, error) {
	type BalanceUpdateBakingBonuses_no_json_marshaller BalanceUpdateBakingBonuses

	type json_BalanceUpdateBakingBonuses struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateBakingBonuses_no_json_marshaller
	}

	tmp := json_BalanceUpdateBakingBonuses {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateBakingBonuses_no_json_marshaller: BalanceUpdateBakingBonuses_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateStorageFees) MarshalJSON() ([]byte, error) {
	type BalanceUpdateStorageFees_no_json_marshaller BalanceUpdateStorageFees

	type json_BalanceUpdateStorageFees struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateStorageFees_no_json_marshaller
	}

	tmp := json_BalanceUpdateStorageFees {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateStorageFees_no_json_marshaller: BalanceUpdateStorageFees_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateDoubleSigningPunishments) MarshalJSON() ([]byte, error) {
	type BalanceUpdateDoubleSigningPunishments_no_json_marshaller BalanceUpdateDoubleSigningPunishments

	type json_BalanceUpdateDoubleSigningPunishments struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateDoubleSigningPunishments_no_json_marshaller
	}

	tmp := json_BalanceUpdateDoubleSigningPunishments {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateDoubleSigningPunishments_no_json_marshaller: BalanceUpdateDoubleSigningPunishments_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateLiquidityBakingSubsidies) MarshalJSON() ([]byte, error) {
	type BalanceUpdateLiquidityBakingSubsidies_no_json_marshaller BalanceUpdateLiquidityBakingSubsidies

	type json_BalanceUpdateLiquidityBakingSubsidies struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateLiquidityBakingSubsidies_no_json_marshaller
	}

	tmp := json_BalanceUpdateLiquidityBakingSubsidies {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateLiquidityBakingSubsidies_no_json_marshaller: BalanceUpdateLiquidityBakingSubsidies_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateBurned) MarshalJSON() ([]byte, error) {
	type BalanceUpdateBurned_no_json_marshaller BalanceUpdateBurned

	type json_BalanceUpdateBurned struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateBurned_no_json_marshaller
	}

	tmp := json_BalanceUpdateBurned {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateBurned_no_json_marshaller: BalanceUpdateBurned_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateBootstrap) MarshalJSON() ([]byte, error) {
	type BalanceUpdateBootstrap_no_json_marshaller BalanceUpdateBootstrap

	type json_BalanceUpdateBootstrap struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateBootstrap_no_json_marshaller
	}

	tmp := json_BalanceUpdateBootstrap {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateBootstrap_no_json_marshaller: BalanceUpdateBootstrap_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateInvoice) MarshalJSON() ([]byte, error) {
	type BalanceUpdateInvoice_no_json_marshaller BalanceUpdateInvoice

	type json_BalanceUpdateInvoice struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateInvoice_no_json_marshaller
	}

	tmp := json_BalanceUpdateInvoice {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateInvoice_no_json_marshaller: BalanceUpdateInvoice_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateInitialCommitments) MarshalJSON() ([]byte, error) {
	type BalanceUpdateInitialCommitments_no_json_marshaller BalanceUpdateInitialCommitments

	type json_BalanceUpdateInitialCommitments struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateInitialCommitments_no_json_marshaller
	}

	tmp := json_BalanceUpdateInitialCommitments {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateInitialCommitments_no_json_marshaller: BalanceUpdateInitialCommitments_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateMinted) MarshalJSON() ([]byte, error) {
	type BalanceUpdateMinted_no_json_marshaller BalanceUpdateMinted

	type json_BalanceUpdateMinted struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateMinted_no_json_marshaller
	}

	tmp := json_BalanceUpdateMinted {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateMinted_no_json_marshaller: BalanceUpdateMinted_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateLostEndorsingRewards) MarshalJSON() ([]byte, error) {
	type BalanceUpdateLostEndorsingRewards_no_json_marshaller BalanceUpdateLostEndorsingRewards

	type json_BalanceUpdateLostEndorsingRewards struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateLostEndorsingRewards_no_json_marshaller
	}

	tmp := json_BalanceUpdateLostEndorsingRewards {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateLostEndorsingRewards_no_json_marshaller: BalanceUpdateLostEndorsingRewards_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateCommitments) MarshalJSON() ([]byte, error) {
	type BalanceUpdateCommitments_no_json_marshaller BalanceUpdateCommitments

	type json_BalanceUpdateCommitments struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateCommitments_no_json_marshaller
	}

	tmp := json_BalanceUpdateCommitments {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateCommitments_no_json_marshaller: BalanceUpdateCommitments_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateLegacyRewards) MarshalJSON() ([]byte, error) {
	type BalanceUpdateLegacyRewards_no_json_marshaller BalanceUpdateLegacyRewards

	type json_BalanceUpdateLegacyRewards struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateLegacyRewards_no_json_marshaller
	}

	tmp := json_BalanceUpdateLegacyRewards {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateLegacyRewards_no_json_marshaller: BalanceUpdateLegacyRewards_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateLegacyDeposits) MarshalJSON() ([]byte, error) {
	type BalanceUpdateLegacyDeposits_no_json_marshaller BalanceUpdateLegacyDeposits

	type json_BalanceUpdateLegacyDeposits struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateLegacyDeposits_no_json_marshaller
	}

	tmp := json_BalanceUpdateLegacyDeposits {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateLegacyDeposits_no_json_marshaller: BalanceUpdateLegacyDeposits_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateLegacyFees) MarshalJSON() ([]byte, error) {
	type BalanceUpdateLegacyFees_no_json_marshaller BalanceUpdateLegacyFees

	type json_BalanceUpdateLegacyFees struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateLegacyFees_no_json_marshaller
	}

	tmp := json_BalanceUpdateLegacyFees {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateLegacyFees_no_json_marshaller: BalanceUpdateLegacyFees_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *Transaction) MarshalJSON() ([]byte, error) {
	type Transaction_no_json_marshaller Transaction

	type json_Transaction struct {
		Marker0 any `json:"kind"`
		Transaction_no_json_marshaller
	}

	tmp := json_Transaction {
		Marker0: self.OperationKind(),
		Transaction_no_json_marshaller: Transaction_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *TransactionSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type TransactionSuccessfulManagerResult_no_json_marshaller TransactionSuccessfulManagerResult

	type json_TransactionSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		TransactionSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_TransactionSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		TransactionSuccessfulManagerResult_no_json_marshaller: TransactionSuccessfulManagerResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *TransactionContentsAndResult) MarshalJSON() ([]byte, error) {
	type TransactionContentsAndResult_no_json_marshaller TransactionContentsAndResult

	type json_TransactionContentsAndResult struct {
		Marker0 any `json:"kind"`
		TransactionContentsAndResult_no_json_marshaller
	}

	tmp := json_TransactionContentsAndResult {
		Marker0: self.OperationKind(),
		TransactionContentsAndResult_no_json_marshaller: TransactionContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *TransactionInternalOperationResult) MarshalJSON() ([]byte, error) {
	type TransactionInternalOperationResult_no_json_marshaller TransactionInternalOperationResult

	type json_TransactionInternalOperationResult struct {
		Marker0 any `json:"kind"`
		TransactionInternalOperationResult_no_json_marshaller
	}

	tmp := json_TransactionInternalOperationResult {
		Marker0: self.OperationKind(),
		TransactionInternalOperationResult_no_json_marshaller: TransactionInternalOperationResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}


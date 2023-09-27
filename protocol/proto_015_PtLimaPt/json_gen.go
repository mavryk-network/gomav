package proto_015_PtLimaPt

import "encoding/json"

// Code generated by genmarshaller.go DO NOT EDIT.

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

func (self *UpdateConsensusKey) MarshalJSON() ([]byte, error) {
	type UpdateConsensusKey_no_json_marshaller UpdateConsensusKey

	type json_UpdateConsensusKey struct {
		Marker0 any `json:"kind"`
		UpdateConsensusKey_no_json_marshaller
	}

	tmp := json_UpdateConsensusKey {
		Marker0: self.OperationKind(),
		UpdateConsensusKey_no_json_marshaller: UpdateConsensusKey_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DrainDelegate) MarshalJSON() ([]byte, error) {
	type DrainDelegate_no_json_marshaller DrainDelegate

	type json_DrainDelegate struct {
		Marker0 any `json:"kind"`
		DrainDelegate_no_json_marshaller
	}

	tmp := json_DrainDelegate {
		Marker0: self.OperationKind(),
		DrainDelegate_no_json_marshaller: DrainDelegate_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *UpdateConsensusKeySuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type UpdateConsensusKeySuccessfulManagerResult_no_json_marshaller UpdateConsensusKeySuccessfulManagerResult

	type json_UpdateConsensusKeySuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		UpdateConsensusKeySuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_UpdateConsensusKeySuccessfulManagerResult {
		Marker0: self.OperationKind(),
		UpdateConsensusKeySuccessfulManagerResult_no_json_marshaller: UpdateConsensusKeySuccessfulManagerResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DALPublishSlotHeader) MarshalJSON() ([]byte, error) {
	type DALPublishSlotHeader_no_json_marshaller DALPublishSlotHeader

	type json_DALPublishSlotHeader struct {
		Marker0 any `json:"kind"`
		DALPublishSlotHeader_no_json_marshaller
	}

	tmp := json_DALPublishSlotHeader {
		Marker0: self.OperationKind(),
		DALPublishSlotHeader_no_json_marshaller: DALPublishSlotHeader_no_json_marshaller(*self),
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

func (self *VDFRevelationContentsAndResult) MarshalJSON() ([]byte, error) {
	type VDFRevelationContentsAndResult_no_json_marshaller VDFRevelationContentsAndResult

	type json_VDFRevelationContentsAndResult struct {
		Marker0 any `json:"kind"`
		VDFRevelationContentsAndResult_no_json_marshaller
	}

	tmp := json_VDFRevelationContentsAndResult {
		Marker0: self.OperationKind(),
		VDFRevelationContentsAndResult_no_json_marshaller: VDFRevelationContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DrainDelegateContentsAndResult) MarshalJSON() ([]byte, error) {
	type DrainDelegateContentsAndResult_no_json_marshaller DrainDelegateContentsAndResult

	type json_DrainDelegateContentsAndResult struct {
		Marker0 any `json:"kind"`
		DrainDelegateContentsAndResult_no_json_marshaller
	}

	tmp := json_DrainDelegateContentsAndResult {
		Marker0: self.OperationKind(),
		DrainDelegateContentsAndResult_no_json_marshaller: DrainDelegateContentsAndResult_no_json_marshaller(*self),
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

func (self *UpdateConsensusKeyContentsAndResult) MarshalJSON() ([]byte, error) {
	type UpdateConsensusKeyContentsAndResult_no_json_marshaller UpdateConsensusKeyContentsAndResult

	type json_UpdateConsensusKeyContentsAndResult struct {
		Marker0 any `json:"kind"`
		UpdateConsensusKeyContentsAndResult_no_json_marshaller
	}

	tmp := json_UpdateConsensusKeyContentsAndResult {
		Marker0: self.OperationKind(),
		UpdateConsensusKeyContentsAndResult_no_json_marshaller: UpdateConsensusKeyContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *IncreasePaidStorageSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller IncreasePaidStorageSuccessfulManagerResult

	type json_IncreasePaidStorageSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_IncreasePaidStorageSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller: IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *IncreasePaidStorageContentsAndResult) MarshalJSON() ([]byte, error) {
	type IncreasePaidStorageContentsAndResult_no_json_marshaller IncreasePaidStorageContentsAndResult

	type json_IncreasePaidStorageContentsAndResult struct {
		Marker0 any `json:"kind"`
		IncreasePaidStorageContentsAndResult_no_json_marshaller
	}

	tmp := json_IncreasePaidStorageContentsAndResult {
		Marker0: self.OperationKind(),
		IncreasePaidStorageContentsAndResult_no_json_marshaller: IncreasePaidStorageContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *TransferTicketContentsAndResult) MarshalJSON() ([]byte, error) {
	type TransferTicketContentsAndResult_no_json_marshaller TransferTicketContentsAndResult

	type json_TransferTicketContentsAndResult struct {
		Marker0 any `json:"kind"`
		TransferTicketContentsAndResult_no_json_marshaller
	}

	tmp := json_TransferTicketContentsAndResult {
		Marker0: self.OperationKind(),
		TransferTicketContentsAndResult_no_json_marshaller: TransferTicketContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DALPublishSlotHeaderContentsAndResult) MarshalJSON() ([]byte, error) {
	type DALPublishSlotHeaderContentsAndResult_no_json_marshaller DALPublishSlotHeaderContentsAndResult

	type json_DALPublishSlotHeaderContentsAndResult struct {
		Marker0 any `json:"kind"`
		DALPublishSlotHeaderContentsAndResult_no_json_marshaller
	}

	tmp := json_DALPublishSlotHeaderContentsAndResult {
		Marker0: self.OperationKind(),
		DALPublishSlotHeaderContentsAndResult_no_json_marshaller: DALPublishSlotHeaderContentsAndResult_no_json_marshaller(*self),
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

func (self *ZkRollupOrigination) MarshalJSON() ([]byte, error) {
	type ZkRollupOrigination_no_json_marshaller ZkRollupOrigination

	type json_ZkRollupOrigination struct {
		Marker0 any `json:"kind"`
		ZkRollupOrigination_no_json_marshaller
	}

	tmp := json_ZkRollupOrigination {
		Marker0: self.OperationKind(),
		ZkRollupOrigination_no_json_marshaller: ZkRollupOrigination_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *ZkRollupPublish) MarshalJSON() ([]byte, error) {
	type ZkRollupPublish_no_json_marshaller ZkRollupPublish

	type json_ZkRollupPublish struct {
		Marker0 any `json:"kind"`
		ZkRollupPublish_no_json_marshaller
	}

	tmp := json_ZkRollupPublish {
		Marker0: self.OperationKind(),
		ZkRollupPublish_no_json_marshaller: ZkRollupPublish_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *ZkRollupOriginationContentsAndResult) MarshalJSON() ([]byte, error) {
	type ZkRollupOriginationContentsAndResult_no_json_marshaller ZkRollupOriginationContentsAndResult

	type json_ZkRollupOriginationContentsAndResult struct {
		Marker0 any `json:"kind"`
		ZkRollupOriginationContentsAndResult_no_json_marshaller
	}

	tmp := json_ZkRollupOriginationContentsAndResult {
		Marker0: self.OperationKind(),
		ZkRollupOriginationContentsAndResult_no_json_marshaller: ZkRollupOriginationContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *ZkRollupPublishContentsAndResult) MarshalJSON() ([]byte, error) {
	type ZkRollupPublishContentsAndResult_no_json_marshaller ZkRollupPublishContentsAndResult

	type json_ZkRollupPublishContentsAndResult struct {
		Marker0 any `json:"kind"`
		ZkRollupPublishContentsAndResult_no_json_marshaller
	}

	tmp := json_ZkRollupPublishContentsAndResult {
		Marker0: self.OperationKind(),
		ZkRollupPublishContentsAndResult_no_json_marshaller: ZkRollupPublishContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BalanceUpdateScRollupRefutationRewards) MarshalJSON() ([]byte, error) {
	type BalanceUpdateScRollupRefutationRewards_no_json_marshaller BalanceUpdateScRollupRefutationRewards

	type json_BalanceUpdateScRollupRefutationRewards struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		BalanceUpdateScRollupRefutationRewards_no_json_marshaller
	}

	tmp := json_BalanceUpdateScRollupRefutationRewards {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateScRollupRefutationRewards_no_json_marshaller: BalanceUpdateScRollupRefutationRewards_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}


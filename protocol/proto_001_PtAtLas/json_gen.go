package proto_001_PtAtLas

import "encoding/json"

// Code generated by genmarshaller.go DO NOT EDIT.

func (self *SmartRollupOriginate) MarshalJSON() ([]byte, error) {
	type SmartRollupOriginate_no_json_marshaller SmartRollupOriginate

	type json_SmartRollupOriginate struct {
		Marker0 any `json:"kind"`
		SmartRollupOriginate_no_json_marshaller
	}

	tmp := json_SmartRollupOriginate {
		Marker0: self.OperationKind(),
		SmartRollupOriginate_no_json_marshaller: SmartRollupOriginate_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SmartRollupCement) MarshalJSON() ([]byte, error) {
	type SmartRollupCement_no_json_marshaller SmartRollupCement

	type json_SmartRollupCement struct {
		Marker0 any `json:"kind"`
		SmartRollupCement_no_json_marshaller
	}

	tmp := json_SmartRollupCement {
		Marker0: self.OperationKind(),
		SmartRollupCement_no_json_marshaller: SmartRollupCement_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SmartRollupOriginateContentsAndResult) MarshalJSON() ([]byte, error) {
	type SmartRollupOriginateContentsAndResult_no_json_marshaller SmartRollupOriginateContentsAndResult

	type json_SmartRollupOriginateContentsAndResult struct {
		Marker0 any `json:"kind"`
		SmartRollupOriginateContentsAndResult_no_json_marshaller
	}

	tmp := json_SmartRollupOriginateContentsAndResult {
		Marker0: self.OperationKind(),
		SmartRollupOriginateContentsAndResult_no_json_marshaller: SmartRollupOriginateContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SmartRollupAddMessagesContentsAndResult) MarshalJSON() ([]byte, error) {
	type SmartRollupAddMessagesContentsAndResult_no_json_marshaller SmartRollupAddMessagesContentsAndResult

	type json_SmartRollupAddMessagesContentsAndResult struct {
		Marker0 any `json:"kind"`
		SmartRollupAddMessagesContentsAndResult_no_json_marshaller
	}

	tmp := json_SmartRollupAddMessagesContentsAndResult {
		Marker0: self.OperationKind(),
		SmartRollupAddMessagesContentsAndResult_no_json_marshaller: SmartRollupAddMessagesContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SmartRollupCementContentsAndResult) MarshalJSON() ([]byte, error) {
	type SmartRollupCementContentsAndResult_no_json_marshaller SmartRollupCementContentsAndResult

	type json_SmartRollupCementContentsAndResult struct {
		Marker0 any `json:"kind"`
		SmartRollupCementContentsAndResult_no_json_marshaller
	}

	tmp := json_SmartRollupCementContentsAndResult {
		Marker0: self.OperationKind(),
		SmartRollupCementContentsAndResult_no_json_marshaller: SmartRollupCementContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SmartRollupPublishContentsAndResult) MarshalJSON() ([]byte, error) {
	type SmartRollupPublishContentsAndResult_no_json_marshaller SmartRollupPublishContentsAndResult

	type json_SmartRollupPublishContentsAndResult struct {
		Marker0 any `json:"kind"`
		SmartRollupPublishContentsAndResult_no_json_marshaller
	}

	tmp := json_SmartRollupPublishContentsAndResult {
		Marker0: self.OperationKind(),
		SmartRollupPublishContentsAndResult_no_json_marshaller: SmartRollupPublishContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SmartRollupTimeoutContentsAndResult) MarshalJSON() ([]byte, error) {
	type SmartRollupTimeoutContentsAndResult_no_json_marshaller SmartRollupTimeoutContentsAndResult

	type json_SmartRollupTimeoutContentsAndResult struct {
		Marker0 any `json:"kind"`
		SmartRollupTimeoutContentsAndResult_no_json_marshaller
	}

	tmp := json_SmartRollupTimeoutContentsAndResult {
		Marker0: self.OperationKind(),
		SmartRollupTimeoutContentsAndResult_no_json_marshaller: SmartRollupTimeoutContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SmartRollupRefuteContentsAndResult) MarshalJSON() ([]byte, error) {
	type SmartRollupRefuteContentsAndResult_no_json_marshaller SmartRollupRefuteContentsAndResult

	type json_SmartRollupRefuteContentsAndResult struct {
		Marker0 any `json:"kind"`
		SmartRollupRefuteContentsAndResult_no_json_marshaller
	}

	tmp := json_SmartRollupRefuteContentsAndResult {
		Marker0: self.OperationKind(),
		SmartRollupRefuteContentsAndResult_no_json_marshaller: SmartRollupRefuteContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SmartRollupExecuteOutboxMessageContentsAndResult) MarshalJSON() ([]byte, error) {
	type SmartRollupExecuteOutboxMessageContentsAndResult_no_json_marshaller SmartRollupExecuteOutboxMessageContentsAndResult

	type json_SmartRollupExecuteOutboxMessageContentsAndResult struct {
		Marker0 any `json:"kind"`
		SmartRollupExecuteOutboxMessageContentsAndResult_no_json_marshaller
	}

	tmp := json_SmartRollupExecuteOutboxMessageContentsAndResult {
		Marker0: self.OperationKind(),
		SmartRollupExecuteOutboxMessageContentsAndResult_no_json_marshaller: SmartRollupExecuteOutboxMessageContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SmartRollupRecoverBondContentsAndResult) MarshalJSON() ([]byte, error) {
	type SmartRollupRecoverBondContentsAndResult_no_json_marshaller SmartRollupRecoverBondContentsAndResult

	type json_SmartRollupRecoverBondContentsAndResult struct {
		Marker0 any `json:"kind"`
		SmartRollupRecoverBondContentsAndResult_no_json_marshaller
	}

	tmp := json_SmartRollupRecoverBondContentsAndResult {
		Marker0: self.OperationKind(),
		SmartRollupRecoverBondContentsAndResult_no_json_marshaller: SmartRollupRecoverBondContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SmartRollupOriginateSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type SmartRollupOriginateSuccessfulManagerResult_no_json_marshaller SmartRollupOriginateSuccessfulManagerResult

	type json_SmartRollupOriginateSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		SmartRollupOriginateSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_SmartRollupOriginateSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		SmartRollupOriginateSuccessfulManagerResult_no_json_marshaller: SmartRollupOriginateSuccessfulManagerResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DoubleAttestationEvidence) MarshalJSON() ([]byte, error) {
	type DoubleAttestationEvidence_no_json_marshaller DoubleAttestationEvidence

	type json_DoubleAttestationEvidence struct {
		Marker0 any `json:"kind"`
		DoubleAttestationEvidence_no_json_marshaller
	}

	tmp := json_DoubleAttestationEvidence {
		Marker0: self.OperationKind(),
		DoubleAttestationEvidence_no_json_marshaller: DoubleAttestationEvidence_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *Attestation) MarshalJSON() ([]byte, error) {
	type Attestation_no_json_marshaller Attestation

	type json_Attestation struct {
		Marker0 any `json:"kind"`
		Attestation_no_json_marshaller
	}

	tmp := json_Attestation {
		Marker0: self.OperationKind(),
		Attestation_no_json_marshaller: Attestation_no_json_marshaller(*self),
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

func (self *DoublePreattestationEvidence) MarshalJSON() ([]byte, error) {
	type DoublePreattestationEvidence_no_json_marshaller DoublePreattestationEvidence

	type json_DoublePreattestationEvidence struct {
		Marker0 any `json:"kind"`
		DoublePreattestationEvidence_no_json_marshaller
	}

	tmp := json_DoublePreattestationEvidence {
		Marker0: self.OperationKind(),
		DoublePreattestationEvidence_no_json_marshaller: DoublePreattestationEvidence_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *Preattestation) MarshalJSON() ([]byte, error) {
	type Preattestation_no_json_marshaller Preattestation

	type json_Preattestation struct {
		Marker0 any `json:"kind"`
		Preattestation_no_json_marshaller
	}

	tmp := json_Preattestation {
		Marker0: self.OperationKind(),
		Preattestation_no_json_marshaller: Preattestation_no_json_marshaller(*self),
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

func (self *DoubleAttestationEvidenceContentsAndResult) MarshalJSON() ([]byte, error) {
	type DoubleAttestationEvidenceContentsAndResult_no_json_marshaller DoubleAttestationEvidenceContentsAndResult

	type json_DoubleAttestationEvidenceContentsAndResult struct {
		Marker0 any `json:"kind"`
		DoubleAttestationEvidenceContentsAndResult_no_json_marshaller
	}

	tmp := json_DoubleAttestationEvidenceContentsAndResult {
		Marker0: self.OperationKind(),
		DoubleAttestationEvidenceContentsAndResult_no_json_marshaller: DoubleAttestationEvidenceContentsAndResult_no_json_marshaller(*self),
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

func (self *DoublePreattestationEvidenceContentsAndResult) MarshalJSON() ([]byte, error) {
	type DoublePreattestationEvidenceContentsAndResult_no_json_marshaller DoublePreattestationEvidenceContentsAndResult

	type json_DoublePreattestationEvidenceContentsAndResult struct {
		Marker0 any `json:"kind"`
		DoublePreattestationEvidenceContentsAndResult_no_json_marshaller
	}

	tmp := json_DoublePreattestationEvidenceContentsAndResult {
		Marker0: self.OperationKind(),
		DoublePreattestationEvidenceContentsAndResult_no_json_marshaller: DoublePreattestationEvidenceContentsAndResult_no_json_marshaller(*self),
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

func (self *AttestationContentsAndResult) MarshalJSON() ([]byte, error) {
	type AttestationContentsAndResult_no_json_marshaller AttestationContentsAndResult

	type json_AttestationContentsAndResult struct {
		Marker0 any `json:"kind"`
		AttestationContentsAndResult_no_json_marshaller
	}

	tmp := json_AttestationContentsAndResult {
		Marker0: self.OperationKind(),
		AttestationContentsAndResult_no_json_marshaller: AttestationContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *PreattestationContentsAndResult) MarshalJSON() ([]byte, error) {
	type PreattestationContentsAndResult_no_json_marshaller PreattestationContentsAndResult

	type json_PreattestationContentsAndResult struct {
		Marker0 any `json:"kind"`
		PreattestationContentsAndResult_no_json_marshaller
	}

	tmp := json_PreattestationContentsAndResult {
		Marker0: self.OperationKind(),
		PreattestationContentsAndResult_no_json_marshaller: PreattestationContentsAndResult_no_json_marshaller(*self),
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

func (self *ZkRollupUpdateContentsAndResult) MarshalJSON() ([]byte, error) {
	type ZkRollupUpdateContentsAndResult_no_json_marshaller ZkRollupUpdateContentsAndResult

	type json_ZkRollupUpdateContentsAndResult struct {
		Marker0 any `json:"kind"`
		ZkRollupUpdateContentsAndResult_no_json_marshaller
	}

	tmp := json_ZkRollupUpdateContentsAndResult {
		Marker0: self.OperationKind(),
		ZkRollupUpdateContentsAndResult_no_json_marshaller: ZkRollupUpdateContentsAndResult_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

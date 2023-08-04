package proto_016_PtMumbai

import "encoding/json"

// Code generated by genmarshaller.go DO NOT EDIT.

func (self *TransactionSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type TransactionSuccessfulManagerResult_no_json_marshaller TransactionSuccessfulManagerResult

	type json_TransactionSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*TransactionSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_TransactionSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		TransactionSuccessfulManagerResult_no_json_marshaller: (*TransactionSuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *TransactionInternalOperationResult) MarshalJSON() ([]byte, error) {
	type TransactionInternalOperationResult_no_json_marshaller TransactionInternalOperationResult

	type json_TransactionInternalOperationResult struct {
		Marker0 any `json:"kind"`
		*TransactionInternalOperationResult_no_json_marshaller
	}

	tmp := json_TransactionInternalOperationResult {
		Marker0: self.OperationKind(),
		TransactionInternalOperationResult_no_json_marshaller: (*TransactionInternalOperationResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *OriginationSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type OriginationSuccessfulManagerResult_no_json_marshaller OriginationSuccessfulManagerResult

	type json_OriginationSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*OriginationSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_OriginationSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		OriginationSuccessfulManagerResult_no_json_marshaller: (*OriginationSuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *OriginationInternalOperationResult) MarshalJSON() ([]byte, error) {
	type OriginationInternalOperationResult_no_json_marshaller OriginationInternalOperationResult

	type json_OriginationInternalOperationResult struct {
		Marker0 any `json:"kind"`
		*OriginationInternalOperationResult_no_json_marshaller
	}

	tmp := json_OriginationInternalOperationResult {
		Marker0: self.OperationKind(),
		OriginationInternalOperationResult_no_json_marshaller: (*OriginationInternalOperationResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *ZkRollupUpdate) MarshalJSON() ([]byte, error) {
	type ZkRollupUpdate_no_json_marshaller ZkRollupUpdate

	type json_ZkRollupUpdate struct {
		Marker0 any `json:"kind"`
		*ZkRollupUpdate_no_json_marshaller
	}

	tmp := json_ZkRollupUpdate {
		Marker0: self.OperationKind(),
		ZkRollupUpdate_no_json_marshaller: (*ZkRollupUpdate_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *BalanceUpdateFrozenBonds) MarshalJSON() ([]byte, error) {
	type BalanceUpdateFrozenBonds_no_json_marshaller BalanceUpdateFrozenBonds

	type json_BalanceUpdateFrozenBonds struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		*BalanceUpdateFrozenBonds_no_json_marshaller
	}

	tmp := json_BalanceUpdateFrozenBonds {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateFrozenBonds_no_json_marshaller: (*BalanceUpdateFrozenBonds_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *BalanceUpdateSmartRollupRefutationPunishments) MarshalJSON() ([]byte, error) {
	type BalanceUpdateSmartRollupRefutationPunishments_no_json_marshaller BalanceUpdateSmartRollupRefutationPunishments

	type json_BalanceUpdateSmartRollupRefutationPunishments struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		*BalanceUpdateSmartRollupRefutationPunishments_no_json_marshaller
	}

	tmp := json_BalanceUpdateSmartRollupRefutationPunishments {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateSmartRollupRefutationPunishments_no_json_marshaller: (*BalanceUpdateSmartRollupRefutationPunishments_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *BalanceUpdateSmartRollupRefutationRewards) MarshalJSON() ([]byte, error) {
	type BalanceUpdateSmartRollupRefutationRewards_no_json_marshaller BalanceUpdateSmartRollupRefutationRewards

	type json_BalanceUpdateSmartRollupRefutationRewards struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		*BalanceUpdateSmartRollupRefutationRewards_no_json_marshaller
	}

	tmp := json_BalanceUpdateSmartRollupRefutationRewards {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateSmartRollupRefutationRewards_no_json_marshaller: (*BalanceUpdateSmartRollupRefutationRewards_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *DALAttestation) MarshalJSON() ([]byte, error) {
	type DALAttestation_no_json_marshaller DALAttestation

	type json_DALAttestation struct {
		Marker0 any `json:"kind"`
		*DALAttestation_no_json_marshaller
	}

	tmp := json_DALAttestation {
		Marker0: self.OperationKind(),
		DALAttestation_no_json_marshaller: (*DALAttestation_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *IncreasePaidStorageSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller IncreasePaidStorageSuccessfulManagerResult

	type json_IncreasePaidStorageSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_IncreasePaidStorageSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller: (*IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *DoubleBakingEvidence) MarshalJSON() ([]byte, error) {
	type DoubleBakingEvidence_no_json_marshaller DoubleBakingEvidence

	type json_DoubleBakingEvidence struct {
		Marker0 any `json:"kind"`
		*DoubleBakingEvidence_no_json_marshaller
	}

	tmp := json_DoubleBakingEvidence {
		Marker0: self.OperationKind(),
		DoubleBakingEvidence_no_json_marshaller: (*DoubleBakingEvidence_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *DALPublishSlotHeader) MarshalJSON() ([]byte, error) {
	type DALPublishSlotHeader_no_json_marshaller DALPublishSlotHeader

	type json_DALPublishSlotHeader struct {
		Marker0 any `json:"kind"`
		*DALPublishSlotHeader_no_json_marshaller
	}

	tmp := json_DALPublishSlotHeader {
		Marker0: self.OperationKind(),
		DALPublishSlotHeader_no_json_marshaller: (*DALPublishSlotHeader_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *SmartRollupOriginate) MarshalJSON() ([]byte, error) {
	type SmartRollupOriginate_no_json_marshaller SmartRollupOriginate

	type json_SmartRollupOriginate struct {
		Marker0 any `json:"kind"`
		*SmartRollupOriginate_no_json_marshaller
	}

	tmp := json_SmartRollupOriginate {
		Marker0: self.OperationKind(),
		SmartRollupOriginate_no_json_marshaller: (*SmartRollupOriginate_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *SmartRollupOriginateSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type SmartRollupOriginateSuccessfulManagerResult_no_json_marshaller SmartRollupOriginateSuccessfulManagerResult

	type json_SmartRollupOriginateSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*SmartRollupOriginateSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_SmartRollupOriginateSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		SmartRollupOriginateSuccessfulManagerResult_no_json_marshaller: (*SmartRollupOriginateSuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *SmartRollupAddMessages) MarshalJSON() ([]byte, error) {
	type SmartRollupAddMessages_no_json_marshaller SmartRollupAddMessages

	type json_SmartRollupAddMessages struct {
		Marker0 any `json:"kind"`
		*SmartRollupAddMessages_no_json_marshaller
	}

	tmp := json_SmartRollupAddMessages {
		Marker0: self.OperationKind(),
		SmartRollupAddMessages_no_json_marshaller: (*SmartRollupAddMessages_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *SmartRollupCement) MarshalJSON() ([]byte, error) {
	type SmartRollupCement_no_json_marshaller SmartRollupCement

	type json_SmartRollupCement struct {
		Marker0 any `json:"kind"`
		*SmartRollupCement_no_json_marshaller
	}

	tmp := json_SmartRollupCement {
		Marker0: self.OperationKind(),
		SmartRollupCement_no_json_marshaller: (*SmartRollupCement_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *SmartRollupPublish) MarshalJSON() ([]byte, error) {
	type SmartRollupPublish_no_json_marshaller SmartRollupPublish

	type json_SmartRollupPublish struct {
		Marker0 any `json:"kind"`
		*SmartRollupPublish_no_json_marshaller
	}

	tmp := json_SmartRollupPublish {
		Marker0: self.OperationKind(),
		SmartRollupPublish_no_json_marshaller: (*SmartRollupPublish_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *SmartRollupRefute) MarshalJSON() ([]byte, error) {
	type SmartRollupRefute_no_json_marshaller SmartRollupRefute

	type json_SmartRollupRefute struct {
		Marker0 any `json:"kind"`
		*SmartRollupRefute_no_json_marshaller
	}

	tmp := json_SmartRollupRefute {
		Marker0: self.OperationKind(),
		SmartRollupRefute_no_json_marshaller: (*SmartRollupRefute_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *SmartRollupTimeout) MarshalJSON() ([]byte, error) {
	type SmartRollupTimeout_no_json_marshaller SmartRollupTimeout

	type json_SmartRollupTimeout struct {
		Marker0 any `json:"kind"`
		*SmartRollupTimeout_no_json_marshaller
	}

	tmp := json_SmartRollupTimeout {
		Marker0: self.OperationKind(),
		SmartRollupTimeout_no_json_marshaller: (*SmartRollupTimeout_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *SmartRollupExecuteOutboxMessage) MarshalJSON() ([]byte, error) {
	type SmartRollupExecuteOutboxMessage_no_json_marshaller SmartRollupExecuteOutboxMessage

	type json_SmartRollupExecuteOutboxMessage struct {
		Marker0 any `json:"kind"`
		*SmartRollupExecuteOutboxMessage_no_json_marshaller
	}

	tmp := json_SmartRollupExecuteOutboxMessage {
		Marker0: self.OperationKind(),
		SmartRollupExecuteOutboxMessage_no_json_marshaller: (*SmartRollupExecuteOutboxMessage_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *SmartRollupRecoverBond) MarshalJSON() ([]byte, error) {
	type SmartRollupRecoverBond_no_json_marshaller SmartRollupRecoverBond

	type json_SmartRollupRecoverBond struct {
		Marker0 any `json:"kind"`
		*SmartRollupRecoverBond_no_json_marshaller
	}

	tmp := json_SmartRollupRecoverBond {
		Marker0: self.OperationKind(),
		SmartRollupRecoverBond_no_json_marshaller: (*SmartRollupRecoverBond_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}


package proto_014_PtKathma

import "encoding/json"

// Code generated by genmarshaller.go DO NOT EDIT.

func (self *RevealSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type RevealSuccessfulManagerResult_no_json_marshaller RevealSuccessfulManagerResult

	type json_RevealSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*RevealSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_RevealSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		RevealSuccessfulManagerResult_no_json_marshaller: (*RevealSuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *DelegationSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type DelegationSuccessfulManagerResult_no_json_marshaller DelegationSuccessfulManagerResult

	type json_DelegationSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*DelegationSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_DelegationSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		DelegationSuccessfulManagerResult_no_json_marshaller: (*DelegationSuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *DelegationInternalOperationResult) MarshalJSON() ([]byte, error) {
	type DelegationInternalOperationResult_no_json_marshaller DelegationInternalOperationResult

	type json_DelegationInternalOperationResult struct {
		Marker0 any `json:"kind"`
		*DelegationInternalOperationResult_no_json_marshaller
	}

	tmp := json_DelegationInternalOperationResult {
		Marker0: self.OperationKind(),
		DelegationInternalOperationResult_no_json_marshaller: (*DelegationInternalOperationResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *SetDepositsLimitSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type SetDepositsLimitSuccessfulManagerResult_no_json_marshaller SetDepositsLimitSuccessfulManagerResult

	type json_SetDepositsLimitSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*SetDepositsLimitSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_SetDepositsLimitSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		SetDepositsLimitSuccessfulManagerResult_no_json_marshaller: (*SetDepositsLimitSuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *IncreasePaidStorage) MarshalJSON() ([]byte, error) {
	type IncreasePaidStorage_no_json_marshaller IncreasePaidStorage

	type json_IncreasePaidStorage struct {
		Marker0 any `json:"kind"`
		*IncreasePaidStorage_no_json_marshaller
	}

	tmp := json_IncreasePaidStorage {
		Marker0: self.OperationKind(),
		IncreasePaidStorage_no_json_marshaller: (*IncreasePaidStorage_no_json_marshaller)(self),
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

func (self *VDFRevelation) MarshalJSON() ([]byte, error) {
	type VDFRevelation_no_json_marshaller VDFRevelation

	type json_VDFRevelation struct {
		Marker0 any `json:"kind"`
		*VDFRevelation_no_json_marshaller
	}

	tmp := json_VDFRevelation {
		Marker0: self.OperationKind(),
		VDFRevelation_no_json_marshaller: (*VDFRevelation_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *DALSlotAvailability) MarshalJSON() ([]byte, error) {
	type DALSlotAvailability_no_json_marshaller DALSlotAvailability

	type json_DALSlotAvailability struct {
		Marker0 any `json:"kind"`
		*DALSlotAvailability_no_json_marshaller
	}

	tmp := json_DALSlotAvailability {
		Marker0: self.OperationKind(),
		DALSlotAvailability_no_json_marshaller: (*DALSlotAvailability_no_json_marshaller)(self),
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

func (self *EventInternalOperationResult) MarshalJSON() ([]byte, error) {
	type EventInternalOperationResult_no_json_marshaller EventInternalOperationResult

	type json_EventInternalOperationResult struct {
		Marker0 any `json:"kind"`
		*EventInternalOperationResult_no_json_marshaller
	}

	tmp := json_EventInternalOperationResult {
		Marker0: self.OperationKind(),
		EventInternalOperationResult_no_json_marshaller: (*EventInternalOperationResult_no_json_marshaller)(self),
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

func (self *BalanceUpdateScRollupRefutationPunishments) MarshalJSON() ([]byte, error) {
	type BalanceUpdateScRollupRefutationPunishments_no_json_marshaller BalanceUpdateScRollupRefutationPunishments

	type json_BalanceUpdateScRollupRefutationPunishments struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		*BalanceUpdateScRollupRefutationPunishments_no_json_marshaller
	}

	tmp := json_BalanceUpdateScRollupRefutationPunishments {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateScRollupRefutationPunishments_no_json_marshaller: (*BalanceUpdateScRollupRefutationPunishments_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

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

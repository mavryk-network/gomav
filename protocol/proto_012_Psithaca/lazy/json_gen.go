package lazy

import "encoding/json"

// Code generated by genmarshaller.go DO NOT EDIT.

func (self *BigMap) MarshalJSON() ([]byte, error) {
	type BigMap_no_json_marshaller BigMap

	type json_BigMap struct {
		Marker0 any `json:"kind"`
		BigMap_no_json_marshaller
	}

	tmp := json_BigMap {
		Marker0: self.LazyStorageDiffKind(),
		BigMap_no_json_marshaller: BigMap_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BigMapUpdate) MarshalJSON() ([]byte, error) {
	type BigMapUpdate_no_json_marshaller BigMapUpdate

	type json_BigMapUpdate struct {
		Marker0 any `json:"action"`
		BigMapUpdate_no_json_marshaller
	}

	tmp := json_BigMapUpdate {
		Marker0: self.LazyStorageBigMapOp(),
		BigMapUpdate_no_json_marshaller: BigMapUpdate_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BigMapRemove) MarshalJSON() ([]byte, error) {
	type BigMapRemove_no_json_marshaller BigMapRemove

	type json_BigMapRemove struct {
		Marker0 any `json:"action"`
		BigMapRemove_no_json_marshaller
	}

	tmp := json_BigMapRemove {
		Marker0: self.LazyStorageBigMapOp(),
		BigMapRemove_no_json_marshaller: BigMapRemove_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BigMapCopy) MarshalJSON() ([]byte, error) {
	type BigMapCopy_no_json_marshaller BigMapCopy

	type json_BigMapCopy struct {
		Marker0 any `json:"action"`
		BigMapCopy_no_json_marshaller
	}

	tmp := json_BigMapCopy {
		Marker0: self.LazyStorageBigMapOp(),
		BigMapCopy_no_json_marshaller: BigMapCopy_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *BigMapAlloc) MarshalJSON() ([]byte, error) {
	type BigMapAlloc_no_json_marshaller BigMapAlloc

	type json_BigMapAlloc struct {
		Marker0 any `json:"action"`
		BigMapAlloc_no_json_marshaller
	}

	tmp := json_BigMapAlloc {
		Marker0: self.LazyStorageBigMapOp(),
		BigMapAlloc_no_json_marshaller: BigMapAlloc_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SaplingState) MarshalJSON() ([]byte, error) {
	type SaplingState_no_json_marshaller SaplingState

	type json_SaplingState struct {
		Marker0 any `json:"kind"`
		SaplingState_no_json_marshaller
	}

	tmp := json_SaplingState {
		Marker0: self.LazyStorageDiffKind(),
		SaplingState_no_json_marshaller: SaplingState_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SaplingStateUpdate) MarshalJSON() ([]byte, error) {
	type SaplingStateUpdate_no_json_marshaller SaplingStateUpdate

	type json_SaplingStateUpdate struct {
		Marker0 any `json:"action"`
		SaplingStateUpdate_no_json_marshaller
	}

	tmp := json_SaplingStateUpdate {
		Marker0: self.LazyStorageSaplingStateOp(),
		SaplingStateUpdate_no_json_marshaller: SaplingStateUpdate_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SaplingStateRemove) MarshalJSON() ([]byte, error) {
	type SaplingStateRemove_no_json_marshaller SaplingStateRemove

	type json_SaplingStateRemove struct {
		Marker0 any `json:"action"`
		SaplingStateRemove_no_json_marshaller
	}

	tmp := json_SaplingStateRemove {
		Marker0: self.LazyStorageSaplingStateOp(),
		SaplingStateRemove_no_json_marshaller: SaplingStateRemove_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SaplingStateCopy) MarshalJSON() ([]byte, error) {
	type SaplingStateCopy_no_json_marshaller SaplingStateCopy

	type json_SaplingStateCopy struct {
		Marker0 any `json:"action"`
		SaplingStateCopy_no_json_marshaller
	}

	tmp := json_SaplingStateCopy {
		Marker0: self.LazyStorageSaplingStateOp(),
		SaplingStateCopy_no_json_marshaller: SaplingStateCopy_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *SaplingStateAlloc) MarshalJSON() ([]byte, error) {
	type SaplingStateAlloc_no_json_marshaller SaplingStateAlloc

	type json_SaplingStateAlloc struct {
		Marker0 any `json:"action"`
		SaplingStateAlloc_no_json_marshaller
	}

	tmp := json_SaplingStateAlloc {
		Marker0: self.LazyStorageSaplingStateOp(),
		SaplingStateAlloc_no_json_marshaller: SaplingStateAlloc_no_json_marshaller(*self),
	}

	return json.Marshal(tmp)
}


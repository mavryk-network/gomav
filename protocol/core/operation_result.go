package core

import (
	"math/big"

	mv "github.com/mavryk-network/gomav/v2"
)

type InternalOperationResult interface {
	OperationContents
	OperationWithResult
}

type ManagerOperationResult interface {
	Status() string
	IsApplied() bool
}

type ManagerOperationResultAppliedOrBacktracked interface {
	ManagerOperationResult
	GetResultContents() any
}

// SuccessfulManagerOperationResult is used to represent implicit operations results
type SuccessfulManagerOperationResult interface {
	OperationContents
	ManagerOperationResultAppliedOrBacktracked
}

//json:status=Status()
type OperationResultApplied[T any] struct {
	Contents T `json:"contents"`
}

func (*OperationResultApplied[T]) Status() string           { return "applied" }
func (*OperationResultApplied[T]) IsApplied() bool          { return true }
func (r *OperationResultApplied[T]) GetResultContents() any { return r.Contents }

var _ ManagerOperationResultAppliedOrBacktracked = (*OperationResultApplied[struct{}])(nil)

//json:status=Status()
type OperationResultBacktracked[T any] struct {
	Errors   mv.Option[OperationResultErrors] `json:"errors"`
	Contents T                                `json:"contents"`
}

func (*OperationResultBacktracked[T]) Status() string           { return "backtracked" }
func (*OperationResultBacktracked[T]) IsApplied() bool          { return false }
func (r *OperationResultBacktracked[T]) GetResultContents() any { return r.Contents }

var _ ManagerOperationResultAppliedOrBacktracked = (*OperationResultBacktracked[struct{}])(nil)

type OperationResultErrors struct {
	Errors []Bytes `mv:"dyn" json:"errors"`
}

//json:status=Status()
type OperationResultFailed OperationResultErrors

func (*OperationResultFailed) Status() string  { return "failed" }
func (*OperationResultFailed) IsApplied() bool { return false }

//json:status=Status()
type OperationResultSkipped struct{}

func (*OperationResultSkipped) Status() string  { return "skipped" }
func (*OperationResultSkipped) IsApplied() bool { return false }

type ResultWithConsumedMilligas interface {
	GetConsumedMilligas() mv.BigUint
}

type ResultWithStorageSize interface {
	StorageSizeEstimator
	GetStorageSize() mv.BigInt
}

type ResultWithPaidStorageSizeDiff interface {
	StorageSizeEstimator
	GetPaidStorageSizeDiff() mv.BigInt
}

type StorageSizeEstimator interface {
	EstimateStorageSize(constants Constants) *big.Int
}

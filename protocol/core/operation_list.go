package core

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
)

type GroupContents interface {
	GetSignature() (mv.Option[mv.Signature], error)
	GroupContents()
	Operations() []OperationContents
}

type OperationWithOptionalMetadataContents interface {
	GetSignature() (mv.Option[mv.Signature], error)
	OperationWithOptionalMetadataContents()
	Operations() []OperationContents
}

type OperationsList[T GroupContents] struct {
	Operations []*OperationsGroupImpl[T] `mv:"dyn,dyn" json:"operations"` // yes, twice
}

func (l *OperationsList[T]) GetGroups() []OperationsGroup {
	out := make([]OperationsGroup, len(l.Operations))
	for i, grp := range l.Operations {
		out[i] = grp
	}
	return out
}

type OperationsGroup interface {
	GetChainID() *mv.ChainID
	GetHash() *mv.OperationHash
	GetBranch() *mv.BlockHash
	GetContents() GroupContents
}

type OperationsGroupImpl[T GroupContents] struct {
	ChainID  *mv.ChainID       `json:"chain_id"`
	Hash     *mv.OperationHash `json:"hash"`
	Branch   *mv.BlockHash     `mv:"dyn" json:"branch"`
	Contents T                 `mv:"dyn" json:"contents"`
}

func (g *OperationsGroupImpl[T]) GetChainID() *mv.ChainID    { return g.ChainID }
func (g *OperationsGroupImpl[T]) GetHash() *mv.OperationHash { return g.Hash }
func (g *OperationsGroupImpl[T]) GetBranch() *mv.BlockHash   { return g.Branch }
func (g *OperationsGroupImpl[T]) GetContents() GroupContents { return g.Contents }

type OperationWithoutMetadata[T OperationContents] struct {
	Contents  []T                  `json:"contents"`
	Signature *mv.GenericSignature `json:"signature"`
}

func (op *OperationWithoutMetadata[T]) DecodeMV(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	if len(data) < mv.GenericSignatureBytesLen {
		return nil, encoding.ErrBuffer(len(data))
	}
	tmp := data[:len(data)-mv.GenericSignatureBytesLen]
	data = data[len(data)-mv.GenericSignatureBytesLen:]
	if _, err := encoding.Decode(tmp, &op.Contents, encoding.Ctx(ctx)); err != nil {
		return nil, err
	}
	return encoding.Decode(data, &op.Signature, encoding.Ctx(ctx))
}

func (ops *OperationWithoutMetadata[T]) Operations() []OperationContents {
	out := make([]OperationContents, len(ops.Contents))
	for i, op := range ops.Contents {
		out[i] = op
	}
	return out
}

func (*OperationWithoutMetadata[T]) GroupContents() {}

type OperationWithOptionalMetadata[T OperationWithOptionalMetadataContents] struct {
	Contents T `json:"contents"`
}

func (ops OperationWithOptionalMetadata[T]) Operations() []OperationContents {
	return ops.Contents.Operations()
}

func (op OperationWithOptionalMetadata[T]) GetSignature() (mv.Option[mv.Signature], error) {
	return op.Contents.GetSignature()
}

func (OperationWithOptionalMetadata[T]) GroupContents() {}

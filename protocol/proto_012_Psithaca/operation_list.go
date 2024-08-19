package proto_012_Psithaca

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
)

type GroupContents interface {
	core.GroupContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[GroupContents]{
		Variants: encoding.Variants[GroupContents]{
			0: (*OperationWithTooLargeMetadata[OperationContents])(nil),
			1: (*OperationWithoutMetadata[OperationContents])(nil),
			2: (*core.OperationWithOptionalMetadata[OperationWithOptionalMetadataContents])(nil),
		},
	})
}

type OperationWithoutMetadata[T core.OperationContents] struct {
	core.OperationWithoutMetadata[T]
}

func (op *OperationWithoutMetadata[T]) GetSignature() (mv.Option[mv.Signature], error) {
	return mv.Some[mv.Signature](op.Signature), nil
}

type OperationWithTooLargeMetadata[T core.OperationContents] struct {
	OperationWithoutMetadata[T]
}

type OperationWithOptionalMetadataContents interface {
	core.OperationWithOptionalMetadataContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationWithOptionalMetadataContents]{
		Variants: encoding.Variants[OperationWithOptionalMetadataContents]{
			0: (*OperationWithOptionalMetadataWithMetadata[OperationContentsAndResult])(nil),
			1: (*OperationWithOptionalMetadataWithoutMetadata[OperationContents])(nil),
		},
	})
}

type OperationWithOptionalMetadataWithMetadata[T core.OperationContentsAndResult] struct {
	Contents  []T                             `mv:"dyn" json:"contents"`
	Signature mv.Option[*mv.GenericSignature] `json:"signature"`
}

func (ops *OperationWithOptionalMetadataWithMetadata[T]) Operations() []core.OperationContents {
	out := make([]core.OperationContents, len(ops.Contents))
	for i, op := range ops.Contents {
		out[i] = op
	}
	return out
}

func (*OperationWithOptionalMetadataWithMetadata[T]) OperationWithOptionalMetadataContents() {}
func (op *OperationWithOptionalMetadataWithMetadata[T]) GetSignature() (mv.Option[mv.Signature], error) {
	if sig, ok := op.Signature.CheckUnwrap(); ok {
		return mv.Some[mv.Signature](sig), nil
	}
	return mv.None[mv.Signature](), nil
}

type OperationWithOptionalMetadataWithoutMetadata[T core.OperationContents] struct {
	Contents  []T                             `mv:"dyn" json:"contents"`
	Signature mv.Option[*mv.GenericSignature] `json:"signature"`
}

func (ops *OperationWithOptionalMetadataWithoutMetadata[T]) Operations() []core.OperationContents {
	out := make([]core.OperationContents, len(ops.Contents))
	for i, op := range ops.Contents {
		out[i] = op
	}
	return out
}

func (*OperationWithOptionalMetadataWithoutMetadata[T]) OperationWithOptionalMetadataContents() {}
func (op *OperationWithOptionalMetadataWithoutMetadata[T]) GetSignature() (mv.Option[mv.Signature], error) {
	if sig, ok := op.Signature.CheckUnwrap(); ok {
		return mv.Some[mv.Signature](sig), nil
	}
	return mv.None[mv.Signature](), nil
}

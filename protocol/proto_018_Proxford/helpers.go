package proto_018_Proxford

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
)

type UnsignedOperation = UnsignedOperationImpl[OperationContents]
type SignedOperation = SignedOperationImpl[OperationContents]

type UnsignedOperationImpl[T core.OperationContents] struct {
	Branch   *mv.BlockHash `json:"branch"`
	Contents []T           `json:"contents"`
}

type SignedOperationImpl[T core.OperationContents] struct {
	UnsignedOperationImpl[T]
	Signature *mv.GenericSignature `json:"signature"`
}

func (*SignedOperationImpl[T]) RunOperationRequestContents() {}

func (op *SignedOperationImpl[T]) DecodeMV(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	if data, err = encoding.Decode(data, &op.Branch, encoding.Ctx(ctx)); err != nil {
		return nil, err
	}
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

func NewRunOperationRequest(op *SignedOperation, chain *mv.ChainID) *RunOperationRequest {
	return &RunOperationRequest{
		Operation: op,
		ChainID:   chain,
	}
}

func NewUnsignedOperation(branch *mv.BlockHash, contents []OperationContents) *UnsignedOperation {
	return &UnsignedOperation{
		Branch:   branch,
		Contents: contents,
	}
}

func NewSignedOperation(operation *UnsignedOperation, signature *mv.GenericSignature) *SignedOperation {
	return &SignedOperation{
		UnsignedOperationImpl: *operation,
		Signature:             signature,
	}
}

type RunOperationRequest = RunOperationRequestImpl[RunOperationRequestContents]
type RunOperationRequestImpl[T RunOperationRequestContents] struct {
	Operation T           `json:"operation"`
	ChainID   *mv.ChainID `json:"chain_id"`
}

type RunOperationRequestContents interface {
	RunOperationRequestContents()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RunOperationRequestContents]{
		Variants: encoding.Variants[RunOperationRequestContents]{
			0: (*SignedOperation)(nil),
		},
	})
}

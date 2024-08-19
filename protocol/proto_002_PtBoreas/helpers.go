package proto_002_PtBoreas

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/proto_001_PtAtLas"
)

type UnsignedOperation = proto_001_PtAtLas.UnsignedOperationImpl[OperationContents]
type SignedOperation = proto_001_PtAtLas.SignedOperationImpl[OperationContents]
type RunOperationRequest = proto_001_PtAtLas.RunOperationRequestImpl[RunOperationRequestContents]

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

type RunOperationRequestContents interface {
	proto_001_PtAtLas.RunOperationRequestContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RunOperationRequestContents]{
		Variants: encoding.Variants[RunOperationRequestContents]{
			0: (*SignedOperation)(nil),
		},
	})
}

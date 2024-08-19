package proto_019_PtParisB

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/proto_018_Proxford"
)

type UnsignedOperation = proto_018_Proxford.UnsignedOperationImpl[OperationContents]
type SignedOperation = proto_018_Proxford.SignedOperationImpl[OperationContents]
type RunOperationRequest = proto_018_Proxford.RunOperationRequestImpl[RunOperationRequestContents]

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
	proto_018_Proxford.RunOperationRequestContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RunOperationRequestContents]{
		Variants: encoding.Variants[RunOperationRequestContents]{
			0: (*SignedOperation)(nil),
		},
	})
}

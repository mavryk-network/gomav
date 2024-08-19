package proto_013_PtJakart

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
)

type UnsignedOperation struct {
	Branch   *mv.BlockHash       `json:"branch"`
	Contents []OperationContents `json:"contents"`
}

type SignedOperation struct {
	UnsignedOperation
	Signature *mv.GenericSignature `json:"signature"`
}

func (op *SignedOperation) DecodeMV(data []byte, ctx *encoding.Context) (rest []byte, err error) {
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

type RunOperationRequest struct {
	Operation SignedOperation `json:"operation"`
	ChainID   *mv.ChainID     `json:"chain_id"`
}

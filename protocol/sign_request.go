package protocol

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/latest"
)

type SignRequest interface {
	SignRequestKind() string
}

type BlockSignRequest struct {
	Chain       *mv.ChainID
	BlockHeader latest.UnsignedBlockHeader
}

func (r *BlockSignRequest) GetChainID() *mv.ChainID { return r.Chain }
func (r *BlockSignRequest) GetLevel() int32         { return r.BlockHeader.Level }
func (r *BlockSignRequest) GetRound() int32         { return r.BlockHeader.PayloadRound }
func (*BlockSignRequest) SignRequestKind() string   { return "block" }

type PreattestationSignRequest struct {
	Chain     *mv.ChainID
	Branch    *mv.BlockHash
	Operation latest.InlinedPreendorsementContents
}

type PreendorsementSignRequest = PreattestationSignRequest

func (r *PreattestationSignRequest) GetChainID() *mv.ChainID { return r.Chain }
func (r *PreattestationSignRequest) GetLevel() int32 {
	return r.Operation.(*latest.Preattestation).Level
}
func (r *PreattestationSignRequest) GetRound() int32 {
	return r.Operation.(*latest.Preattestation).Round
}
func (*PreattestationSignRequest) SignRequestKind() string { return "preendorsement" }

type AttestationSignRequest struct {
	Chain     *mv.ChainID
	Branch    *mv.BlockHash
	Operation latest.InlinedEndorsementContents
}

type EndorsementSignRequest = AttestationSignRequest

func (r *AttestationSignRequest) GetChainID() *mv.ChainID { return r.Chain }
func (r *AttestationSignRequest) GetLevel() int32         { return r.Operation.(*latest.Attestation).Level }
func (r *AttestationSignRequest) GetRound() int32         { return r.Operation.(*latest.Attestation).Round }
func (*AttestationSignRequest) SignRequestKind() string   { return "endorsement" }

type GenericOperationSignRequest latest.UnsignedOperation

func (*GenericOperationSignRequest) SignRequestKind() string { return "generic" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[SignRequest]{
		Variants: encoding.Variants[SignRequest]{
			0x03: (*GenericOperationSignRequest)(nil),
			0x11: (*BlockSignRequest)(nil),
			0x12: (*PreattestationSignRequest)(nil),
			0x13: (*AttestationSignRequest)(nil),
		},
	})
}

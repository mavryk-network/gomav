package protocol

import (
	"fmt"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/proto_012_Psithaca"
	"github.com/mavryk-network/gomav/v2/protocol/proto_013_PtJakart"
	"github.com/mavryk-network/gomav/v2/protocol/proto_014_PtKathma"
	"github.com/mavryk-network/gomav/v2/protocol/proto_015_PtLimaPt"
	"github.com/mavryk-network/gomav/v2/protocol/proto_016_PtMumbai"
	"github.com/mavryk-network/gomav/v2/protocol/proto_017_PtNairob"
	"github.com/mavryk-network/gomav/v2/protocol/proto_018_Proxford"
	"github.com/mavryk-network/gomav/v2/protocol/proto_019_PtParisB"
)

type BlockInfo interface {
	GetChainID() *mv.ChainID
	GetHash() *mv.BlockHash
	GetHeader() core.BlockHeader
	GetMetadata() mv.Option[core.BlockMetadata]
	GetOperations() [][]core.OperationsGroup
}

type BlockHeaderInfo interface {
	GetChainID() *mv.ChainID
	GetHash() *mv.BlockHash
	core.BlockHeader
}

func NewBlockInfo(proto *mv.ProtocolHash) (BlockInfo, error) {
	switch *proto {
	case core.Proto019PtParisB:
		return new(proto_019_PtParisB.BlockInfo), nil
	case core.Proto018Proxford:
		return new(proto_018_Proxford.BlockInfo), nil
	case core.Proto017PtNairob:
		return new(proto_017_PtNairob.BlockInfo), nil
	case core.Proto016PtMumbai:
		return new(proto_016_PtMumbai.BlockInfo), nil
	case core.Proto015PtLimaPt:
		return new(proto_015_PtLimaPt.BlockInfo), nil
	case core.Proto014PtKathma:
		return new(proto_014_PtKathma.BlockInfo), nil
	case core.Proto013PtJakart:
		return new(proto_013_PtJakart.BlockInfo), nil
	case core.Proto012Psithaca:
		return new(proto_012_Psithaca.BlockInfo), nil
	default:
		return nil, fmt.Errorf("gomav: NewBlockInfo: unknown protocol %v", proto)
	}
}

func NewBlockHeaderInfo(proto *mv.ProtocolHash) (BlockHeaderInfo, error) {
	switch *proto {
	case core.Proto019PtParisB:
		return new(proto_019_PtParisB.BlockHeaderInfo), nil
	case core.Proto018Proxford:
		return new(proto_018_Proxford.BlockHeaderInfo), nil
	case core.Proto017PtNairob:
		return new(proto_017_PtNairob.BlockHeaderInfo), nil
	case core.Proto016PtMumbai:
		return new(proto_016_PtMumbai.BlockHeaderInfo), nil
	case core.Proto015PtLimaPt:
		return new(proto_015_PtLimaPt.BlockHeaderInfo), nil
	case core.Proto014PtKathma:
		return new(proto_014_PtKathma.BlockHeaderInfo), nil
	case core.Proto013PtJakart:
		return new(proto_013_PtJakart.BlockHeaderInfo), nil
	case core.Proto012Psithaca:
		return new(proto_012_Psithaca.BlockHeaderInfo), nil
	default:
		return nil, fmt.Errorf("gomav: NewBlockHeaderInfo: unknown protocol %v", proto)
	}
}

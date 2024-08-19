package protocol

import (
	"fmt"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/proto_001_PtAtLas"
	"github.com/mavryk-network/gomav/v2/protocol/proto_002_PtBoreas"
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
	case core.Proto002PtBoreas:
		return new(proto_002_PtBoreas.BlockInfo), nil
	case core.Proto001PtAtLas:
		return new(proto_001_PtAtLas.BlockInfo), nil
	default:
		return nil, fmt.Errorf("gomav: NewBlockInfo: unknown protocol %v", proto)
	}
}

func NewBlockHeaderInfo(proto *mv.ProtocolHash) (BlockHeaderInfo, error) {
	switch *proto {
	case core.Proto002PtBoreas:
		return new(proto_002_PtBoreas.BlockHeaderInfo), nil
	case core.Proto001PtAtLas:
		return new(proto_001_PtAtLas.BlockHeaderInfo), nil
	default:
		return nil, fmt.Errorf("gomav: NewBlockHeaderInfo: unknown protocol %v", proto)
	}
}

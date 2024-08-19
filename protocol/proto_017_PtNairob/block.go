package proto_017_PtNairob

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/proto_013_PtJakart"
	"github.com/mavryk-network/gomav/v2/protocol/proto_016_PtMumbai"
)

type UnsignedProtocolBlockHeader = proto_013_PtJakart.UnsignedProtocolBlockHeader
type UnsignedBlockHeader = proto_013_PtJakart.UnsignedBlockHeader
type BlockHeader = proto_016_PtMumbai.BlockHeader
type BlockHeaderInfo = proto_016_PtMumbai.BlockHeaderInfo
type BlockMetadata = proto_016_PtMumbai.BlockMetadata

type BlockInfo struct {
	ChainID    *mv.ChainID                          `json:"chain_id"`
	Hash       *mv.BlockHash                        `json:"hash"`
	Header     BlockHeader                          `mv:"dyn" json:"header"`
	Metadata   mv.Option[BlockMetadata]             `json:"metadata"`
	Operations []core.OperationsList[GroupContents] `mv:"dyn" json:"operations"`
}

func (block *BlockInfo) GetChainID() *mv.ChainID     { return block.ChainID }
func (block *BlockInfo) GetHash() *mv.BlockHash      { return block.Hash }
func (block *BlockInfo) GetHeader() core.BlockHeader { return &block.Header }
func (block *BlockInfo) GetMetadata() mv.Option[core.BlockMetadata] {
	if m, ok := block.Metadata.CheckUnwrapPtr(); ok {
		return mv.Some[core.BlockMetadata](m)
	}
	return mv.None[core.BlockMetadata]()
}

func (block *BlockInfo) GetOperations() [][]core.OperationsGroup {
	out := make([][]core.OperationsGroup, len(block.Operations))
	for i, list := range block.Operations {
		out[i] = list.GetGroups()
	}
	return out
}

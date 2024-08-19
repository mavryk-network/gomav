package proto_015_PtLimaPt

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/proto_013_PtJakart"
)

type UnsignedProtocolBlockHeader = proto_013_PtJakart.UnsignedProtocolBlockHeader
type UnsignedBlockHeader = proto_013_PtJakart.UnsignedBlockHeader
type BlockHeader = proto_013_PtJakart.BlockHeader
type BlockHeaderInfo = proto_013_PtJakart.BlockHeaderInfo

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

type BlockMetadata struct {
	BlockMetadataContents `mv:"dyn"`
}

type BlockMetadataContents struct {
	core.BlockMetadataHeader
	Proposer                  mv.PublicKeyHash                   `json:"proposer"`
	Baker                     mv.PublicKeyHash                   `json:"baker"`
	LevelInfo                 core.LevelInfo                     `json:"level_info"`
	VotingPeriodInfo          core.VotingPeriodInfo              `json:"voting_period_info"`
	NonceHash                 mv.Option1[*mv.CycleNonceHash]     `json:"nonce_hash"`
	Deactivated               []mv.PublicKeyHash                 `mv:"dyn" json:"deactivated"`
	BalanceUpdates            []*BalanceUpdate                   `mv:"dyn" json:"balance_updates"`
	LiquidityBakingToggleEMA  int32                              `json:"liquidity_baking_toggle_ema"`
	ImplicitOperationsResults []SuccessfulManagerOperationResult `mv:"dyn" json:"implicit_operations_results"`
	ProposerConsensusKey      mv.PublicKeyHash                   `json:"proposer_consensus_key"`
	BakerConsensusKey         mv.PublicKeyHash                   `json:"baker_consensus_key"`
	ConsumedMilligas          mv.BigUint                         `json:"consumed_milligas"`
	DALSlotAvailability       mv.Option[mv.BigInt]               `json:"dal_slot_availability"`
}

func (m *BlockMetadata) GetMetadataHeader() *core.BlockMetadataHeader { return &m.BlockMetadataHeader }
func (m *BlockMetadata) GetProposer() mv.PublicKeyHash                { return m.Proposer }
func (m *BlockMetadata) GetBaker() mv.PublicKeyHash                   { return m.Baker }
func (m *BlockMetadata) GetLevelInfo() *core.LevelInfo                { return &m.LevelInfo }
func (m *BlockMetadata) GetVotingPeriodInfo() *core.VotingPeriodInfo  { return &m.VotingPeriodInfo }
func (m *BlockMetadata) GetNonceHash() mv.Option[*mv.CycleNonceHash]  { return m.NonceHash.Option }
func (m *BlockMetadata) GetConsumedGas() mv.Option[mv.BigUint]        { return mv.None[mv.BigUint]() }
func (m *BlockMetadata) GetConsumedMilligas() mv.Option[mv.BigUint] {
	return mv.Some(m.ConsumedMilligas)
}
func (m *BlockMetadata) GetDeactivated() []mv.PublicKeyHash { return m.Deactivated }
func (m *BlockMetadata) GetBalanceUpdates() (updates []core.BalanceUpdate) {
	updates = make([]core.BalanceUpdate, len(m.BalanceUpdates))
	for i, u := range m.BalanceUpdates {
		updates[i] = u
	}
	return
}
func (m *BlockMetadata) GetImplicitOperationsResults() []core.SuccessfulManagerOperationResult {
	out := make([]core.SuccessfulManagerOperationResult, len(m.ImplicitOperationsResults))
	for i, v := range m.ImplicitOperationsResults {
		out[i] = v
	}
	return out
}
func (m *BlockMetadata) GetProposerConsensusKey() mv.Option[mv.PublicKeyHash] {
	return mv.Some(m.ProposerConsensusKey)
}
func (m *BlockMetadata) GetBakerConsensusKey() mv.Option[mv.PublicKeyHash] {
	return mv.Some(m.BakerConsensusKey)
}

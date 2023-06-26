package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_013_PtJakart"
)

type UnsignedProtocolBlockHeader = proto_013_PtJakart.UnsignedProtocolBlockHeader
type UnsignedBlockHeader = proto_013_PtJakart.UnsignedBlockHeader

type BlockInfo struct {
	ChainID    *tz.ChainID                          `json:"chain_id"`
	Hash       *tz.BlockHash                        `json:"hash"`
	Header     BlockHeader                          `tz:"dyn" json:"header"`
	Metadata   tz.Option[BlockMetadata]             `json:"metadata"`
	Operations []core.OperationsList[GroupContents] `tz:"dyn" json:"operations"`
}

func (block *BlockInfo) GetChainID() *tz.ChainID     { return block.ChainID }
func (block *BlockInfo) GetHash() *tz.BlockHash      { return block.Hash }
func (block *BlockInfo) GetHeader() core.BlockHeader { return &block.Header }
func (block *BlockInfo) GetMetadata() tz.Option[core.BlockMetadata] {
	if m, ok := block.Metadata.CheckUnwrapPtr(); ok {
		return tz.Some[core.BlockMetadata](m)
	}
	return tz.None[core.BlockMetadata]()
}

func (block *BlockInfo) GetOperations() [][]core.OperationsGroup {
	out := make([][]core.OperationsGroup, len(block.Operations))
	for i, list := range block.Operations {
		out[i] = list.GetGroups()
	}
	return out
}

type BlockMetadata struct {
	BlockMetadataContents `tz:"dyn"`
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature tz.AnySignature `json:"signature"`
}

func (header *BlockHeader) GetShellHeader() *core.ShellHeader {
	return &header.ShellHeader
}

func (header *BlockHeader) GetSignature() (tz.Signature, error) {
	return header.Signature.Signature()
}

type BlockHeaderInfo struct {
	ChainID *tz.ChainID   `json:"chain_id"`
	Hash    *tz.BlockHash `json:"hash"`
	BlockHeader
}

func (block *BlockHeaderInfo) GetChainID() *tz.ChainID { return block.ChainID }
func (block *BlockHeaderInfo) GetHash() *tz.BlockHash  { return block.Hash }

type BlockMetadataContents struct {
	core.BlockMetadataHeader
	Proposer                  tz.PublicKeyHash                   `json:"proposer"`
	Baker                     tz.PublicKeyHash                   `json:"baker"`
	LevelInfo                 core.LevelInfo                     `json:"level_info"`
	VotingPeriodInfo          core.VotingPeriodInfo              `json:"voting_period_info"`
	NonceHash                 tz.Option1[*tz.CycleNonceHash]     `json:"nonce_hash"`
	Deactivated               []tz.PublicKeyHash                 `tz:"dyn" json:"deactivated"`
	BalanceUpdates            []*BalanceUpdate                   `tz:"dyn" json:"balance_updates"`
	LiquidityBakingToggleEMA  int32                              `json:"liquidity_baking_toggle_ema"`
	ImplicitOperationsResults []SuccessfulManagerOperationResult `tz:"dyn" json:"implicit_operations_results"`
	ProposerConsensusKey      tz.PublicKeyHash                   `json:"proposer_consensus_key"`
	BakerConsensusKey         tz.PublicKeyHash                   `json:"baker_consensus_key"`
	ConsumedMilligas          tz.BigUint                         `json:"consumed_milligas"`
	DALAttestation            tz.Option[tz.BigInt]               `json:"dal_attestation"`
}

func (m *BlockMetadata) GetMetadataHeader() *core.BlockMetadataHeader { return &m.BlockMetadataHeader }
func (m *BlockMetadata) GetProposer() tz.PublicKeyHash                { return m.Proposer }
func (m *BlockMetadata) GetBaker() tz.PublicKeyHash                   { return m.Baker }
func (m *BlockMetadata) GetLevelInfo() *core.LevelInfo                { return &m.LevelInfo }
func (m *BlockMetadata) GetVotingPeriodInfo() *core.VotingPeriodInfo  { return &m.VotingPeriodInfo }
func (m *BlockMetadata) GetNonceHash() tz.Option[*tz.CycleNonceHash]  { return m.NonceHash.Option }
func (m *BlockMetadata) GetConsumedGas() tz.Option[tz.BigUint]        { return tz.None[tz.BigUint]() }
func (m *BlockMetadata) GetConsumedMilligas() tz.Option[tz.BigUint] {
	return tz.Some(m.ConsumedMilligas)
}
func (m *BlockMetadata) GetDeactivated() []tz.PublicKeyHash { return m.Deactivated }
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
func (m *BlockMetadata) GetProposerConsensusKey() tz.Option[tz.PublicKeyHash] {
	return tz.Some(m.ProposerConsensusKey)
}
func (m *BlockMetadata) GetBakerConsensusKey() tz.Option[tz.PublicKeyHash] {
	return tz.Some(m.BakerConsensusKey)
}

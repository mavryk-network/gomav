package proto_013_PtJakart

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol/core"
)

type UnsignedProtocolBlockHeader struct {
	PayloadHash               *mv.BlockPayloadHash           `json:"payload_hash"`
	PayloadRound              int32                          `json:"payload_round"`
	ProofOfWorkNonce          *mv.Bytes8                     `json:"proof_of_work_nonce"`
	SeedNonceHash             mv.Option[*mv.CycleNonceHash]  `json:"seed_nonce_hash"`
	LiquidityBakingToggleVote core.LiquidityBakingToggleVote `json:"liquidity_baking_toggle_vote"`
}

func (h *UnsignedProtocolBlockHeader) GetPayloadHash() *mv.BlockPayloadHash { return h.PayloadHash }
func (h *UnsignedProtocolBlockHeader) GetPayloadRound() int32               { return h.PayloadRound }
func (h *UnsignedProtocolBlockHeader) GetProofOfWorkNonce() *mv.Bytes8 {
	return h.ProofOfWorkNonce
}
func (h *UnsignedProtocolBlockHeader) GetSeedNonceHash() mv.Option[*mv.CycleNonceHash] {
	return h.SeedNonceHash
}

func (h *UnsignedProtocolBlockHeader) GetLiquidityBakingToggleVote() core.LiquidityBakingToggleVote {
	return h.LiquidityBakingToggleVote
}

type UnsignedBlockHeader struct {
	core.ShellHeader
	UnsignedProtocolBlockHeader
}

func (header *UnsignedBlockHeader) GetShellHeader() *core.ShellHeader {
	return &header.ShellHeader
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature *mv.GenericSignature `json:"signature"`
}

func (header *BlockHeader) GetSignature() (mv.Signature, error) {
	return header.Signature, nil
}

type BlockHeaderInfo struct {
	ChainID *mv.ChainID   `json:"chain_id"`
	Hash    *mv.BlockHash `json:"hash"`
	BlockHeader
}

func (block *BlockHeaderInfo) GetChainID() *mv.ChainID { return block.ChainID }
func (block *BlockHeaderInfo) GetHash() *mv.BlockHash  { return block.Hash }

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
	ConsumedGas               mv.BigUint                         `json:"consumed_gas"`
	Deactivated               []mv.PublicKeyHash                 `mv:"dyn" json:"deactivated"`
	BalanceUpdates            []*BalanceUpdate                   `mv:"dyn" json:"balance_updates"`
	LiquidityBakingToggleEMA  int32                              `json:"liquidity_baking_toggle_ema"`
	ImplicitOperationsResults []SuccessfulManagerOperationResult `mv:"dyn" json:"implicit_operations_results"`
	ConsumedMilligas          mv.BigUint                         `json:"consumed_milligas"`
}

func (m *BlockMetadata) GetMetadataHeader() *core.BlockMetadataHeader { return &m.BlockMetadataHeader }
func (m *BlockMetadata) GetProposer() mv.PublicKeyHash                { return m.Proposer }
func (m *BlockMetadata) GetBaker() mv.PublicKeyHash                   { return m.Baker }
func (m *BlockMetadata) GetLevelInfo() *core.LevelInfo                { return &m.LevelInfo }
func (m *BlockMetadata) GetVotingPeriodInfo() *core.VotingPeriodInfo  { return &m.VotingPeriodInfo }
func (m *BlockMetadata) GetNonceHash() mv.Option[*mv.CycleNonceHash]  { return m.NonceHash.Option }
func (m *BlockMetadata) GetConsumedGas() mv.Option[mv.BigUint]        { return mv.Some(m.ConsumedGas) }
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
	return mv.None[mv.PublicKeyHash]()
}
func (m *BlockMetadata) GetBakerConsensusKey() mv.Option[mv.PublicKeyHash] {
	return mv.None[mv.PublicKeyHash]()
}
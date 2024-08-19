package proto_012_Psithaca

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
)

type Constants struct {
	ProofOfWorkNonceSize                             uint8             `json:"proof_of_work_nonce_size"`
	NonceLength                                      uint8             `json:"nonce_length"`
	MaxAnonOpsPerBlock                               uint8             `json:"max_anon_ops_per_block"`
	MaxOperationDataLength                           int32             `json:"max_operation_data_length"`
	MaxProposalsPerDelegate                          uint8             `json:"max_proposals_per_delegate"`
	MaxMichelineNodeCount                            int32             `json:"max_micheline_node_count"`
	MaxMichelineBytesLimit                           int32             `json:"max_micheline_bytes_limit"`
	MaxAllowedGlobalConstantsDepth                   int32             `json:"max_allowed_global_constants_depth"`
	CacheLayout                                      []int64           `mv:"dyn" json:"cache_layout"`
	MichelsonMaximumTypeSize                         uint16            `json:"michelson_maximum_type_size"`
	PreservedCycles                                  uint8             `json:"preserved_cycles"`
	BlocksPerCycle                                   int32             `json:"blocks_per_cycle"`
	BlocksPerCommitment                              int32             `json:"blocks_per_commitment"`
	BlocksPerStakeSnapshot                           int32             `json:"blocks_per_stake_snapshot"`
	BlocksPerVotingPeriod                            int32             `json:"blocks_per_voting_period"`
	HardGasLimitPerOperation                         mv.BigInt         `json:"hard_gas_limit_per_operation"`
	HardGasLimitPerBlock                             mv.BigInt         `json:"hard_gas_limit_per_block"`
	ProofOfWorkThreshold                             int64             `json:"proof_of_work_threshold"`
	TokensPerRoll                                    mv.BigUint        `json:"tokens_per_roll"`
	SeedNonceRevelationTip                           mv.BigUint        `json:"seed_nonce_revelation_tip"`
	OriginationSize                                  int32             `json:"origination_size"`
	BakingRewardFixedPortion                         mv.BigUint        `json:"baking_reward_fixed_portion"`
	BakingRewardBonusPerSlot                         mv.BigUint        `json:"baking_reward_bonus_per_slot"`
	EndorsingRewardPerSlot                           mv.BigUint        `json:"endorsing_reward_per_slot"`
	CostPerByte                                      mv.BigUint        `json:"cost_per_byte"`
	HardStorageLimitPerOperation                     mv.BigInt         `json:"hard_storage_limit_per_operation"`
	QuorumMin                                        int32             `json:"quorum_min"`
	QuorumMax                                        int32             `json:"quorum_max"`
	MinProposalQuorum                                int32             `json:"min_proposal_quorum"`
	LiquidityBakingSubsidy                           mv.BigUint        `json:"liquidity_baking_subsidy"`
	LiquidityBakingSunsetLevel                       int32             `json:"liquidity_baking_sunset_level"`
	LiquidityBakingEscapeEmaThreshold                int32             `json:"liquidity_baking_escape_ema_threshold"`
	MaxOperationsTimeToLive                          int16             `json:"max_operations_time_to_live"`
	MinimalBlockDelay                                int64             `json:"minimal_block_delay"`
	DelayIncrementPerRound                           int64             `json:"delay_increment_per_round"`
	ConsensusCommitteeSize                           int32             `json:"consensus_committee_size"`
	ConsensusThreshold                               int32             `json:"consensus_threshold"`
	MinimalParticipationRatio                        core.Rat          `json:"minimal_participation_ratio"`
	MaxSlashingPeriod                                int32             `json:"max_slashing_period"`
	FrozenDepositsPercentage                         int32             `json:"frozen_deposits_percentage"`
	DoubleBakingPunishment                           mv.BigUint        `json:"double_baking_punishment"`
	RatioOfFrozenDepositsSlashedPerDoubleEndorsement core.Rat          `json:"ratio_of_frozen_deposits_slashed_per_double_endorsement"`
	DelegateSelection                                DelegateSelection `json:"delegate_selection"`
}

func (c *Constants) GetProofOfWorkNonceSize() uint8    { return c.ProofOfWorkNonceSize }
func (c *Constants) GetNonceLength() uint8             { return c.NonceLength }
func (c *Constants) GetMaxAnonOpsPerBlock() uint8      { return c.MaxAnonOpsPerBlock }
func (c *Constants) GetMaxOperationDataLength() int32  { return c.MaxOperationDataLength }
func (c *Constants) GetMaxProposalsPerDelegate() uint8 { return c.MaxProposalsPerDelegate }
func (c *Constants) GetMaxMichelineNodeCount() int32   { return c.MaxMichelineNodeCount }
func (c *Constants) GetMaxMichelineBytesLimit() int32  { return c.MaxMichelineBytesLimit }
func (c *Constants) GetMaxAllowedGlobalConstantsDepth() int32 {
	return c.MaxAllowedGlobalConstantsDepth
}
func (c *Constants) GetMichelsonMaximumTypeSize() uint16     { return c.MichelsonMaximumTypeSize }
func (c *Constants) GetPreservedCycles() uint8               { return c.PreservedCycles }
func (c *Constants) GetBlocksPerCycle() int32                { return c.BlocksPerCycle }
func (c *Constants) GetBlocksPerCommitment() int32           { return c.BlocksPerCommitment }
func (c *Constants) GetBlocksPerStakeSnapshot() int32        { return c.BlocksPerStakeSnapshot }
func (c *Constants) GetHardGasLimitPerOperation() mv.BigInt  { return c.HardGasLimitPerOperation }
func (c *Constants) GetHardGasLimitPerBlock() mv.BigInt      { return c.HardGasLimitPerBlock }
func (c *Constants) GetProofOfWorkThreshold() int64          { return c.ProofOfWorkThreshold }
func (c *Constants) GetSeedNonceRevelationTip() mv.BigUint   { return c.SeedNonceRevelationTip }
func (c *Constants) GetOriginationSize() int32               { return c.OriginationSize }
func (c *Constants) GetBakingRewardFixedPortion() mv.BigUint { return c.BakingRewardFixedPortion }
func (c *Constants) GetBakingRewardBonusPerSlot() mv.BigUint { return c.BakingRewardBonusPerSlot }
func (c *Constants) GetEndorsingRewardPerSlot() mv.BigUint   { return c.EndorsingRewardPerSlot }
func (c *Constants) GetCostPerByte() mv.BigUint              { return c.CostPerByte }
func (c *Constants) GetHardStorageLimitPerOperation() mv.BigInt {
	return c.HardStorageLimitPerOperation
}
func (c *Constants) GetQuorumMin() int32                     { return c.QuorumMin }
func (c *Constants) GetQuorumMax() int32                     { return c.QuorumMax }
func (c *Constants) GetMinProposalQuorum() int32             { return c.MinProposalQuorum }
func (c *Constants) GetLiquidityBakingSubsidy() mv.BigUint   { return c.LiquidityBakingSubsidy }
func (c *Constants) GetMaxOperationsTimeToLive() int16       { return c.MaxOperationsTimeToLive }
func (c *Constants) GetMinimalBlockDelay() int64             { return c.MinimalBlockDelay }
func (c *Constants) GetDelayIncrementPerRound() int64        { return c.DelayIncrementPerRound }
func (c *Constants) GetConsensusCommitteeSize() int32        { return c.ConsensusCommitteeSize }
func (c *Constants) GetConsensusThreshold() int32            { return c.ConsensusThreshold }
func (c *Constants) GetMinimalParticipationRatio() *core.Rat { return &c.MinimalParticipationRatio }
func (c *Constants) GetMaxSlashingPeriod() int32             { return c.MaxSlashingPeriod }
func (c *Constants) GetFrozenDepositsPercentage() int32      { return c.FrozenDepositsPercentage }
func (c *Constants) GetDoubleBakingPunishment() mv.BigUint   { return c.DoubleBakingPunishment }
func (c *Constants) GetRatioOfFrozenDepositsSlashedPerDoubleEndorsement() *core.Rat {
	return &c.RatioOfFrozenDepositsSlashedPerDoubleEndorsement
}

type DelegateSelection interface {
	DelegateSelection()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[DelegateSelection]{
		Variants: encoding.Variants[DelegateSelection]{
			0: DelegateSelectionRandom{},
			1: DelegateSelectionRoundRobin{},
		},
	})
}

type DelegateSelectionRandom struct{}

func (DelegateSelectionRandom) DelegateSelection() {}

type DelegateSelectionRoundRobin struct {
	Contents []*RoundRobinPublicKeys `mv:"dyn"`
}

func (DelegateSelectionRoundRobin) DelegateSelection() {}

type RoundRobinPublicKeys struct {
	PublicKeys []mv.PublicKey `mv:"dyn"`
}

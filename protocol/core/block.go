package core

import (
	"strconv"

	mv "github.com/mavryk-network/gomav/v2"
)

type LiquidityBakingToggleVote int8

const (
	LiquidityBakingOn LiquidityBakingToggleVote = iota
	LiquidityBakingOff
	LiquidityBakingPass
)

type LevelInfo struct {
	Level              int32 `json:"level"`
	LevelPosition      int32 `json:"level_position"`
	Cycle              int32 `json:"cycle"`
	CyclePosition      int32 `json:"cycle_position"`
	ExpectedCommitment bool  `json:"expected_commitment"`
}

type VotingPeriodInfo struct {
	VotingPeriod VotingPeriod `json:"voting_period"`
	Position     int32        `json:"position"`
	Remaining    int32        `json:"remaining"`
}

type VotingPeriod struct {
	Index         int32            `json:"index"`
	Kind          VotingPeriodKind `json:"kind"`
	StartPosition int32            `json:"start_position"`
}

type VotingPeriodKind uint8

const (
	VotingPeriodProposal VotingPeriodKind = iota
	VotingPeriodExploration
	VotingPeriodCooldown
	VotingPeriodPromotion
	VotingPeriodAdoption
)

func (k VotingPeriodKind) MarshalText() (text []byte, err error) {
	return []byte(k.String()), nil
}

func (k VotingPeriodKind) String() string {
	switch k {
	case VotingPeriodProposal:
		return "proposal"
	case VotingPeriodExploration:
		return "exploration"
	case VotingPeriodCooldown:
		return "cooldown"
	case VotingPeriodPromotion:
		return "promotion"
	case VotingPeriodAdoption:
		return "adoption"
	default:
		return strconv.FormatInt(int64(k), 10)
	}
}

type BlockMetadata interface {
	WithBalanceUpdates
	GetMetadataHeader() *BlockMetadataHeader
	GetProposer() mv.PublicKeyHash
	GetBaker() mv.PublicKeyHash
	GetLevelInfo() *LevelInfo
	GetVotingPeriodInfo() *VotingPeriodInfo
	GetNonceHash() mv.Option[*mv.CycleNonceHash]
	GetConsumedGas() mv.Option[mv.BigUint]
	GetConsumedMilligas() mv.Option[mv.BigUint]
	GetDeactivated() []mv.PublicKeyHash
	GetImplicitOperationsResults() []SuccessfulManagerOperationResult
	GetProposerConsensusKey() mv.Option[mv.PublicKeyHash]
	GetBakerConsensusKey() mv.Option[mv.PublicKeyHash]
}

type BlockHeader interface {
	Signed
	GetShellHeader() *ShellHeader
	GetPayloadHash() *mv.BlockPayloadHash
	GetPayloadRound() int32
	GetProofOfWorkNonce() *mv.Bytes8
	GetSeedNonceHash() mv.Option[*mv.CycleNonceHash]
}

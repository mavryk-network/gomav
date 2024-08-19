package core

import mv "github.com/mavryk-network/gomav/v2"

type DelegatesList struct {
	Delegates []mv.PublicKeyHash `mv:"dyn" json:"delegates"`
}

type PendingConsensusKey interface {
	GetCycle() int32
	GetPKH() mv.PublicKeyHash
}

type DelegateInfo interface {
	GetFullBalance() mv.BigUint
	GetCurrentFrozenDeposits() mv.BigUint
	GetFrozenDeposits() mv.BigUint
	GetStakingBalance() mv.BigUint
	GetFrozenDepositsLimit() mv.Option[mv.BigUint]
	GetDelegatedContracts() []ContractID
	GetDelegatedBalance() mv.BigUint
	GetDeactivated() bool
	GetGracePeriod() int32
	GetVotingPower() mv.Option[int64]
	GetCurrentBallot() mv.Option[BallotKind]
	GetCurrentProposals() mv.Option[[]*mv.ProtocolHash]
	GetRemainingProposals() mv.Option[int32]
	GetActiveConsensusKey() mv.Option[mv.PublicKeyHash]
	GetPendingConsensusKeys() mv.Option[[]PendingConsensusKey]
}

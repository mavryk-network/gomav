package proto_012_Psithaca

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol/core"
)

type DelegateInfo struct {
	FullBalance           mv.BigUint            `json:"full_balance"`
	CurrentFrozenDeposits mv.BigUint            `json:"current_frozen_deposits"`
	FrozenDeposits        mv.BigUint            `json:"frozen_deposits"`
	StakingBalance        mv.BigUint            `json:"staking_balance"`
	FrozenDepositsLimit   mv.Option[mv.BigUint] `json:"frozen_deposits_limit"`
	DelegatedContracts    []core.ContractID     `mv:"dyn" json:"delegated_contracts"`
	DelegatedBalance      mv.BigUint            `json:"delegated_balance"`
	Deactivated           bool                  `json:"deactivated"`
	GracePeriod           int32                 `json:"grace_period"`
	VotingPower           int32                 `json:"voting_power"`
}

func (d *DelegateInfo) GetFullBalance() mv.BigUint                    { return d.FullBalance }
func (d *DelegateInfo) GetCurrentFrozenDeposits() mv.BigUint          { return d.CurrentFrozenDeposits }
func (d *DelegateInfo) GetFrozenDeposits() mv.BigUint                 { return d.FrozenDeposits }
func (d *DelegateInfo) GetStakingBalance() mv.BigUint                 { return d.StakingBalance }
func (d *DelegateInfo) GetFrozenDepositsLimit() mv.Option[mv.BigUint] { return d.FrozenDepositsLimit }
func (d *DelegateInfo) GetDelegatedContracts() []core.ContractID      { return d.DelegatedContracts }
func (d *DelegateInfo) GetDelegatedBalance() mv.BigUint               { return d.DelegatedBalance }
func (d *DelegateInfo) GetDeactivated() bool                          { return d.Deactivated }
func (d *DelegateInfo) GetGracePeriod() int32                         { return d.GracePeriod }
func (d *DelegateInfo) GetVotingPower() mv.Option[int64]              { return mv.Some(int64(d.VotingPower)) }
func (d *DelegateInfo) GetCurrentBallot() mv.Option[core.BallotKind] {
	return mv.None[core.BallotKind]()
}
func (d *DelegateInfo) GetCurrentProposals() mv.Option[[]*mv.ProtocolHash] {
	return mv.None[[]*mv.ProtocolHash]()
}
func (d *DelegateInfo) GetRemainingProposals() mv.Option[int32] { return mv.None[int32]() }
func (d *DelegateInfo) GetActiveConsensusKey() mv.Option[mv.PublicKeyHash] {
	return mv.None[mv.PublicKeyHash]()
}
func (d *DelegateInfo) GetPendingConsensusKeys() mv.Option[[]core.PendingConsensusKey] {
	return mv.None[[]core.PendingConsensusKey]()
}

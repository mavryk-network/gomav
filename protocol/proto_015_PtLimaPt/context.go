package proto_015_PtLimaPt

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol/core"
)

type DelegateInfo struct {
	FullBalance           mv.BigUint                 `json:"full_balance"`
	CurrentFrozenDeposits mv.BigUint                 `json:"current_frozen_deposits"`
	FrozenDeposits        mv.BigUint                 `json:"frozen_deposits"`
	StakingBalance        mv.BigUint                 `json:"staking_balance"`
	FrozenDepositsLimit   mv.Option[mv.BigUint]      `json:"frozen_deposits_limit"`
	DelegatedContracts    []core.ContractID          `mv:"dyn" json:"delegated_contracts,omitempty"`
	DelegatedBalance      mv.BigUint                 `json:"delegated_balance"`
	Deactivated           bool                       `json:"deactivated"`
	GracePeriod           int32                      `json:"grace_period"`
	VotingPower           mv.Option[int64]           `json:"voting_power"`
	CurrentBallot         mv.Option[core.BallotKind] `json:"current_ballot"`
	CurrentProposals      []*mv.ProtocolHash         `mv:"dyn" json:"current_proposals,omitempty"`
	RemainingProposals    int32                      `json:"remaining_proposals"`
	ActiveConsensusKey    mv.PublicKeyHash           `json:"active_consensus_key"`
	PendingConsensusKeys  []*PendingConsensusKey     `mv:"dyn" json:"pending_consensus_keys,omitempty"`
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
func (d *DelegateInfo) GetVotingPower() mv.Option[int64]              { return d.VotingPower }
func (d *DelegateInfo) GetCurrentBallot() mv.Option[core.BallotKind]  { return d.CurrentBallot }
func (d *DelegateInfo) GetCurrentProposals() mv.Option[[]*mv.ProtocolHash] {
	return mv.Some(d.CurrentProposals)
}
func (d *DelegateInfo) GetRemainingProposals() mv.Option[int32] { return mv.Some(d.RemainingProposals) }
func (d *DelegateInfo) GetActiveConsensusKey() mv.Option[mv.PublicKeyHash] {
	return mv.Some(d.ActiveConsensusKey)
}
func (d *DelegateInfo) GetPendingConsensusKeys() mv.Option[[]core.PendingConsensusKey] {
	keys := make([]core.PendingConsensusKey, len(d.PendingConsensusKeys))
	for i, k := range d.PendingConsensusKeys {
		keys[i] = k
	}
	return mv.Some(keys)
}

type PendingConsensusKey struct {
	Cycle int32            `json:"cycle"`
	PKH   mv.PublicKeyHash `json:"pkh"`
}

func (k *PendingConsensusKey) GetCycle() int32          { return k.Cycle }
func (k *PendingConsensusKey) GetPKH() mv.PublicKeyHash { return k.PKH }

package proto_015_PtLimaPt

import (
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/v2/protocol/proto_014_PtKathma"
)

type BalanceUpdates struct {
	BalanceUpdates []*BalanceUpdate `tz:"dyn" json:"balance_updates"`
}

func (b *BalanceUpdates) GetBalanceUpdates() []core.BalanceUpdate {
	out := make([]core.BalanceUpdate, len(b.BalanceUpdates))
	for i, u := range b.BalanceUpdates {
		out[i] = u
	}
	return out
}

type BalanceUpdate struct {
	Contents BalanceUpdateContents    `json:"contents"`
	Change   int64                    `json:"change"`
	Origin   core.BalanceUpdateOrigin `json:"origin"`
}

func (b *BalanceUpdate) GetContents() core.BalanceUpdateContents { return b.Contents }
func (b *BalanceUpdate) GetChange() int64                        { return b.Change }
func (b *BalanceUpdate) GetOrigin() core.BalanceUpdateOrigin     { return b.Origin }

type BalanceUpdateContents interface {
	core.BalanceUpdateContents
}

type BalanceUpdateContract = proto_012_Psithaca.BalanceUpdateContract
type BalanceUpdateDeposits = proto_012_Psithaca.BalanceUpdateDeposits
type BalanceUpdateLostEndorsingRewards = proto_012_Psithaca.BalanceUpdateLostEndorsingRewards
type BalanceUpdateCommitments = proto_012_Psithaca.BalanceUpdateCommitments
type BalanceUpdateBlockFees = proto_012_Psithaca.BalanceUpdateBlockFees
type BalanceUpdateNonceRevelationRewards = proto_012_Psithaca.BalanceUpdateNonceRevelationRewards
type BalanceUpdateDoubleSigningEvidenceRewards = proto_012_Psithaca.BalanceUpdateDoubleSigningEvidenceRewards
type BalanceUpdateEndorsingRewards = proto_012_Psithaca.BalanceUpdateEndorsingRewards
type BalanceUpdateBakingRewards = proto_012_Psithaca.BalanceUpdateBakingRewards
type BalanceUpdateBakingBonuses = proto_012_Psithaca.BalanceUpdateBakingBonuses
type BalanceUpdateStorageFees = proto_012_Psithaca.BalanceUpdateStorageFees
type BalanceUpdateDoubleSigningPunishments = proto_012_Psithaca.BalanceUpdateDoubleSigningPunishments
type BalanceUpdateLiquidityBakingSubsidies = proto_012_Psithaca.BalanceUpdateLiquidityBakingSubsidies
type BalanceUpdateBurned = proto_012_Psithaca.BalanceUpdateBurned
type BalanceUpdateBootstrap = proto_012_Psithaca.BalanceUpdateBootstrap
type BalanceUpdateInvoice = proto_012_Psithaca.BalanceUpdateInvoice
type BalanceUpdateInitialCommitments = proto_012_Psithaca.BalanceUpdateInitialCommitments
type BalanceUpdateMinted = proto_012_Psithaca.BalanceUpdateMinted
type BalanceUpdateTxRollupRejectionRewards = proto_013_PtJakart.BalanceUpdateTxRollupRejectionRewards
type BalanceUpdateTxRollupRejectionPunishments = proto_013_PtJakart.BalanceUpdateTxRollupRejectionPunishments
type BalanceUpdateScRollupRefutationPunishments = proto_014_PtKathma.BalanceUpdateScRollupRefutationPunishments
type BalanceUpdateFrozenBonds = proto_014_PtKathma.BalanceUpdateFrozenBonds

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateScRollupRefutationRewards struct{}

func (BalanceUpdateScRollupRefutationRewards) BalanceUpdateCategory() string {
	return "smart_rollup_refutation_rewards"
}
func (BalanceUpdateScRollupRefutationRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateKindMinted
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BalanceUpdateContents]{
		Variants: encoding.Variants[BalanceUpdateContents]{
			0:  (*BalanceUpdateContract)(nil),
			2:  BalanceUpdateBlockFees{},
			4:  (*BalanceUpdateDeposits)(nil),
			5:  BalanceUpdateNonceRevelationRewards{},
			6:  BalanceUpdateDoubleSigningEvidenceRewards{},
			7:  BalanceUpdateEndorsingRewards{},
			8:  BalanceUpdateBakingRewards{},
			9:  BalanceUpdateBakingBonuses{},
			11: BalanceUpdateStorageFees{},
			12: BalanceUpdateDoubleSigningPunishments{},
			13: (*BalanceUpdateLostEndorsingRewards)(nil),
			14: BalanceUpdateLiquidityBakingSubsidies{},
			15: BalanceUpdateBurned{},
			16: (*BalanceUpdateCommitments)(nil),
			17: BalanceUpdateBootstrap{},
			18: BalanceUpdateInvoice{},
			19: BalanceUpdateInitialCommitments{},
			20: BalanceUpdateMinted{},
			21: (*BalanceUpdateFrozenBonds)(nil),
			22: BalanceUpdateTxRollupRejectionRewards{},
			23: BalanceUpdateTxRollupRejectionPunishments{},
			24: BalanceUpdateScRollupRefutationPunishments{},
			25: BalanceUpdateScRollupRefutationRewards{},
		},
	})
}

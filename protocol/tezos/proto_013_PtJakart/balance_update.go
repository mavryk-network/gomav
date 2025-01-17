package proto_013_PtJakart

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_012_Psithaca"
)

type BalanceUpdates struct {
	BalanceUpdates []*BalanceUpdate `mv:"dyn" json:"balance_updates"`
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

type BondID interface {
	BondID()
}

type TxRollupBondID struct {
	Address *mv.TXRollupAddress `json:"address"`
}

func (TxRollupBondID) BondID() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BondID]{
		Variants: encoding.Variants[BondID]{
			0: TxRollupBondID{},
		},
	})
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateFrozenBonds struct {
	Contract core.ContractID `json:"contract"`
	BondID   BondID          `json:"bond_id"`
}

func (*BalanceUpdateFrozenBonds) BalanceUpdateCategory() string { return "frozen_bonds" }
func (*BalanceUpdateFrozenBonds) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateKindFreezer
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateTxRollupRejectionRewards struct{}

func (BalanceUpdateTxRollupRejectionRewards) BalanceUpdateCategory() string {
	return "tx_rollup_rejection_rewards"
}
func (BalanceUpdateTxRollupRejectionRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateKindMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateTxRollupRejectionPunishments struct{}

func (BalanceUpdateTxRollupRejectionPunishments) BalanceUpdateCategory() string {
	return "tx_rollup_rejection_punishments"
}
func (BalanceUpdateTxRollupRejectionPunishments) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateKindBurned
}

type BalanceUpdateContents interface {
	core.BalanceUpdateContents
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
		},
	})
}

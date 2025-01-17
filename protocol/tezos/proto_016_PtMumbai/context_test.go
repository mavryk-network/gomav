package proto_016_PtMumbai

import (
	"fmt"
	"testing"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDelegateInfo(t *testing.T) {
	src := []byte{
		0x00, 0x00, 0x00, 0x99, 0xaf, 0x96, 0xb1, 0x9b, 0xba, 0x9d, 0x01, 0xb7, 0xa0, 0xbb, 0xd6, 0x95,
		0x32, 0xb7, 0xa0, 0xbb, 0xd6, 0x95, 0x32, 0x97, 0x8d, 0x97, 0xa7, 0xec, 0xf5, 0x03, 0x00, 0x00,
		0x00, 0x00, 0x42, 0x01, 0x13, 0x52, 0x0f, 0x61, 0x45, 0xbf, 0x8a, 0xb7, 0x85, 0x43, 0x86, 0xe3,
		0x73, 0xe0, 0xf1, 0x4d, 0xff, 0x72, 0xa2, 0xd7, 0x00, 0x00, 0x00, 0xe0, 0x69, 0x63, 0x44, 0x45,
		0xd2, 0xf4, 0xf8, 0x36, 0x61, 0xcf, 0x46, 0x65, 0xb6, 0x92, 0xd1, 0x98, 0x05, 0x59, 0xfd, 0x00,
		0x00, 0x84, 0x97, 0x32, 0x8c, 0xe7, 0x20, 0x4d, 0x44, 0x78, 0xbb, 0xa7, 0xdd, 0x56, 0xf3, 0xfb,
		0xa8, 0x01, 0x05, 0x73, 0x35, 0xe8, 0xf6, 0xe5, 0x8b, 0xb2, 0xd8, 0x02, 0x00, 0x00, 0x00, 0x02,
		0x6d, 0xff, 0x00, 0x00, 0x0f, 0xaa, 0xe4, 0xfb, 0xe4, 0x48, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x69, 0x63, 0x44, 0x45, 0xd2, 0xf4, 0xf8, 0x36, 0x61, 0xcf,
		0x46, 0x65, 0xb6, 0x92, 0xd1, 0x98, 0x05, 0x59, 0xfd, 0x00, 0x00, 0x00, 0x00,
	}

	expect := DelegateInfo{
		FullBalance:           mv.BigUint{0xaf, 0x96, 0xb1, 0x9b, 0xba, 0x9d, 0x01},
		CurrentFrozenDeposits: mv.BigUint{0xb7, 0xa0, 0xbb, 0xd6, 0x95, 0x32},
		FrozenDeposits:        mv.BigUint{0xb7, 0xa0, 0xbb, 0xd6, 0x95, 0x32},
		StakingBalance:        mv.BigUint{0x97, 0x8d, 0x97, 0xa7, 0xec, 0xf5, 0x03},
		DelegatedContracts: []core.ContractID{
			core.OriginatedContract{
				ContractHash: &mv.ContractHash{
					0x13, 0x52, 0x0f, 0x61, 0x45, 0xbf, 0x8a, 0xb7, 0x85, 0x43, 0x86, 0xe3, 0x73, 0xe0, 0xf1, 0x4d,
					0xff, 0x72, 0xa2, 0xd7,
				},
			},
			core.ImplicitContract{
				PublicKeyHash: &mv.Ed25519PublicKeyHash{
					0xe0, 0x69, 0x63, 0x44, 0x45, 0xd2, 0xf4, 0xf8, 0x36, 0x61, 0xcf, 0x46, 0x65, 0xb6, 0x92, 0xd1,
					0x98, 0x05, 0x59, 0xfd,
				},
			},
			core.ImplicitContract{
				PublicKeyHash: &mv.Ed25519PublicKeyHash{
					0x84, 0x97, 0x32, 0x8c, 0xe7, 0x20, 0x4d, 0x44, 0x78, 0xbb, 0xa7, 0xdd, 0x56, 0xf3, 0xfb, 0xa8,
					0x01, 0x05, 0x73, 0x35,
				},
			},
		},
		DelegatedBalance: mv.BigUint{
			0xe8, 0xf6, 0xe5, 0x8b, 0xb2, 0xd8, 0x02,
		},
		Deactivated:        false,
		GracePeriod:        621,
		VotingPower:        mv.Some(int64(17226660570184)),
		CurrentBallot:      mv.Some(core.BallotYay),
		CurrentProposals:   []*mv.ProtocolHash{},
		RemainingProposals: 0,
		ActiveConsensusKey: &mv.Ed25519PublicKeyHash{
			0xe0, 0x69, 0x63, 0x44, 0x45, 0xd2, 0xf4, 0xf8, 0x36, 0x61, 0xcf, 0x46, 0x65, 0xb6, 0x92, 0xd1,
			0x98, 0x05, 0x59, 0xfd,
		},
		PendingConsensusKeys: []*PendingConsensusKey{},
	}

	var out DelegateInfo
	_, err := encoding.Decode(src, &out, encoding.Dynamic())
	if !assert.NoError(t, err) {
		if err, ok := err.(*encoding.Error); ok {
			fmt.Println(err.Path)
		}
	} else {
		require.Equal(t, &expect, &out)
	}
}

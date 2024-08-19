package protocol_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type protoTestData struct {
	proto  *mv.ProtocolHash
	blocks []string
}

var testData = []protoTestData{
	{
		proto: &core.Proto002PtBoreas,
		blocks: []string{
			"83409",
			"83410",
			"83411",
			"83412",
			"83413",
			"83414",
			"83415",
			"83416",
			"83417",
			"83418",
		},
	},
	{
		proto: &core.Proto001PtAtLas,
		blocks: []string{
			"494220",
			"494221",
			"494222",
			"494223",
		},
	},
}

func TestBlock(t *testing.T) {
	for _, protoData := range testData {
		t.Run(protoData.proto.String(), func(t *testing.T) {
			for _, block := range protoData.blocks {
				t.Run(block, func(t *testing.T) {
					fileName := filepath.Join("test_data", core.ProtocolShortName(protoData.proto), block+".bin")
					buf, err := os.ReadFile(fileName)
					require.NoError(t, err)
					out, err := protocol.NewBlockInfo(protoData.proto)
					require.NoError(t, err)
					_, err = encoding.Decode(buf, out, encoding.Dynamic())
					if !assert.NoError(t, err) {
						//pretty.Println(out)
						if err, ok := err.(*encoding.Error); ok {
							fmt.Println(err.Path)
						}
						return
					}

					// check operation group signature lengths
					for _, list := range out.GetOperations() {
						for _, grp := range list {
							_, err := grp.GetContents().GetSignature()
							assert.NoError(t, err)
						}
					}
				})
			}
		})
	}
}

var headerTestData = []protoTestData{
	{
		proto: &core.Proto001PtAtLas,
		blocks: []string{
			"3279466",
		},
	},
}

func TestBlockHeader(t *testing.T) {
	for _, protoData := range headerTestData {
		t.Run(protoData.proto.String(), func(t *testing.T) {
			for _, block := range protoData.blocks {
				t.Run(block, func(t *testing.T) {
					fileName := filepath.Join("test_data", core.ProtocolShortName(protoData.proto), "header_"+block+".bin")
					buf, err := os.ReadFile(fileName)
					require.NoError(t, err)
					out, err := protocol.NewBlockHeaderInfo(protoData.proto)
					require.NoError(t, err)
					_, err = encoding.Decode(buf, out, encoding.Dynamic())
					if !assert.NoError(t, err) {
						if err, ok := err.(*encoding.Error); ok {
							fmt.Println(err.Path)
						}
					}
				})
			}
		})
	}
}

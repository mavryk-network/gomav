package core

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
)

type ShellHeader struct {
	Level          int32              `json:"level"`
	Proto          uint8              `json:"proto"`
	Predecessor    *mv.BlockHash      `json:"predecessor"`
	Timestamp      mv.Timestamp       `json:"timestamp"`
	ValidationPass uint8              `json:"validation_pass"`
	OperationsHash *mv.OperationsHash `json:"operations_hash"`
	Fitness        mv.Bytes           `mv:"dyn" json:"fitness"`
	Context        *mv.ContextHash    `json:"context"`
}

type BlockMetadataHeader struct {
	TestChainStatus        TestChainStatus           `json:"test_chain_status"`
	MaxOperationsTTL       int32                     `json:"max_operations_ttl"`
	MaxOperationDataLength int32                     `json:"max_operation_data_length"`
	MaxBlockHeaderLength   int32                     `json:"max_block_header_length"`
	MaxOperationListLength []*MaxOperationListLength `mv:"dyn,dyn" json:"max_operation_list_length"`
}

func (*BlockMetadataHeader) BlockMetadataContents() {}

type TestChainStatus interface {
	TestChainStatus() string
}

type TestChainStatusNotRunning struct{}

func (TestChainStatusNotRunning) TestChainStatus() string { return "not_running" }

func (t TestChainStatusNotRunning) MarshalText() (text []byte, err error) {
	return []byte(t.TestChainStatus()), nil
}

//json:status=TestChainStatus()
type TestChainStatusForking struct {
	Protocol   *mv.ProtocolHash `json:"protocol"`
	Expiration int64            `json:"expiration"`
}

func (TestChainStatusForking) TestChainStatus() string { return "forking" }

//json:status=TestChainStatus()
type TestChainStatusRunning struct {
	ChainID    *mv.ChainID      `json:"chain_id,omitempty"`
	Genesis    *mv.BlockHash    `json:"genesis,omitempty"`
	Protocol   *mv.ProtocolHash `json:"protocol"`
	Expiration int64            `json:"expiration"`
}

func (TestChainStatusRunning) TestChainStatus() string { return "running" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[TestChainStatus]{
		Variants: encoding.Variants[TestChainStatus]{
			0: TestChainStatusNotRunning{},
			1: (*TestChainStatusForking)(nil),
			2: (*TestChainStatusRunning)(nil),
		},
	})
}

type MaxOperationListLength struct {
	MaxSize int32            `json:"max_size"`
	MaxOp   mv.Option[int32] `json:"max_op"`
}

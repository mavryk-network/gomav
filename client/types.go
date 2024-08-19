package client

//go:generate go run generate.go

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/latest"
)

type BlockInfo = protocol.BlockInfo
type BlockHeaderInfo = protocol.BlockHeaderInfo
type BigUint = mv.BigUint
type ChainID = mv.ChainID
type OperationWithOptionalMetadata = latest.OperationWithOptionalMetadata
type Constants = core.Constants
type BlockShellHeader = core.ShellHeader
type OperationHash = mv.OperationHash
type BlockProtocols = core.BlockProtocols
type BlockHash = mv.BlockHash

type MetadataMode int

const (
	MetadataDefault MetadataMode = iota
	MetadataAlways
	MetadataNever
)

func (m MetadataMode) String() string {
	switch m {
	case MetadataAlways:
		return "always"
	case MetadataNever:
		return "never"
	default:
		return "default"
	}
}

type SimpleRequest struct {
	Chain string
	Block string
}

type BlockRequest struct {
	Chain    string
	Block    string
	Metadata MetadataMode
	Protocol *mv.ProtocolHash
}

type ContractRequest struct {
	Chain string
	Block string
	ID    core.ContractID
}

type ContextRequest struct {
	Chain    string
	Block    string
	Protocol *mv.ProtocolHash
}

type RunOperationRequest struct {
	Chain   string
	Block   string
	Payload *latest.RunOperationRequest
}

type InjectOperationRequest struct {
	Chain   string
	Async   Flag
	Payload *InjectRequestPayload
}

type InjectRequestPayload struct {
	Contents []byte `mv:"dyn"`
}

type BasicBlockInfo struct {
	Hash     *mv.BlockHash
	Protocol *mv.ProtocolHash
}

type HeadsRequest struct {
	Chain        string
	Protocol     *mv.ProtocolHash
	NextProtocol *mv.ProtocolHash
}

type Head struct {
	Hash *mv.BlockHash `json:"hash"`
	core.ShellHeader
	ProtocolData []byte `json:"protocol_data"` // not dyn, takes the rest
}

type Flag bool

func newConstants(p *mv.ProtocolHash) (Constants, error) { return protocol.NewConstants(p) }
func newBlockInfo(p *mv.ProtocolHash) (BlockInfo, error) { return protocol.NewBlockInfo(p) }
func newBlockHeaderInfo(p *mv.ProtocolHash) (BlockHeaderInfo, error) {
	return protocol.NewBlockHeaderInfo(p)
}

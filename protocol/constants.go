package protocol

import (
	"fmt"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/proto_001_PtAtLas"
	"github.com/mavryk-network/gomav/v2/protocol/proto_002_PtBoreas"
)

func NewConstants(proto *mv.ProtocolHash) (constants core.Constants, err error) {
	switch *proto {
	case core.Proto002PtBoreas:
		constants = new(proto_002_PtBoreas.Constants)
	case core.Proto001PtAtLas:
		constants = new(proto_001_PtAtLas.Constants)
	default:
		return nil, fmt.Errorf("gomav: NewConstants: unknown protocol version %d", proto)
	}
	return
}

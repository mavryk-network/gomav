package protocol

import (
	"fmt"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/proto_001_PtAtLas"
	"github.com/mavryk-network/gomav/v2/protocol/proto_002_PtBoreas"
)

func NewDelegateInfo(proto *mv.ProtocolHash) (delegate core.DelegateInfo, err error) {
	switch *proto {
	case core.Proto002PtBoreas:
		delegate = new(proto_002_PtBoreas.DelegateInfo)
	case core.Proto001PtAtLas:
		delegate = new(proto_001_PtAtLas.DelegateInfo)
	default:
		return nil, fmt.Errorf("gomav: NewDelegateInfo: unknown protocol version %d", proto)
	}
	return
}

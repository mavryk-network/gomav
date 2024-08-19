package protocol

import (
	"fmt"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/proto_012_Psithaca"
	"github.com/mavryk-network/gomav/v2/protocol/proto_013_PtJakart"
	"github.com/mavryk-network/gomav/v2/protocol/proto_014_PtKathma"
	"github.com/mavryk-network/gomav/v2/protocol/proto_015_PtLimaPt"
	"github.com/mavryk-network/gomav/v2/protocol/proto_016_PtMumbai"
	"github.com/mavryk-network/gomav/v2/protocol/proto_017_PtNairob"
	"github.com/mavryk-network/gomav/v2/protocol/proto_018_Proxford"
	"github.com/mavryk-network/gomav/v2/protocol/proto_019_PtParisB"
)

func NewDelegateInfo(proto *mv.ProtocolHash) (delegate core.DelegateInfo, err error) {
	switch *proto {
	case core.Proto019PtParisB:
		delegate = new(proto_019_PtParisB.DelegateInfo)
	case core.Proto018Proxford:
		delegate = new(proto_018_Proxford.DelegateInfo)
	case core.Proto017PtNairob:
		delegate = new(proto_017_PtNairob.DelegateInfo)
	case core.Proto016PtMumbai:
		delegate = new(proto_016_PtMumbai.DelegateInfo)
	case core.Proto015PtLimaPt:
		delegate = new(proto_015_PtLimaPt.DelegateInfo)
	case core.Proto014PtKathma:
		delegate = new(proto_014_PtKathma.DelegateInfo)
	case core.Proto013PtJakart:
		delegate = new(proto_013_PtJakart.DelegateInfo)
	case core.Proto012Psithaca:
		delegate = new(proto_012_Psithaca.DelegateInfo)
	default:
		return nil, fmt.Errorf("gomav: NewDelegateInfo: unknown protocol version %d", proto)
	}
	return
}

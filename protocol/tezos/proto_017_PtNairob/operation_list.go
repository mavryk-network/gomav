package proto_017_PtNairob

import (
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_016_PtMumbai"
)

type OperationWithOptionalMetadata = core.OperationWithOptionalMetadata[OperationWithOptionalMetadataContents]

type GroupContents interface {
	core.GroupContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[GroupContents]{
		Variants: encoding.Variants[GroupContents]{
			0: (*proto_016_PtMumbai.OperationWithTooLargeMetadata[OperationContents])(nil),
			1: (*proto_016_PtMumbai.OperationWithoutMetadata[OperationContents])(nil),
			2: (*OperationWithOptionalMetadata)(nil),
		},
	})
}

type OperationWithOptionalMetadataContents interface {
	core.OperationWithOptionalMetadataContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationWithOptionalMetadataContents]{
		Variants: encoding.Variants[OperationWithOptionalMetadataContents]{
			0: (*proto_016_PtMumbai.OperationWithOptionalMetadataWithMetadata[OperationContentsAndResult])(nil),
			1: (*proto_016_PtMumbai.OperationWithOptionalMetadataWithoutMetadata[OperationContents])(nil),
		},
	})
}

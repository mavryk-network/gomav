package proto_013_PtJakart

import (
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_012_Psithaca"
)

type GroupContents interface {
	core.GroupContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[GroupContents]{
		Variants: encoding.Variants[GroupContents]{
			0: (*proto_012_Psithaca.OperationWithTooLargeMetadata[OperationContents])(nil),
			1: (*proto_012_Psithaca.OperationWithoutMetadata[OperationContents])(nil),
			2: (*core.OperationWithOptionalMetadata[OperationWithOptionalMetadataContents])(nil),
		},
	})
}

type OperationWithOptionalMetadataContents interface {
	core.OperationWithOptionalMetadataContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationWithOptionalMetadataContents]{
		Variants: encoding.Variants[OperationWithOptionalMetadataContents]{
			0: (*proto_012_Psithaca.OperationWithOptionalMetadataWithMetadata[OperationContentsAndResult])(nil),
			1: (*proto_012_Psithaca.OperationWithOptionalMetadataWithoutMetadata[OperationContents])(nil),
		},
	})
}

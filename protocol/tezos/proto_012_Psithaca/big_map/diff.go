package big_map

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core/expression"
)

//go:generate go run ../../../cmd/genmarshaller.go

type Diff struct {
	Contents []Op `mv:"dyn" json:"contents"`
}

type Op interface {
	BigMapDiffOp() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[Op]{
		Variants: encoding.Variants[Op]{
			0: (*Update)(nil),
			1: (*Remove)(nil),
			2: (*Copy)(nil),
			3: (*Alloc)(nil),
		},
	})
}

//json:action=BigMapDiffOp()
type Update struct {
	BigMap  mv.BigInt                        `json:"big_map"`
	KeyHash *mv.ScriptExprHash               `json:"key_hash"`
	Key     expression.Expression            `json:"key"`
	Value   mv.Option[expression.Expression] `json:"value"`
}

func (*Update) BigMapDiffOp() string { return "update" }

//json:action=BigMapDiffOp()
type Remove struct {
	BigMap mv.BigInt `json:"big_map"`
}

func (*Remove) BigMapDiffOp() string { return "remove" }

//json:action=BigMapDiffOp()
type Copy struct {
	SourceBigMap      mv.BigInt `json:"source_big_map"`
	DestinationBigMap mv.BigInt `json:"destination_big_map"`
}

func (*Copy) BigMapDiffOp() string { return "copy" }

//json:action=BigMapDiffOp()
type Alloc struct {
	BigMap    mv.BigInt             `json:"big_map"`
	KeyType   expression.Expression `json:"key_type"`
	ValueType expression.Expression `json:"value_type"`
}

func (*Alloc) BigMapDiffOp() string { return "alloc" }

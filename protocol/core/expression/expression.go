package expression

import (
	"encoding/json"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
)

type Expression interface {
	Expression()
}

type Int struct {
	Int mv.BigInt `json:"int"`
}

func (Int) Expression() {}

type String struct {
	String string `mv:"dyn" json:"string"`
}

func (String) Expression() {}

type Seq struct {
	Value []Expression `mv:"dyn"`
}

func (s Seq) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Value)
}

func (Seq) Expression() {}

type Prim00 Prim

func (Prim00) Expression() {}

func (p Prim00) MarshalJSON() ([]byte, error) {
	return json.Marshal(&PrimXX{
		Prim: Prim(p),
	})
}

type Prim0X struct {
	Prim   Prim   `json:"prim"`
	Annots string `mv:"dyn" json:"annots"`
}

func (*Prim0X) Expression() {}

type Prim10 struct {
	Prim Prim
	Arg  Expression
}

func (*Prim10) Expression() {}

func (p *Prim10) MarshalJSON() ([]byte, error) {
	return json.Marshal(&PrimXX{
		Prim: p.Prim,
		Args: []Expression{p.Arg},
	})
}

type Prim1X struct {
	Prim   Prim
	Arg    Expression
	Annots string `mv:"dyn"`
}

func (*Prim1X) Expression() {}

func (p *Prim1X) MarshalJSON() ([]byte, error) {
	return json.Marshal(&PrimXX{
		Prim:   p.Prim,
		Args:   []Expression{p.Arg},
		Annots: p.Annots,
	})
}

type Prim20 struct {
	Prim Prim          `json:"prim"`
	Args [2]Expression `json:"args"`
}

func (*Prim20) Expression() {}

type Prim2X struct {
	Prim   Prim          `json:"prim"`
	Args   [2]Expression `json:"args"`
	Annots string        `mv:"dyn" json:"annots,omitempty"`
}

func (*Prim2X) Expression() {}

type PrimXX struct {
	Prim   Prim         `json:"prim"`
	Args   []Expression `mv:"dyn" json:"args,omitempty"`
	Annots string       `mv:"dyn" json:"annots,omitempty"`
}

func (*PrimXX) Expression() {}

type Bytes struct {
	Bytes mv.Bytes `mv:"dyn" json:"bytes"`
}

func (Bytes) Expression() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[Expression]{
		Variants: encoding.Variants[Expression]{
			0:  Int{},
			1:  String{},
			2:  Seq{},
			3:  Prim00(0),
			4:  (*Prim0X)(nil),
			5:  (*Prim10)(nil),
			6:  (*Prim1X)(nil),
			7:  (*Prim20)(nil),
			8:  (*Prim2X)(nil),
			9:  (*PrimXX)(nil),
			10: Bytes{},
		},
	})
}

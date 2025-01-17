//go:build ignore

package main

import (
	"log"
	"os"
	"text/template"
)

type mapping struct {
	Type   string
	Length string
	Prefix string
}

const outName = "b58_gen.go"

var data = []mapping{
	{
		Type:   "BlockHash",
		Length: "BlockHashBytesLen",
		Prefix: "BlockHash",
	},
	{
		Type:   "OperationsHash",
		Length: "OperationListListHashBytesLen",
		Prefix: "OperationListListHash",
	},
	{
		Type:   "ContextHash",
		Length: "ContextHashBytesLen",
		Prefix: "ContextHash",
	},
	{
		Type:   "ChainID",
		Length: "ChainIdBytesLen",
		Prefix: "ChainID",
	},
	{
		Type:   "BlockPayloadHash",
		Length: "BlockPayloadHashBytesLen",
		Prefix: "ValueHash",
	},
	{
		Type:   "CycleNonceHash",
		Length: "CycleNonceBytesLen",
		Prefix: "CycleNonce",
	},
	{
		Type:   "Ed25519PublicKeyHash",
		Length: "PKHBytesLen",
		Prefix: "Ed25519PublicKeyHash",
	},
	{
		Type:   "Secp256k1PublicKeyHash",
		Length: "PKHBytesLen",
		Prefix: "Secp256k1PublicKeyHash",
	},
	{
		Type:   "P256PublicKeyHash",
		Length: "PKHBytesLen",
		Prefix: "P256PublicKeyHash",
	},
	{
		Type:   "BLSPublicKeyHash",
		Length: "PKHBytesLen",
		Prefix: "BLS12_381PublicKeyHash",
	},
	{
		Type:   "ProtocolHash",
		Length: "ProtocolHashBytesLen",
		Prefix: "ProtocolHash",
	},
	{
		Type:   "ContractHash",
		Length: "ContractHashBytesLen",
		Prefix: "ContractHash",
	},
}

const tplSrc = `package b58

import (
	"fmt"

	"github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/b58/base58"
	"github.com/mavryk-network/gomav/v2/b58/prefix"
)

// Code generated by generate.go DO NOT EDIT.
{{range .}}
func Parse{{.Type}}(src []byte) (*gomav.{{.Type}}, error) {
	pre, payload, err := base58.DecodeMV(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.{{.Prefix}} {
		return nil, fmt.Errorf("gomav: invalid {{.Type}} encoding")
	}
	var out gomav.{{.Type}}
	copy(out[:], payload)
	return &out, nil
}
{{end}}
`

var tpl = template.Must(template.New("generate").Parse(tplSrc))

func main() {
	fd, err := os.Create(outName)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	if err = tpl.Execute(fd, data); err != nil {
		log.Fatal(err)
	}
}

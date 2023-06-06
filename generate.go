//go:build ignore

package main

import (
	"log"
	"os"
	"text/template"
)

type base58Type struct {
	Type   string
	Length string
	Prefix string
}

type hexType struct {
	Type   string
	Length string
}

const outName = "types_gen.go"

var base58Data = []base58Type{
	{
		Type:   "BlockHash",
		Length: "HashBytesLen",
		Prefix: "BlockHash",
	},
	{
		Type:   "OperationsHash",
		Length: "HashBytesLen",
		Prefix: "OperationListListHash",
	},
	{
		Type:   "ContextHash",
		Length: "HashBytesLen",
		Prefix: "ContextHash",
	},
	{
		Type:   "ChainID",
		Length: "ChainIdBytesLen",
		Prefix: "ChainID",
	},
	{
		Type:   "BlockPayloadHash",
		Length: "HashBytesLen",
		Prefix: "ValueHash",
	},
	{
		Type:   "CycleNonceHash",
		Length: "CycleNonceBytesLen",
		Prefix: "CycleNonce",
	},
	{
		Type:   "Ed25519PublicKeyHash",
		Length: "AddressBytesLen",
		Prefix: "Ed25519PublicKeyHash",
	},
	{
		Type:   "Secp256k1PublicKeyHash",
		Length: "AddressBytesLen",
		Prefix: "Secp256k1PublicKeyHash",
	},
	{
		Type:   "P256PublicKeyHash",
		Length: "AddressBytesLen",
		Prefix: "P256PublicKeyHash",
	},
	{
		Type:   "BLSPublicKeyHash",
		Length: "AddressBytesLen",
		Prefix: "BLS12_381PublicKeyHash",
	},
	{
		Type:   "ProtocolHash",
		Length: "HashBytesLen",
		Prefix: "ProtocolHash",
	},
	{
		Type:   "ContractHash",
		Length: "AddressBytesLen",
		Prefix: "ContractHash",
	},
	{
		Type:   "Ed25519PublicKey",
		Length: "Ed25519PublicKeyBytesLen",
		Prefix: "Ed25519PublicKey",
	},
	{
		Type:   "Secp256k1PublicKey",
		Length: "Secp256k1PublicKeyBytesLen",
		Prefix: "Secp256k1PublicKey",
	},
	{
		Type:   "P256PublicKey",
		Length: "P256PublicKeyBytesLen",
		Prefix: "P256PublicKey",
	},
	{
		Type:   "BLSPublicKey",
		Length: "BLSPublicKeyBytesLen",
		Prefix: "BLS12_381PublicKey",
	},
	{
		Type:   "Ed25519PrivateKey",
		Length: "Ed25519SeedBytesLen",
		Prefix: "Ed25519Seed",
	},
	{
		Type:   "Secp256k1PrivateKey",
		Length: "Secp256k1PrivateKeyBytesLen",
		Prefix: "Secp256k1SecretKey",
	},
	{
		Type:   "P256PrivateKey",
		Length: "P256PrivateKeyBytesLen",
		Prefix: "P256SecretKey",
	},
	{
		Type:   "BLSPrivateKey",
		Length: "BLSPrivateKeyBytesLen",
		Prefix: "BLS12_381SecretKey",
	},
	{
		Type:   "Ed25519EncryptedPrivateKey",
		Length: "Ed25519EncryptedSeedBytesLen",
		Prefix: "Ed25519EncryptedSeed",
	},
	{
		Type:   "Secp256k1EncryptedPrivateKey",
		Length: "Secp256k1EncryptedPrivateKeyBytesLen",
		Prefix: "Secp256k1EncryptedSecretKey",
	},
	{
		Type:   "P256EncryptedPrivateKey",
		Length: "P256EncryptedPrivateKeyBytesLen",
		Prefix: "P256EncryptedSecretKey",
	},
	{
		Type:   "BLSEncryptedPrivateKey",
		Length: "BLSEncryptedPrivateKeyBytesLen",
		Prefix: "BLS12_381EncryptedSecretKey",
	},
	{
		Type:   "GenericSignature",
		Length: "GenericSignatureBytesLen",
		Prefix: "GenericSignature",
	},
	{
		Type:   "Ed25519Signature",
		Length: "GenericSignatureBytesLen",
		Prefix: "Ed25519Signature",
	},
	{
		Type:   "Secp256k1Signature",
		Length: "GenericSignatureBytesLen",
		Prefix: "Secp256k1Signature",
	},
	{
		Type:   "P256Signature",
		Length: "GenericSignatureBytesLen",
		Prefix: "P256Signature",
	},
	{
		Type:   "BLSSignature",
		Length: "BLSSignatureBytesLen",
		Prefix: "BLS12_381Signature",
	},
	{
		Type:   "BlindedPublicKeyHash",
		Length: "AddressBytesLen",
		Prefix: "BlindedPublicKeyHash",
	},
	{
		Type:   "TXRollupAddress",
		Length: "AddressBytesLen",
		Prefix: "TXRollupAddress",
	},
	{
		Type:   "ZkRollupAddress",
		Length: "AddressBytesLen",
		Prefix: "ZkRollupHash",
	},
	{
		Type:   "DALCommitment",
		Length: "SlotHeaderBytesLen",
		Prefix: "SlotHeader",
	},
	{
		Type:   "ScriptExprHash",
		Length: "HashBytesLen",
		Prefix: "ScriptExpr",
	},
	{
		Type:   "SmartRollupAddress",
		Length: "AddressBytesLen",
		Prefix: "SmartRollupHash",
	},
	{
		Type:   "SmartRollupStateHash",
		Length: "HashBytesLen",
		Prefix: "SmartRollupStateHash",
	},
	{
		Type:   "SmartRollupCommitmentHash",
		Length: "HashBytesLen",
		Prefix: "SmartRollupCommitmentHash",
	},
	{
		Type:   "ScRollupAddress",
		Length: "AddressBytesLen",
		Prefix: "ScRollupHash",
	},
	{
		Type:   "ScRollupStateHash",
		Length: "HashBytesLen",
		Prefix: "ScRollupStateHash",
	},
	{
		Type:   "ScRollupCommitmentHash",
		Length: "HashBytesLen",
		Prefix: "ScRollupCommitmentHash",
	},
}

var hexData = []hexType{
	{
		Type:   "SeedNonce",
		Length: "SeedNonceBytesLen",
	},
	{
		Type:   "CycleNonce",
		Length: "CycleNonceBytesLen",
	},
	{
		Type:   "Bytes32",
		Length: "32",
	},
}

const tplSrc = `package gotez

import (
	"encoding/hex"
	"fmt"
	"github.com/ecadlabs/gotez/v2/b58/base58"
	"github.com/ecadlabs/gotez/v2/b58/prefix"
)

// Code generated by generate.go DO NOT EDIT.
{{range .Base58Types}}
type {{.Type}} [{{.Length}}]byte

func (self *{{.Type}}) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.{{.Prefix}}, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self {{.Type}}) String() string {
	return string(self.ToBase58())
}

func (self {{.Type}}) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.{{.Prefix}}, self[:])
}

func (self *{{.Type}}) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.{{.Prefix}} {
		return fmt.Errorf("gotez: invalid {{.Type}} encoding")
	}
	copy(self[:], payload)
	return nil
}
{{end}}
{{range .HexTypes}}
type {{.Type}} [{{.Length}}]byte

func (self {{.Type}}) String() string {
	out, _ := self.MarshalText()
	return string(out)
}

func (self {{.Type}}) MarshalText() ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(self)))
	hex.Encode(dst, self[:])
	return dst, nil
}

func (self *{{.Type}}) UnmarshalText(src []byte) error {
	hex.Decode(self[:], src)
	return nil
}
{{end}}
`

type tplData struct {
	Base58Types []base58Type
	HexTypes    []hexType
}

var tpl = template.Must(template.New("template").Parse(tplSrc))

func main() {
	fd, err := os.Create(outName)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	data := tplData{
		Base58Types: base58Data,
		HexTypes:    hexData,
	}
	if err = tpl.Execute(fd, &data); err != nil {
		log.Fatal(err)
	}
}

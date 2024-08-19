package core

import (
	"bytes"
	"errors"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/b58/base58"
	"github.com/mavryk-network/gomav/v2/b58/prefix"
	"github.com/mavryk-network/gomav/v2/encoding"
)

type TransactionDestination interface {
	mv.Base58Encoder
	TransactionDestination()
	Eq(other TransactionDestination) bool
}

type Address = TransactionDestination

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionDestination]{
		Variants: encoding.Variants[TransactionDestination]{
			0: ImplicitContract{},
			1: OriginatedContract{},
			2: TxRollupDestination{},
			3: SmartRollupDestination{},
			4: ZkRollupDestination{},
		},
	})
}

type TxRollupDestination struct {
	*mv.TXRollupAddress
	Padding uint8
}

func (TxRollupDestination) TransactionDestination() {}
func (a TxRollupDestination) Eq(b TransactionDestination) bool {
	if other, ok := b.(TxRollupDestination); ok {
		return bytes.Equal(a.TXRollupAddress[:], other.TXRollupAddress[:])
	}
	return false
}

type SmartRollupDestination struct {
	*mv.SmartRollupAddress
	Padding uint8
}

func (SmartRollupDestination) TransactionDestination() {}
func (a SmartRollupDestination) Eq(b TransactionDestination) bool {
	if other, ok := b.(SmartRollupDestination); ok {
		return bytes.Equal(a.SmartRollupAddress[:], other.SmartRollupAddress[:])
	}
	return false
}

type ZkRollupDestination struct {
	*mv.ZkRollupAddress
	Padding uint8
}

func (ZkRollupDestination) TransactionDestination() {}
func (a ZkRollupDestination) Eq(b TransactionDestination) bool {
	if other, ok := b.(ZkRollupDestination); ok {
		return bytes.Equal(a.ZkRollupAddress[:], other.ZkRollupAddress[:])
	}
	return false
}

func ParseTransactionDestination(src []byte) (TransactionDestination, error) {
	pre, payload, err := base58.DecodeMV(src)
	if err != nil {
		return nil, err
	}

	switch pre {
	case &prefix.Ed25519PublicKeyHash:
		var out mv.Ed25519PublicKeyHash
		copy(out[:], payload)
		return ImplicitContract{&out}, nil

	case &prefix.Secp256k1PublicKeyHash:
		var out mv.Secp256k1PublicKeyHash
		copy(out[:], payload)
		return ImplicitContract{&out}, nil

	case &prefix.P256PublicKeyHash:
		var out mv.P256PublicKeyHash
		copy(out[:], payload)
		return ImplicitContract{&out}, nil

	case &prefix.BLS12_381PublicKeyHash:
		var out mv.BLSPublicKeyHash
		copy(out[:], payload)
		return ImplicitContract{&out}, nil

	case &prefix.ContractHash:
		var out mv.ContractHash
		copy(out[:], payload)
		return OriginatedContract{&out, 0}, nil

	case &prefix.TXRollupAddress:
		var out mv.TXRollupAddress
		copy(out[:], payload)
		return TxRollupDestination{&out, 0}, nil

	case &prefix.SmartRollupHash:
		var out mv.SmartRollupAddress
		copy(out[:], payload)
		return SmartRollupDestination{&out, 0}, nil

	case &prefix.ZkRollupHash:
		var out mv.ZkRollupAddress
		copy(out[:], payload)
		return ZkRollupDestination{&out, 0}, nil

	default:
		return nil, errors.New("gomav: unknown destination prefix")
	}
}

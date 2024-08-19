// Package b58 contains Base58 decoding functions of various Mavryk types
package b58

import (
	"errors"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/b58/base58"
	"github.com/mavryk-network/gomav/v2/b58/prefix"
)

//go:generate go run generate.go

func ParsePublicKey(src []byte) (mv.PublicKey, error) {
	pre, payload, err := base58.DecodeMV(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.Ed25519PublicKey:
		var out mv.Ed25519PublicKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1PublicKey:
		var out mv.Secp256k1PublicKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256PublicKey:
		var out mv.P256PublicKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381PublicKey:
		var out mv.BLSPublicKey
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, errors.New("gomav: unknown public key prefix")
	}
}

func ParsePublicKeyHash(src []byte) (mv.PublicKeyHash, error) {
	pre, payload, err := base58.DecodeMV(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.Ed25519PublicKeyHash:
		var out mv.Ed25519PublicKeyHash
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1PublicKeyHash:
		var out mv.Secp256k1PublicKeyHash
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256PublicKeyHash:
		var out mv.P256PublicKeyHash
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381PublicKeyHash:
		var out mv.BLSPublicKeyHash
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, errors.New("gomav: unknown public key prefix")
	}
}

func ParsePrivateKey(src []byte) (mv.PrivateKey, error) {
	pre, payload, err := base58.DecodeMV(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.Ed25519Seed:
		var out mv.Ed25519PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1SecretKey:
		var out mv.Secp256k1PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256SecretKey:
		var out mv.P256PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381SecretKey:
		var out mv.BLSPrivateKey
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, errors.New("gomav: unknown private key prefix")
	}
}

func ParseEncryptedPrivateKey(src []byte) (mv.EncryptedPrivateKey, error) {
	pre, payload, err := base58.DecodeMV(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.Ed25519Seed:
		var out mv.Ed25519PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1SecretKey:
		var out mv.Secp256k1PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256SecretKey:
		var out mv.P256PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381SecretKey:
		var out mv.BLSPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.Ed25519EncryptedSeed:
		var out mv.Ed25519EncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1EncryptedSecretKey:
		var out mv.Secp256k1EncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256EncryptedSecretKey:
		var out mv.P256EncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381EncryptedSecretKey:
		var out mv.BLSEncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, errors.New("gomav: unknown private key prefix")
	}
}

func ParseSignature(src []byte) (mv.Signature, error) {
	pre, payload, err := base58.DecodeMV(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.GenericSignature:
		var out mv.GenericSignature
		copy(out[:], payload)
		return &out, nil

	case &prefix.Ed25519Signature:
		var out mv.Ed25519Signature
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1Signature:
		var out mv.Secp256k1Signature
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256Signature:
		var out mv.P256Signature
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381Signature:
		var out mv.BLSSignature
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, errors.New("gomav: unknown signature prefix")
	}
}

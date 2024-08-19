package core

//go:generate go run ../../cmd/genmarshaller.go

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/b58/base58"
	"github.com/mavryk-network/gomav/v2/b58/prefix"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core/expression"
)

type OperationContents interface {
	OperationKind() string
}

type OperationContentsAndResult interface {
	OperationContents
	GetMetadata() any
}

type OperationWithResult interface {
	GetResult() ManagerOperationResult
}

type ManagerOperationMetadata interface {
	WithBalanceUpdates
	OperationWithResult
	GetInternalOperationResults() []InternalOperationResult
}

type Bytes struct {
	Bytes []byte `mv:"dyn"`
}

type BallotKind uint8

const (
	BallotYay BallotKind = iota
	BallotNay
	BallotPass
)

func (b BallotKind) String() string {
	switch b {
	case BallotYay:
		return "yay"
	case BallotNay:
		return "nay"
	case BallotPass:
		return "pass"
	default:
		return strconv.FormatInt(int64(b), 10)
	}
}

func (b BallotKind) MarshalText() (text []byte, err error) {
	return []byte(b.String()), nil
}

type ContractID interface {
	mv.Base58Encoder
	TransactionDestination
	ContractID()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ContractID]{
		Variants: encoding.Variants[ContractID]{
			0: ImplicitContract{},
			1: OriginatedContract{},
		},
	})
}

type OriginatedContract struct {
	*mv.ContractHash
	Padding uint8
}

func (OriginatedContract) ContractID()             {}
func (OriginatedContract) OriginatedContractID()   {}
func (OriginatedContract) TransactionDestination() {}
func (a OriginatedContract) Eq(b TransactionDestination) bool {
	if other, ok := b.(OriginatedContract); ok {
		return bytes.Equal(a.ContractHash[:], other.ContractHash[:])
	}
	return false
}

type ImplicitContract struct {
	mv.PublicKeyHash
}

func (ImplicitContract) ContractID()             {}
func (ImplicitContract) TransactionDestination() {}
func (a ImplicitContract) Eq(b TransactionDestination) bool {
	if other, ok := b.(ImplicitContract); ok {
		return a.PublicKeyHash.Eq(other.PublicKeyHash)
	}
	return false
}

type OriginatedContractID interface {
	mv.Base58Encoder
	OriginatedContractID()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OriginatedContractID]{
		Variants: encoding.Variants[OriginatedContractID]{
			1: OriginatedContract{},
		},
	})
}

func ParseContractID(src []byte) (ContractID, error) {
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

	default:
		return nil, errors.New("gomav: unknown contract id prefix")
	}
}

type Entrypoint interface {
	Entrypoint() string
}

type Signed interface {
	GetSignature() (mv.Signature, error)
}

type OperationWithSource interface {
	GetSource() Address
}

type ManagerOperation interface {
	OperationContents
	OperationWithSource
	GetFee() mv.BigUint
	GetCounter() mv.BigUint
	GetGasLimit() mv.BigUint
	GetStorageLimit() mv.BigUint
	SetFee(mv.BigUint)
	SetCounter(mv.BigUint)
	SetGasLimit(mv.BigUint)
	SetStorageLimit(mv.BigUint)
}

type TransactionBase interface {
	OperationContents
	OperationWithSource
	GetAmount() mv.BigUint
	GetDestination() Address
	GetParameters() mv.Option[Parameters]
}

type Transaction interface {
	OperationContents
	ManagerOperation
	TransactionBase
}

type TransactionInternalOperationResult interface {
	InternalOperationResult
	TransactionBase
	GetNonce() uint16
}

type Parameters interface {
	GetEntrypoint() string
	GetValue() expression.Expression
}

type Rat [2]uint16

func (r *Rat) String() string {
	return fmt.Sprintf("%d/%d", r[0], r[1])
}

type BigRat [2]mv.BigInt

func (r *BigRat) String() string {
	return fmt.Sprintf("%v/%v", r[0], r[1])
}

type BlockProtocols struct {
	Protocol     *mv.ProtocolHash
	NextProtocol *mv.ProtocolHash
}

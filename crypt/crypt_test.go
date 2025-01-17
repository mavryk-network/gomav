package crypt

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"sync"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ecadlabs/goblst/minpk"
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/b58"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	title  string
	genKey func() PrivateKey
}

var cases = []testCase{
	{
		title: "Ed25519",
		genKey: func() PrivateKey {
			_, k, _ := ed25519.GenerateKey(rand.Reader)
			return Ed25519PrivateKey(k)
		},
	},
	{
		title: "Secp256k1",
		genKey: func() PrivateKey {
			k, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
			return (*ECDSAPrivateKey)(k)
		},
	},
	{
		title: "P256",
		genKey: func() PrivateKey {
			k, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
			return (*ECDSAPrivateKey)(k)
		},
	},
	{
		title: "BLS",
		genKey: func() PrivateKey {
			k, _ := minpk.GenerateKey(rand.Reader)
			return (*BLSPrivateKey)(k)
		},
	},
}

func TestKey(t *testing.T) {
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// generate key
			priv := c.genKey()
			// encode to internal roundtrip
			mvPriv := priv.ToProtocol()
			tmp, err := NewPrivateKey(mvPriv)
			require.NoError(t, err)
			require.True(t, priv.Equal(tmp))
			require.Equal(t, priv, tmp)

			// encode to base58 roundtrip
			tmp2, err := ParsePrivateKey(priv.ToBase58())
			require.NoError(t, err)
			require.True(t, priv.Equal(tmp2))

			// encode to base58 roundtrip using encrypted type
			tmp3, err := b58.ParsePrivateKey(mvPriv.ToBase58())
			require.NoError(t, err)
			decrypted, err := tmp3.Decrypt(nil)
			require.NoError(t, err)
			require.Equal(t, mvPriv, decrypted)

			// get public
			pub := priv.Public()
			// encode to internal roundtrip
			mvPub := pub.ToProtocol()
			tmp4, err := NewPublicKey(mvPub)
			require.NoError(t, err)
			require.True(t, pub.Equal(tmp4))
			require.Equal(t, pub, tmp4)

			// encode to base58 roundtrip
			tmp5, err := ParsePublicKey(pub.ToBase58())
			require.NoError(t, err)
			require.True(t, pub.Equal(tmp5))
		})
	}
}

func asGeneric(sig mv.Signature) mv.Signature {
	switch sig := sig.(type) {
	case *mv.Ed25519Signature:
		return (*mv.GenericSignature)(sig)
	case *mv.Secp256k1Signature:
		return (*mv.GenericSignature)(sig)
	case *mv.P256Signature:
		return (*mv.GenericSignature)(sig)
	case *mv.BLSSignature:
		return nil
	default:
		panic("unknown")
	}
}

func TestSignature(t *testing.T) {
	var message = []byte("message text")

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			priv := c.genKey()
			sig, err := priv.Sign(message)
			require.NoError(t, err)

			sig1, err := NewSignature(sig.ToProtocol())
			require.NoError(t, err)
			require.Equal(t, sig, sig1)

			sig2, err := ParseSignature(sig.ToBase58())
			require.NoError(t, err)
			require.Equal(t, sig, sig2)

			require.True(t, sig.Verify(priv.Public(), message))

			// via generic
			if genSig := asGeneric(sig.ToProtocol()); genSig != nil {
				sig, err := NewSignature(genSig)
				require.NoError(t, err)
				require.True(t, sig.Verify(priv.Public(), message))
			}
		})
	}
}

var wg sync.WaitGroup

func MinpkGenkey(t *testing.T) {
	defer wg.Done()
	k, err := minpk.GenerateKey(rand.Reader)
	require.NoError(t, err)
	require.NotNil(t, k)
}

func TestMinpkGenkey1kSerial(t *testing.T) {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		MinpkGenkey(t)
	}
}

func TestMinpkGenkey1kParallel(t *testing.T) {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go MinpkGenkey(t)
	}
	wg.Wait()
}

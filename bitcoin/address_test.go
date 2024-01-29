package bitcoin

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/decred/dcrd/crypto/blake256"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/assert"
)

func TestPubKeyToAddr(t *testing.T) {
	network := &chaincfg.TestNet3Params
	pubKeyHex := "0357bbb2d4a9cb8a2357633f201b9c518c2795ded682b7913c6beef3fe23bd6d2f"
	publicKey, err := hex.DecodeString(pubKeyHex)
	assert.Nil(t, err)

	p2pkh, err := PubKeyToAddr(publicKey, LEGACY, network)
	assert.Nil(t, err)
	assert.Equal(t, "mouQtmBWDS7JnT65Grj2tPzdSmGKJgRMhE", p2pkh)

	p2wpkh, err := PubKeyToAddr(publicKey, SEGWIT_NATIVE, network)
	assert.Nil(t, err)
	assert.Equal(t, "tb1qtsq9c4fje6qsmheql8gajwtrrdrs38kdzeersc", p2wpkh)

	p2sh, err := PubKeyToAddr(publicKey, SEGWIT_NESTED, network)
	assert.Nil(t, err)
	assert.Equal(t, "2NF33rckfiQTiE5Guk5ufUdwms8PgmtnEdc", p2sh)

	p2tr, err := PubKeyToAddr(publicKey, TAPROOT, network)
	assert.Nil(t, err)
	assert.Equal(t, "tb1pklh8lqax5l7m2ycypptv2emc4gata2dy28svnwcp9u32wlkenvsspcvhsr", p2tr)
}

// This example demonstrates signing a message with the EC-Schnorr-DCRv0 scheme
// using a secp256k1 private key that is first parsed from raw bytes and
// serializing the generated signature.
func TestExampleSign(t *testing.T) {
	// Decode a hex-encoded private key.
	pkBytes, err := hex.DecodeString("22a47fa09a223f2aa079edf85a7c2d4f8720ee6" +
		"3e502ee2869afab7de234b80c")
	if err != nil {
		fmt.Println(err)
		return
	}
	privKey := secp256k1.PrivKeyFromBytes(pkBytes)

	// Sign a message using the private key.
	message := "test message"
	messageHash := blake256.Sum256([]byte(message))
	signature, err := schnorr.Sign(privKey, messageHash[:])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Serialize and display the signature.
	fmt.Printf("Serialized Signature: %x\n", signature.Serialize())

	// Verify the signature for the message using the public key.
	pubKey := privKey.PubKey()
	verified := signature.Verify(messageHash[:], pubKey)
	fmt.Printf("Signature Verified? %v\n", verified)

	// Output:
	// Serialized Signature: 970603d8ccd2475b1ff66cfb3ce7e622c5938348304c5a7bc2e6015fb98e3b457d4e912fcca6ca87c04390aa5e6e0e613bbbba7ffd6f15bc59f95bbd92ba50f0
	// Signature Verified? true
}

// This example demonstrates verifying an EC-Schnorr-DCRv0 signature against a
// public key that is first parsed from raw bytes.  The signature is also parsed
// from raw bytes.
func TestExampleSignature_Verify(t *testing.T) {
	// Decode hex-encoded serialized public key.
	pubKeyBytes, err := hex.DecodeString("032809768349c8293f896e4f17b82ba4f21edfaa6c28b356f396a6101ad0b9d9a6")
	if err != nil {
		fmt.Println(err)
		return
	}
	pubKey, err := schnorr.ParsePubKey(pubKeyBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decode hex-encoded serialized signature. this sign message is copied from wallet demo
	sigBytes, err := hex.DecodeString("HD9l91NC7J+mhsMHRG6EW29iajNMFmXfb7BaRwxt2K1ITCo1YAi3FhucaF198FGYnLkv5VSCvL3/CizP4uA6Vfc=970603d8ccd2475b1ff66cfb3ce7e622c59383")
	if err != nil {
		fmt.Println(err)
		return
	}
	signature, err := schnorr.ParseSignature(sigBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Verify the signature for the message using the public key.
	message := "test message"
	messageHash := blake256.Sum256([]byte(message))
	verified := signature.Verify(messageHash[:], pubKey)
	fmt.Println("Signature Verified?", verified)

	// Output:
	// Signature Verified? true
}

package client

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/decred/dcrd/crypto/blake256"
	bitcoin "github.com/okx/go-wallet-sdk/coins/bitcoin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListUnspent(t *testing.T) {
	// https://mempool.space/signet/api/address/tb1p8lh4np5824u48ppawq3numsm7rss0de4kkxry0z70dcfwwwn2fcspyyhc7/utxo
	netParams := &chaincfg.SigNetParams
	client := NewClient(netParams)
	address, _ := btcutil.DecodeAddress("tb1p8lh4np5824u48ppawq3numsm7rss0de4kkxry0z70dcfwwwn2fcspyyhc7", netParams)
	unspentList, err := client.ListUnspent(address)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(len(unspentList))
		for _, output := range unspentList {
			t.Log(output.Outpoint.Hash.String(), "    ", output.Outpoint.Index)
		}
	}
}

// get address from pk
func TestPubKeyToAddr(t *testing.T) {
	network := &chaincfg.TestNet3Params
	pubKeyHex := "0357bbb2d4a9cb8a2357633f201b9c518c2795ded682b7913c6beef3fe23bd6d2f"
	publicKey, err := hex.DecodeString(pubKeyHex)
	assert.Nil(t, err)

	p2pkh, err := bitcoin.PubKeyToAddr(publicKey, bitcoin.LEGACY, network)
	assert.Nil(t, err)
	assert.Equal(t, "mouQtmBWDS7JnT65Grj2tPzdSmGKJgRMhE", p2pkh)

	p2wpkh, err := bitcoin.PubKeyToAddr(publicKey, bitcoin.SEGWIT_NATIVE, network)
	assert.Nil(t, err)
	assert.Equal(t, "tb1qtsq9c4fje6qsmheql8gajwtrrdrs38kdzeersc", p2wpkh)

	p2sh, err := bitcoin.PubKeyToAddr(publicKey, bitcoin.SEGWIT_NESTED, network)
	assert.Nil(t, err)
	assert.Equal(t, "2NF33rckfiQTiE5Guk5ufUdwms8PgmtnEdc", p2sh)

	p2tr, err := bitcoin.PubKeyToAddr(publicKey, bitcoin.TAPROOT, network)
	assert.Nil(t, err)
	assert.Equal(t, "tb1pklh8lqax5l7m2ycypptv2emc4gata2dy28svnwcp9u32wlkenvsspcvhsr", p2tr)
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

package client

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/schnorr"
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
	pubKeyBytes, err := hex.DecodeString("03e89b066ce5d807289c22b511341bbe5ed70c7f236ae9e2d05ba740aec0af2ae3")
	if err != nil {
		fmt.Println(err)
		return
	}

	internalKey, err := schnorr.ParsePubKey(pubKeyBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := "hello world~"

	msgHash := fmt.Sprintf(
		"%x",
		sha256.Sum256([]byte(message)),
	)

	hash, hashDecodeError := hex.DecodeString(msgHash)
	if hashDecodeError != nil {
		fmt.Println(err)
		return
	}

	signatureBytes, err := base64.StdEncoding.DecodeString("GxrVlzUTwz3LPaFkfPtbzFKh0fchOGac3RA3PDrFtOfcaQ9gn7r3/wxfWe4xCzX+0ZCBhfYetWBuSads43E52fA=")

	fmt.Println(signatureBytes) // Output: 48656c6c6f
	parseSignature, err := ecdsa.ParseSignature([]byte("304402201ad5973513c33dcb3da1647cfb5bcc52a1d1f72138669cdd10373c3ac5b4e7dc0220690f609fbaf7ff0c5f59ee310b35fed1908185f61eb5606e49a76ce37139d9f0"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Verify the signature for the message using the public key.
	verified := parseSignature.Verify(hash[:], internalKey)
	fmt.Println("Signature Verified?", verified)
	assert.True(t, verified)
	// Output:
	// Signature Verified? true
}

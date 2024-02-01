package main

import (
	"encoding/hex"
	"fmt"
	"github.com/CubicGames/cubic-btc-wallet-server/brc/btcapi/mempool"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	bitcoin "github.com/okx/go-wallet-sdk/coins/bitcoin"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

///01/31 14:16:59 commitTxHash, 2efdd963a14f7e77b1b80eeeb426a8b634e6eb4ce3438c00afb31382475ac850
//2024/01/31 14:16:59 revealTxHash, cf1055732f28e309542b1e596693c2f7e17e8ebb095d72c5d0653d23090d6987
//2024/01/31 14:16:59 inscription, cf1055732f28e309542b1e596693c2f7e17e8ebb095d72c5d0653d23090d6987i0
//2024/01/31 14:16:59 fees:  919

//2024/01/31 14:34:34 commitTxHash, b024a0aac91142b3e97acedfad1e9b6c17f1f9f6c559f3285305d3601f3c145c
//2024/01/31 14:34:34 revealTxHash, d8906bbe09fd064ec5f7e62bf15f26d7ce79fb892bb1a39ad19c4f349da9457e
//2024/01/31 14:34:34 inscription, d8906bbe09fd064ec5f7e62bf15f26d7ce79fb892bb1a39ad19c4f349da9457ei0
//1411

//2024/01/31 14:36:34 file contentType image/png
//2024/01/31 14:36:37 commitTxHash, 267721bf9e918a659977a7cb67daaf102951ba33c3f7bb4a5ab4508d8052fa28
//2024/01/31 14:36:37 revealTxHash, e4cfab0933a8bbdc5b527fed286ed39bd67ca663cf4d0f9dfacdc2629ebb7f5b
//2024/01/31 14:36:37 inscription, e4cfab0933a8bbdc5b527fed286ed39bd67ca663cf4d0f9dfacdc2629ebb7f5bi0
//2024/01/31 14:36:37 fees:  1001

//2024/01/31 14:42:30 commitTxHash, c2a93d9b2773d0300d7231c53ea1a56b29ffb8412befdd3f2972edce7dde3d71
//2024/01/31 14:42:30 revealTxHash, dafdf9b12d76fd9eb1502d91730892311a5c48bd15eed267137943e9aeafc5ee
//2024/01/31 14:42:30 inscription, dafdf9b12d76fd9eb1502d91730892311a5c48bd15eed267137943e9aeafc5eei0
//2024/01/31 14:42:30 fees:  1001

//pk: 58c467cf4f03ad3c73e582e829b04a71fd4bb80517a9d4806f7829b93da1e308 - addr : tb1p8fekvp3f5s92rjl539nmn7wkl8xa4xh0h4n5lh9r8gdnqeyjn00q5v883t
//2024/01/31 21:48:21 commitTxHash, dd37659431012d3acd124ccd44c47fc6ce442df11ef3a466517f08f88a9c063c
//2024/01/31 21:48:21 revealTxHash, d3bc9472970224f614192773dac9bed13b44690e66a123b812cc17ccdda5799d
//2024/01/31 21:48:21 inscription, d3bc9472970224f614192773dac9bed13b44690e66a123b812cc17ccdda5799di0
//2024/01/31 21:48:21 fees:  919

//2024/01/31 21:59:14 file contentType image/png
//2024/01/31 21:59:15 commitTxHash, 4fccebe2f70f43ff1a4ec92a938eab6f85e64b4e1997523ca57a738fdc7eb113
//2024/01/31 21:59:15 revealTxHash, 266b6e3a1acd5be7c59cbf5b12c735e03b7f57346ee8045357f9a33a175b7445
//2024/01/31 21:59:15 inscription, 266b6e3a1acd5be7c59cbf5b12c735e03b7f57346ee8045357f9a33a175b7445i0
//2024/01/31 21:59:15 fees:  919

//2024/01/31 22:01:32 file contentType image/png
//2024/01/31 22:01:33 commitTxHash, 438f91f0f6f4326fe642351b9c84000f75536a7feb941b5a52ffbef1a5df0498
//2024/01/31 22:01:33 revealTxHash, 6b3faaefddd1404ab9213d7ca14c8010b983b425d115be5d54f5aa4b0bc1cb9e
//2024/01/31 22:01:33 inscription, 6b3faaefddd1404ab9213d7ca14c8010b983b425d115be5d54f5aa4b0bc1cb9ei0
//2024/01/31 22:01:33 fees:  919

// image04
//2024/01/31 22:03:00 commitTxHash, 2da575c0ed139fd156b7a743e15ce4fe5578f6c0129ef8bb8655f2a0275e99de
//2024/01/31 22:03:00 revealTxHash, ab413ad087673bbb43086cb84f0da573c85b9bf54c438a49f71227b008bc6e8f
//2024/01/31 22:03:00 inscription, ab413ad087673bbb43086cb84f0da573c85b9bf54c438a49f71227b008bc6e8fi0
//2024/01/31 22:03:00 fees:  849

func TestBtcTx1(t *testing.T) {
	netParams := &chaincfg.TestNet3Params
	btcApiClient := mempool.NewClient(netParams)

	utxoPrivateKeyHex := "58c467cf4f03ad3c73e582e829b04a71fd4bb80517a9d4806f7829b93da1e308"
	//toAddress := "tb1pgd2asxalejy3muv0uw8eeryaejmlj5tpwydpevtsuqcfnumk737qrv7pqu"
	//tb1q078xsrxh5m7kkh0u2urkegja7kve9r5qm9vegk
	//var outpont *wire.OutPoint
	{
		utxoPrivateKeyBytes, err := hex.DecodeString(utxoPrivateKeyHex)
		if err != nil {
			log.Fatal(err)
		}
		utxoPrivateKey, _ := btcec.PrivKeyFromBytes(utxoPrivateKeyBytes)

		utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(utxoPrivateKey.PubKey())), netParams)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf(utxoTaprootAddress.String())

		unspentList, err := btcApiClient.ListUnspent(utxoTaprootAddress)

		if err != nil {
			log.Fatalf("list unspent err %v", err)
		}

		// get one inscription
		for i := range unspentList {
			hash := unspentList[i].Outpoint.Hash.String()
			index := unspentList[i].Outpoint.Index
			log.Printf("hash=%v, index=%v", hash, index)

			value := unspentList[i].Output.Value
			script := unspentList[i].Output.PkScript
			log.Printf("value=%v, script=%v", value, script)
		}
	}
}
func TestBtcTx(t *testing.T) {
	utxoPrivateKeyHex := "cQZFf1boNVxjpRvpPNxELsUK1kh8oZt4Ax6JztPPy6ejhd7uoFh3"
	destination := "tb1q078xsrxh5m7kkh0u2urkegja7kve9r5qm9vegk"
	sender := "tb1p8fekvp3f5s92rjl539nmn7wkl8xa4xh0h4n5lh9r8gdnqeyjn00q5v883t"
	wif, err := btcutil.DecodeWIF(utxoPrivateKeyHex)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//txBuild := bitcoin.NewTxBuild(1, &chaincfg.TestNet3Params)
	//txBuild.AddInput("dafdf9b12d76fd9eb1502d91730892311a5c48bd15eed267137943e9aeafc5ee", 0, utxoPrivateKeyHex, "", "", 0) // replace to your private key
	//txBuild.AddOutput(toAddress, 0)
	//txHex, err := txBuild.SingleBuild()
	//require.Nil(t, err)
	//t.Log(txHex)

	netParams := &chaincfg.TestNet3Params
	btcApiClient := mempool.NewClient(netParams)
	txBuild := bitcoin.NewTxBuild(1, &chaincfg.TestNet3Params)
	// taproot_keypath address
	txBuild = bitcoin.NewTxBuild(1, &chaincfg.TestNet3Params)
	txBuild.AddInput2("fbc3ce8ef5b86f8fd05c9c3c137767f05818322bfa0b6f882e353835a4b510f2", 0, wif.String(), sender, 500)
	txBuild.AddInput2("9c9aa87715f33ff83cff861d8bffbb20bfc06c38dd0cbc9d4dee929dfc91c14c", 0, wif.String(), sender, 98199)
	txBuild.AddOutput(destination, 500)
	txBuild.AddOutput(sender, 97998)
	tx, err := txBuild.Build()
	assert.Nil(t, err)
	txHex, err := bitcoin.GetTxHex(tx)
	assert.Nil(t, err)
	log.Printf(txHex)
	transaction, err := btcApiClient.SendRawTransaction(txHex)
	assert.Nil(t, err)
	log.Printf(transaction.String())

}

func TestParserScript(t *testing.T) {
	scriptHex := "7b2270223a226272632d3230222c226f70223a227472616e73666572222c227469636b223a226f726469222c22616d74223a2237227d"
	bytes, err := hex.DecodeString(scriptHex)
	if err != nil {
		// handle error
	}
	utf8Str := string(bytes)
	fmt.Println(utf8Str)
}

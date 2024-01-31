package main

import (
	"encoding/hex"
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

func main() {
	netParams := &chaincfg.TestNet3Params
	btcApiClient := mempool.NewClient(netParams)

	utxoPrivateKeyHex := "8f89a8d5e05f117af4f91300ce643eb174ef763263f16939771a62b698035499"
	//toAddress := "tb1pgd2asxalejy3muv0uw8eeryaejmlj5tpwydpevtsuqcfnumk737qrv7pqu"

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
			log.Printf(hash)
			if hash == "dafdf9b12d76fd9eb1502d91730892311a5c48bd15eed267137943e9aeafc5ee" {
				//outpont = unspentList[i].Outpoint
			}
		}
	}

}
func TestBtcTx(t *testing.T) {
	utxoPrivateKeyHex := "cSPij8GwFNvKfj4J9Zf9ixDa67ZzWfo9cTypyDMPEKjuTaC1yLJX"
	destination := "tb1pfkd72zchxehrnd3jnxsq80fuqjjqfhh3pfgeh4zchfdtj956dz8qfrs9af"
	//wif, err := btcutil.DecodeWIF(utxoPrivateKeyHex)
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
	txBuild.AddInput2("dafdf9b12d76fd9eb1502d91730892311a5c48bd15eed267137943e9aeafc5ee", 0, utxoPrivateKeyHex, destination, 546)
	txBuild.AddOutput(destination, 300)
	txBuild.AddOutput(destination, 300)

	tx, err := txBuild.Build()
	assert.Nil(t, err)
	txHex, err := bitcoin.GetTxHex(tx)
	t.Log(txHex)
	_, err = btcApiClient.SendRawTransaction(txHex)
	if err != nil {
		log.Fatal(err)
		return
	}

}

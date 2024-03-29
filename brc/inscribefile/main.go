package main

import (
	"encoding/hex"
	"fmt"
	"github.com/CubicGames/cubic-btc-wallet-server/brc/btcapi/client"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"log"
	"net/http"
	"os"
)

func main() {
	netParams := &chaincfg.TestNet3Params
	btcApiClient := client.NewClient(netParams)

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory, %v", err)
	}
	filePath := fmt.Sprintf("%s/images/0004.png", workingDir)
	// if file size too max will return sendrawtransaction RPC error: {"code":-26,"message":"tx-size"}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file %v", err)
	}

	contentType := http.DetectContentType(fileContent)
	log.Printf("file contentType %s", contentType)

	//tb1pflq6z6mdduna235j3k3wn8tu6r39d4lc5celw9c7tfu6agp2yxvqfpyzqh
	//utxoPrivateKeyHex := "8f89a8d5e05f117af4f91300ce643eb174ef763263f16939771a62b698035499"
	utxoPrivateKeyHex := "58c467cf4f03ad3c73e582e829b04a71fd4bb80517a9d4806f7829b93da1e308"
	destination := "tb1p8fekvp3f5s92rjl539nmn7wkl8xa4xh0h4n5lh9r8gdnqeyjn00q5v883t"

	commitTxOutPointList := make([]*wire.OutPoint, 0)
	commitTxPrivateKeyList := make([]*btcec.PrivateKey, 0)
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

		unspentList, err := btcApiClient.ListUnspent(utxoTaprootAddress)

		if err != nil {
			log.Fatalf("list unspent err %v", err)
		}

		for i := range unspentList {
			if "438f91f0f6f4326fe642351b9c84000f75536a7feb941b5a52ffbef1a5df0498" == unspentList[i].Outpoint.Hash.String() {
				commitTxOutPointList = append(commitTxOutPointList, unspentList[i].Outpoint)
				commitTxPrivateKeyList = append(commitTxPrivateKeyList, utxoPrivateKey)
			}
		}
	}

	request := InscriptionRequest{
		CommitTxOutPointList:   commitTxOutPointList,
		CommitTxPrivateKeyList: commitTxPrivateKeyList,
		CommitFeeRate:          2,
		FeeRate:                1,
		DataList: []InscriptionData{
			{
				ContentType: contentType,
				Body:        fileContent,
				Destination: destination,
			},
		},
		SingleRevealTxOnly: true,
	}

	tool, err := NewInscriptionToolWithBtcApiClient(netParams, btcApiClient, &request)
	if err != nil {
		log.Fatalf("Failed to create inscription tool: %v", err)
	}
	commitTxHash, revealTxHashList, inscriptions, fees, err := tool.Inscribe()
	if err != nil {
		log.Fatalf("send tx errr, %v", err)
	}
	log.Println("commitTxHash, " + commitTxHash.String())
	for i := range revealTxHashList {
		log.Println("revealTxHash, " + revealTxHashList[i].String())
	}
	for i := range inscriptions {
		log.Println("inscription, " + inscriptions[i])
	}
	log.Println("fees: ", fees)
}

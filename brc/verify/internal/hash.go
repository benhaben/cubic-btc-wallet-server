package internal

import (
	"github.com/CubicGames/cubic-btc-wallet-server/brc/verify/internal/flags"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/samber/lo"
)

// GeneratePublicKeyHash returns the public key hash, either compressed or uncompressed, depending on the recovery flag.
func GeneratePublicKeyHash(recoveryFlag int, publicKey *btcec.PublicKey) []byte {
	if lo.Contains[int](flags.Uncompressed(), recoveryFlag) {
		return btcutil.Hash160(publicKey.SerializeUncompressed())
	}

	return btcutil.Hash160(publicKey.SerializeCompressed())
}

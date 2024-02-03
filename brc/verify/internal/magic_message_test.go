package internal_test

import (
	"github.com/CubicGames/cubic-btc-wallet-server/brc/verify/internal"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateMagicMessage(t *testing.T) {
	t.Parallel()

	message := internal.CreateMagicMessage("random message")
	require.Equal(t, "\x18Bitcoin Signed Message:\n\x0Erandom message", message)
}

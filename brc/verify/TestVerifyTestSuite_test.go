package verifier

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

// This example will output `true, <nil>` for both signed messages, since the signature is valid and there are no errors.
func TestVerifyTestSuite(t *testing.T) {
	// Bitcoin Mainnet
	fmt.Println(VerifyWithChain(SignedMessage{
		Address:   "18J72YSM9pKLvyXX1XAjFXA98zeEvxBYmw",
		Message:   "Test123",
		Signature: "Gzhfsw0ItSrrTCChykFhPujeTyAcvVxiXwywxpHmkwFiKuUR2ETbaoFcocmcSshrtdIjfm8oXlJoTOLosZp3Yc8=",
	}, &chaincfg.MainNetParams))

	// Bitcoin Testnet3
	chain, err := VerifyWithChain(SignedMessage{
		Address:   "tb1pv5yvehnrkksmn6rn309ygcqvj4vdc4jfdvs4gg26wxf5qpmaku5qznl46t",
		Message:   "hello world~",
		Signature: "G920SOk4jEhTz8+UqVxjSw78pRav7YKzyg4DaNJqzp7ealVbZtxufwEOMECzOEzi9AFOiHWF7Pq6vZzclKt0c5M=",
	}, &chaincfg.TestNet3Params)
	assert.Nil(t, err)
	assert.True(t, chain)
}

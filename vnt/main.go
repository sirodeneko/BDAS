package vnt

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/vntchain/go-vnt/common"
	"github.com/vntchain/go-vnt/core/types"
	"github.com/vntchain/go-vnt/crypto"
	"math/big"
	"os"
)

var privateKeyStr string
var toAddressStr string
var formAddressStr string
var privateKey *ecdsa.PrivateKey
var toAddress common.Address

func VntInit() {
	vntRequestInit()
	privateKeyStr = os.Getenv("PRIVSTE_KEY")
	toAddressStr = os.Getenv("TO_ADDRESS")
	formAddressStr = os.Getenv("FORM_ADDRESS")
	privateKey, _ = crypto.HexToECDSA(privateKeyStr)
	toAddress = common.HexToAddress(toAddressStr)
}

func sign(data []byte) (string, error) {
	value := big.NewInt(0)     // in wei (1 eth)
	gasLimit := uint64(500000) // in units
	gasPrice := big.NewInt(500000000000)

	tx := types.NewTransaction(10, toAddress, value, gasLimit, gasPrice, data)
	signer := types.NewHubbleSigner(big.NewInt(2))

	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return "", err
	}
	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)

	return "0x" + rawTxHex, err
}

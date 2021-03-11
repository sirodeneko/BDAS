package vnt

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/vntchain/go-vnt/common"
	"github.com/vntchain/go-vnt/core/types"
	"github.com/vntchain/go-vnt/crypto"
	"math/big"
	"os"
	"strconv"
)

var privateKeyStr string
var ToAddressStr string
var FormAddressStr string
var privateKey *ecdsa.PrivateKey
var toAddress common.Address

func VntInit() {
	vntRequestInit()
	privateKeyStr = os.Getenv("PRIVSTE_KEY")
	ToAddressStr = os.Getenv("TO_ADDRESS")
	FormAddressStr = os.Getenv("FORM_ADDRESS")
	privateKey, _ = crypto.HexToECDSA(privateKeyStr)
	toAddress = common.HexToAddress(ToAddressStr)
}

func Sign(data []byte, nonceHex string) string {
	value := big.NewInt(0)     // in wei (1 eth)
	gasLimit := uint64(500000) // in units
	gasPrice := big.NewInt(500000000000)

	val := nonceHex[2:]
	nonce, _ := strconv.ParseUint(val, 16, 32)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signer := types.NewHubbleSigner(big.NewInt(2))

	signedTx, _ := types.SignTx(tx, signer, privateKey)

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)

	return "0x" + rawTxHex
}

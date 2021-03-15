package vnt

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/vntchain/go-vnt/common"
	"github.com/vntchain/go-vnt/core/types"
	"github.com/vntchain/go-vnt/crypto"
	"math/big"
	"os"
	"singo/util"
	"strconv"
	"sync/atomic"
)

var privateKeyStr string
var ToAddressStr string
var FormAddressStr string
var privateKey *ecdsa.PrivateKey
var toAddress common.Address

// nonce 分发器变量
var nonce *uint64

func VntInit() {
	vntRequestInit()
	privateKeyStr = os.Getenv("PRIVSTE_KEY")
	ToAddressStr = os.Getenv("TO_ADDRESS")
	FormAddressStr = os.Getenv("FORM_ADDRESS")
	privateKey, _ = crypto.HexToECDSA(privateKeyStr)
	toAddress = common.HexToAddress(ToAddressStr)
	nonceInit()
}

func Sign(data []byte, nonce uint64) string {
	value := big.NewInt(0)     // in wei (1 eth)
	gasLimit := uint64(500000) // in units
	gasPrice := big.NewInt(500000000000)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signer := types.NewHubbleSigner(big.NewInt(2))

	signedTx, _ := types.SignTx(tx, signer, privateKey)

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)

	return "0x" + rawTxHex
}

func nonceInit() {
	nonce = new(uint64)
	// 获取nonce
	// 注意，必须先确保FormAddressStr被赋值
	nonceHex, err := GetTransactionCount(FormAddressStr)
	if err != nil {
		util.Log().Panic("nonceHex获取失败，区块链初始化失败: %v", err)
		return
	}
	val := nonceHex[2:]
	num, _ := strconv.ParseUint(val, 16, 32)
	atomic.StoreUint64(nonce, num)
}

//-------------------------------------------
//TODO 此nonce为单机分发器，多机获取nonce必须重构

func GetNonce() uint64 {
	// 加一后返回
	num := atomic.AddUint64(nonce, 1)
	// 因为要加一又要保证原子性，故必须只执行一次，对其+1，返回加一后的值，需要的是原本的值，故减一
	return num - 1
}

func ReNonceInitWithGet() uint64 {
	// 获取nonce
	// 注意，必须先确保FormAddressStr被赋值
	nonceHex, err := GetTransactionCount(FormAddressStr)
	if err != nil {
		util.Log().Panic("nonceHex获取失败，区块链初始化失败: %v", err)
		return 0
	}
	val := nonceHex[2:]
	num, _ := strconv.ParseUint(val, 16, 32)
	atomic.StoreUint64(nonce, num)
	return GetNonce()
}

/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:main.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package vnt

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
	"os"
	"singo/util"
)

var privateKeyStr string
var ToAddressStr string
var FormAddressStr string
var privateKey *ecdsa.PrivateKey
var toAddress common.Address
var fromAddress common.Address
var url string
var client *ethclient.Client

func VntInit() {
	url = os.Getenv("GNT_RPC_URL")
	privateKeyStr = os.Getenv("PRIVSTE_KEY")
	ToAddressStr = os.Getenv("TO_ADDRESS")
	FormAddressStr = os.Getenv("FORM_ADDRESS")
	privateKey, _ = crypto.HexToECDSA(privateKeyStr)
	fromAddress = common.HexToAddress(FormAddressStr)
	toAddress = common.HexToAddress(ToAddressStr)
	c, err := ethclient.Dial(url)
	client = c
	if err != nil {
		util.Log().Panic("以太坊连接失败", err)
	}
}

func SendTransaction(data []byte) (string, error) {

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", errors.Errorf("获取nonce失败:%v", err)
	}

	value := big.NewInt(0)     // in wei (1 eth)
	gasLimit := uint64(210000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", errors.Errorf("获取gas价格失败:%v", err)
	}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", errors.Errorf("获取chainID失败:%v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", errors.Errorf("签名失败:%v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", errors.Errorf("交易发送失败:%v", err)
	}

	return signedTx.Hash().Hex(), nil //0x...
}

// https://faucet.rinkeby.io/ 水龙头
// https://rinkeby.etherscan.io/ 区块链浏览器
// https://rinkeby.infura.io/v3/144451a1ec8e493891a105db4147309b rpc网址
// https://twitter.com/siro59344443/status/1380052088139382785

/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:request.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package vnt

import (
	"encoding/json"
	"errors"
	"github.com/asmcos/requests"
	"os"
)

var url string
var jsonrpc = "2.0"

type RpcPost struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type RpcResult struct {
	Jsonrpc string                 `json:"jsonrpc"`
	Id      int                    `json:"id"`
	Result  interface{}            `json:"result"`
	Err     map[string]interface{} `json:"error"`
}

func vntRequestInit() {
	url = os.Getenv("VNT_RPC_URL")
}

// GetTransactionCount 获取当前交易数量，返回十六进制数量
func GetTransactionCount(address string) (string, error) {
	rpcPost := RpcPost{
		Jsonrpc: jsonrpc,
		Id:      1,
		Method:  "core_getTransactionCount",
		Params:  []interface{}{address, "latest"},
	}
	data, err := json.Marshal(rpcPost)
	if err != nil {
		return "", err
	}
	data, err = createRequest(data)
	if err != nil {
		return "", err
	}
	var rpcResult RpcResult
	err = json.Unmarshal(data, &rpcResult)
	if err != nil {
		return "", err
	}
	if rpcResult.Err != nil || len(rpcResult.Err) != 0 {
		return "", errors.New(rpcResult.Err["message"].(string))
	}

	return rpcResult.Result.(string), nil

}

// GetestimateGas 获取单价
func GetestimateGas(fromAddress string, toAddress string, nonce string, value string, input string) ([]byte, error) {
	rpcPost := RpcPost{
		Jsonrpc: jsonrpc,
		Id:      1,
		Method:  "core_estimateGas",
		Params:  []interface{}{},
	}
	tx := make(map[string]interface{}, 6)

	tx["from"] = fromAddress
	tx["chainId"] = 2
	tx["nonce"] = nonce
	tx["to"] = toAddress
	tx["value"] = value
	tx["data"] = input
	rpcPost.Params = append(rpcPost.Params, tx)

	data, err := json.Marshal(rpcPost)
	if err != nil {
		return nil, err
	}
	return createRequest(data)
}

// SendRawTransaction 广播交易，返回交易地址
func SendRawTransaction(signstr string) (string, error) {
	rpcPost := RpcPost{
		Jsonrpc: jsonrpc,
		Id:      1,
		Method:  "core_sendRawTransaction",
		Params:  []interface{}{},
	}
	rpcPost.Params = append(rpcPost.Params, signstr)
	data, err := json.Marshal(rpcPost)
	if err != nil {
		return "", err
	}
	data, err = createRequest(data)
	if err != nil {
		return "", err
	}

	var rpcResult RpcResult
	err = json.Unmarshal(data, &rpcResult)
	if err != nil {
		return "", err
	}
	if rpcResult.Err != nil || len(rpcResult.Err) != 0 {
		return "", errors.New(rpcResult.Err["message"].(string))
	}
	return rpcResult.Result.(string), nil

}

// GetTransactionByHash 查询交易信息，返回存入的data(十六进制)
func GetTransactionByHash(txHash string) (string, error) {
	rpcPost := RpcPost{
		Jsonrpc: jsonrpc,
		Id:      1,
		Method:  "core_getTransactionByHash",
		Params:  []interface{}{},
	}
	rpcPost.Params = append(rpcPost.Params, txHash)
	data, err := json.Marshal(rpcPost)
	if err != nil {
		return "", err
	}
	data, err = createRequest(data)
	if err != nil {
		return "", err
	}

	var rpcResult RpcResult
	err = json.Unmarshal(data, &rpcResult)
	if err != nil {
		return "", err
	}
	if rpcResult.Err != nil || len(rpcResult.Err) != 0 {
		return "", errors.New(rpcResult.Err["message"].(string))
	}

	return rpcResult.Result.(map[string]interface{})["input"].(string), nil
}
func createRequest(data []byte) ([]byte, error) {
	req := requests.Requests()
	req.Header.Set("Content-Type", "application/json")
	resp, err := req.PostJson(url, string(data))
	if err != nil {
		return nil, err
	}
	redata := resp.Content()
	return redata, err
}

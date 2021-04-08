/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:request.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package vnt

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

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

//// SendRawTransaction 广播交易，返回交易地址
//func SendRawTransaction(signstr string) (string, error) {
//	rpcPost := RpcPost{
//		Jsonrpc: jsonrpc,
//		Id:      1,
//		Method:  "core_sendRawTransaction",
//		Params:  []interface{}{},
//	}
//	rpcPost.Params = append(rpcPost.Params, signstr)
//	data, err := json.Marshal(rpcPost)
//	if err != nil {
//		return "", err
//	}
//	data, err = createRequest(data)
//	if err != nil {
//		return "", err
//	}
//
//	var rpcResult RpcResult
//	err = json.Unmarshal(data, &rpcResult)
//	if err != nil {
//		return "", err
//	}
//	if rpcResult.Err != nil || len(rpcResult.Err) != 0 {
//		return "", errors.New(rpcResult.Err["message"].(string))
//	}
//	return rpcResult.Result.(string), nil
//
//}

// GetTransactionByHash 查询交易信息，返回存入的data(十六进制)
func GetTransactionByHash(txHash string) (string, error) {

	hash := common.HexToHash(txHash)
	tx, _, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return "", errors.Errorf("获取交易失败:%v", err)
	}
	return string(tx.Data()), nil

}

//func createRequest(data []byte) ([]byte, error) {
//	req := requests.Requests()
//	req.Header.Set("Content-Type", "application/json")
//	resp, err := req.PostJson(url, string(data))
//	if err != nil {
//		return nil, err
//	}
//	redata := resp.Content()
//	return redata, err
//}

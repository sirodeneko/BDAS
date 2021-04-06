/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:certificate_with_file_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"encoding/hex"
	"encoding/json"
	"singo/model"
	"singo/serializer"
	"singo/util"
	"singo/vnt"
)

type CertificateWithFileService struct {
	FileUrl string `json:"file_url" form:"file_url" binding:"required"`
	FileID  uint   `json:"file_id" form:"file_id" binding:"required"`
}

func (service *CertificateWithFileService) CertificateWithFile() serializer.Response {
	var certification model.Certification
	err := model.DB.Find(&certification, service.FileID).Error
	if err != nil {
		return serializer.DBErr("证书查询失败", err)
	}
	fileHash := util.FileSHA256(util.FileURL + service.FileUrl)

	dataHex, err := vnt.GetTransactionByHash(certification.Address)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "区块链错误",
			Error: err.Error(),
		}
	}
	data, _ := hex.DecodeString(dataHex[2:])
	pictureInfoJson := util.AesDecrypt(string(data), util.GetJsonKey())
	var pictureInfo util.PictureInfo
	err = json.Unmarshal([]byte(pictureInfoJson), &pictureInfo)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "链上数据转化失败",
			Error: err.Error(),
		}
	}
	if pictureInfo.FileHash == fileHash {
		return serializer.Response{
			Data: true,
		}
	} else {
		return serializer.Response{
			Data: false,
		}
	}
}

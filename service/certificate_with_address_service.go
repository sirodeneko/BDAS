/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:certificate_with_address_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"encoding/json"
	"singo/model"
	"singo/serializer"
	"singo/util"
	"singo/vnt"
)

type CertificateWithAddressService struct {
	Address  string `json:"address" form:"address" binding:"required"`
	CardCode string `json:"card_code" form:"card_code" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
}

func (service *CertificateWithAddressService) CertificateWithAddress() serializer.Response {
	var certification model.Certification
	err := model.DB.Where("address = ?", service.Address).First(&certification).Error
	if err != nil {
		return serializer.DBErr("数据库查询失败", err)
	}
	if certification.CardCode != service.CardCode && certification.Name != service.Name {
		return serializer.Response{
			Code: 403,
			Msg:  "信息错误或者证书不存在",
		}
	}

	dataHex, err := vnt.GetTransactionByHash(certification.Address)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "区块链错误",
			Error: err.Error(),
		}
	}
	pictureInfoJson := util.AesDecrypt(dataHex, util.GetJsonKey())
	var pictureInfo util.PictureInfo
	err = json.Unmarshal([]byte(pictureInfoJson), &pictureInfo)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "链上数据转化失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: pictureInfo.Response(),
	}
}

/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:certification_info_service.go
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

type CertificationInfoService struct {
	ID string `json:"id" form:"id" binding:"required"`
}

func (service *CertificationInfoService) CertificationInfo(user *model.User) serializer.Response {
	var certification model.Certification
	err := model.DB.First(&certification, service.ID).Error
	if err != nil {
		return serializer.DBErr("数据库查询失败", err)
	}

	if user.CardCode != certification.CardCode {
		return serializer.Response{
			Code: 403,
			Msg:  "禁止访问",
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
	// 返回的dataHex，其实就是正常的utf-8不需要转换
	//data, _ := hex.DecodeString(dataHex)
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
		Data: pictureInfo.FileUrl,
	}
}

package service

import (
	"encoding/hex"
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
	return serializer.Response{
		Data: pictureInfo.FileUrl,
	}
}

package service

import (
	"encoding/json"
	"os"
	"singo/model"
	"singo/serializer"
	"singo/util"
	"singo/vnt"
)

type AdminACStudentService struct {
	MsgID string `json:"msg_id" form:"msg_id" binding:"required"`
}

var jsonKey string

func (service *AdminACStudentService) AdminACStudent() serializer.Response {
	var message model.Message

	err := model.DB.Preload("EducationalAcMsg").First(&message, service.MsgID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "消息查询失败",
			Error: err.Error(),
		}
	}
	if message.MsgType != model.EducationalQualifications {
		return serializer.Response{
			Code: 404,
			Msg:  "非法请求",
		}
	}

	pictureInfo := model2PictureInfo(&message.EducationalAcMsg)

	// 写入数据库
	certification := model.Certification{
		CardCode:     pictureInfo.CardCode,
		Level:        pictureInfo.Level,
		Professional: pictureInfo.Professional,
		Name:         pictureInfo.Name,
	}
	err = model.DB.Save(&certification).Error
	if err != nil {
		return serializer.DBErr("信息保存失败", err)
	}

	pictureInfo.FileID = certification.ID

	url, err := pictureInfo.CreateImg()
	if err != nil {
		return serializer.Response{
			Code:  500,
			Data:  nil,
			Msg:   "认证文件生成失败",
			Error: err.Error(),
		}
	}

	pictureInfo.FileUrl = url
	pictureInfo.FileHash = util.FileSHA256(util.SaveURL + url)

	pictureInfoByte, _ := json.Marshal(pictureInfo)
	pictureInfCrypt := util.AesEncrypt(string(pictureInfoByte), getJsonKey())

	// 获取nonce
	nonceHex, err := vnt.GetTransactionCount(vnt.FormAddressStr)
	if err != nil {
		return serializer.Response{
			Code:  5000,
			Data:  nil,
			Msg:   "区块链网络错误",
			Error: err.Error(),
		}
	}
	// 签名
	signHex := vnt.Sign([]byte(pictureInfCrypt), nonceHex)
	// 广播上链
	txAddress, _ := vnt.SendRawTransaction(signHex)

	// 写入数据库
	certification.Address = txAddress
	certification.Url = url
	err = model.DB.Save(certification).Error
	if err != nil {
		return serializer.DBErr("信息保存失败", err)
	}

	// 删除msg
	model.DB.Delete(message.EducationalAcMsg)
	model.DB.Delete(message)

	return serializer.Response{
		Msg: "用户激活成功",
	}
}

func model2PictureInfo(msg *model.EducationalAcMsg) util.PictureInfo {
	return util.PictureInfo{
		CreatedAt:         msg.CreatedAt.Unix(),
		Name:              msg.Name,
		Sex:               msg.Sex,
		Ethnic:            msg.Ethnic,
		Birthday:          msg.Birthday.Unix(),
		CardCode:          msg.CardCode,
		EducationCategory: msg.EducationCategory,
		Level:             msg.Level,
		University:        msg.University,
		Professional:      msg.Professional,
		LearningFormat:    msg.LearningFormat,
		EducationalSystem: msg.EducationalSystem,
		AdmissionDate:     msg.AdmissionDate,
		GraduationDate:    msg.GraduationDate,
		Status:            msg.Status,
		StudentAvatar:     msg.StudentAvatar,
	}
}

func getJsonKey() string {
	if jsonKey == "" {
		jsonKey = os.Getenv("JSON_KEY")
	}
	return jsonKey
}

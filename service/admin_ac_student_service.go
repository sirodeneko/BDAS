package service

import (
	"singo/model"
	"singo/serializer"
	"singo/util"
)

type AdminACStudentService struct {
	MsgID string `json:"msg_id" form:"msg_id" binding:"required"`
}

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

	url, err := pictureInfo.CreateImg()
	if err != nil {
		return serializer.Response{
			Code:  500,
			Data:  nil,
			Msg:   "认证文件生成失败",
			Error: err.Error(),
		}
	}

	url = url
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

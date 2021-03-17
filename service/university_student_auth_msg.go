package service

import (
	"singo/model"
	"singo/serializer"
)

type StudentAuthMsgService struct {
	MsgID string `json:"msg_id" form:"msg_id" binding:"required"`
}

func (service *StudentAuthMsgService) StudentAuthMsg(university *model.University) serializer.Response {

	var message model.Message

	err := model.DB.Preload("EducationalAcMsg").First(&message, service.MsgID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "消息查询失败",
			Error: err.Error(),
		}
	}
	if message.EducationalAcMsg.University != university.UniversityName {
		return serializer.Response{
			Code: 403,
			Msg:  "权限不足，禁止访问",
		}
	}
	return serializer.BuildMessageResponse(message)
}

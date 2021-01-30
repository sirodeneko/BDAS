package service

import (
	"singo/model"
	"singo/serializer"
)

type MsgListService struct {
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
}

func (service *MsgListService) MsgList() serializer.Response {
	var msgs []model.Message
	total := 0

	if service.Limit == 0 {
		service.Limit = 12
	}

	if err := model.DB.Model(model.Message{}).Count(&total).Error; err != nil {
		return serializer.DBErr("数据库链接错误", err)
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Offset).Preload("StudentAcMsg").Preload("EducationalAcMsg").Find(&msgs).Error; err != nil {
		return serializer.DBErr("数据库链接错误", err)
	}

	return serializer.BuildListResponse(serializer.BuildMessages(msgs), uint(total))
}

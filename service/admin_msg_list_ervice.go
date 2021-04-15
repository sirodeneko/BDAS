/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:admin_msg_list_ervice.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"singo/model"
	"singo/serializer"
)

type MsgListService struct {
	Limit  int    `json:"limit" form:"limit"`
	Offset int    `json:"offset" form:"offset"`
	MType  string `json:"m_type" form:"m_type" binding:"required"`
}

func (service *MsgListService) MsgList() serializer.Response {
	var msgs []model.Message
	total := 0
	db := model.DB.Model(model.Message{})
	switch service.MType {
	case model.EducationalQualifications:
		db = db.Where("student_ac_msg_id = ?", 0)
	case model.StudentAccreditation:
		db = db.Where("educational_ac_msg_id = ?", 0)
	default:
		return serializer.Response{Code: 400, Msg: "类型不存在"}
	}

	if service.Limit == 0 {
		service.Limit = 12
	}

	if err := db.Count(&total).Error; err != nil {
		return serializer.DBErr("数据库链接错误", err)
	}

	if err := db.Limit(service.Limit).Offset(service.Offset).Preload("StudentAcMsg").Preload("EducationalAcMsg").Find(&msgs).Error; err != nil {
		return serializer.DBErr("数据库链接错误", err)
	}

	return serializer.BuildListResponse(serializer.BuildMessages(msgs), uint(total))
}

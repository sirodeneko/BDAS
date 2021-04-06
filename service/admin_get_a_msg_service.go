/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:admin_get_a_msg_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"singo/model"
	"singo/serializer"
)

type GetAMsgService struct {
	MType string `form:"m_type" json:"m_type" binding:"required"`
}

func (service *GetAMsgService) GetAMsg() serializer.Response {
	switch service.MType {
	// 学生认证
	case model.StudentAccreditation:
		return getStudentAccreditation()
	// 学历认证
	case model.EducationalQualifications:
		return getEducationalQualifications()
	default:
		return serializer.Response{
			Code:  403,
			Msg:   "类型错误",
			Error: "admin_get_msg:类型错误",
		}
	}

}

func getStudentAccreditation() serializer.Response {
	var message model.Message

	err := model.DB.Where("msg_type = ?", model.StudentAccreditation).Preload("StudentAcMsg").Last(&message).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "消息查询失败或者已全部审核完成",
			Error: err.Error(),
		}
	}

	return serializer.BuildMessageResponse(message)

}

func getEducationalQualifications() serializer.Response {
	var message model.Message

	err := model.DB.Where("msg_type = ?", model.EducationalQualifications).Preload("EducationalAcMsg").Last(&message).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "消息查询失败或者已全部审核完成",
			Error: err.Error(),
		}
	}

	return serializer.BuildMessageResponse(message)
}

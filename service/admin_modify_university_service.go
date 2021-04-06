/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:admin_modify_university_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"singo/model"
	"singo/serializer"
)

// AdminModifyUniversityService 管理员修改学校信息的服务
type AdminModifyUniversityService struct {
	ID             uint   `form:"id" json:"id" binding:"required"`
	Nickname       string `form:"nickname" json:"nickname" binding:"omitempty,min=2,max=30"`
	NewPassword    string `form:"new_password" json:"new_password" binding:"omitempty,min=2,max=30"`
	Status         string `json:"status" form:"status" `
	UniversityName string `json:"university_name" form:"university_name"`
}

// Change 用户修改信息
func (service *AdminModifyUniversityService) AdminModifyUniversity() serializer.Response {
	var university model.University
	//找到用户
	err := model.DB.First(&university, service.ID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询用户失败",
			Error: err.Error(),
		}
	}

	// 如果更新了密码
	if service.NewPassword != "" {
		if err := university.SetPassword(service.NewPassword); err != nil {
			return serializer.Err(
				serializer.CodeEncryptError,
				"密码加密失败",
				err,
			)
		}
	}
	university.Nickname = service.Nickname
	university.Status = service.Status
	university.UniversityName = service.UniversityName

	err = model.DB.Model(&university).Update(&university).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeDBError,
			Msg:   "用户信息保存失败",
			Error: err.Error(),
		}
	}
	model.DB.Find(&university)
	return serializer.Response{
		Data: serializer.BuildUniversity(university),
	}
}

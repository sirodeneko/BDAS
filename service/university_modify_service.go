/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:university_modify_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"singo/model"
	"singo/serializer"
)

// UniversityModifyService 学校用户修改信息的服务
type UniversityModifyService struct {
	Nickname    string `form:"nickname" json:"nickname" binding:"omitempty,min=2,max=30"`
	OldPassword string `form:"old_password" json:"old_password"`
	NewPassword string `form:"new_password" json:"new_password" binding:"omitempty,min=8,max=40"`
}

// Change 用户修改信息
func (service *UniversityModifyService) UniversityModify(ID uint) serializer.Response {
	var university model.University
	//找到用户
	err := model.DB.First(&university, ID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询用户失败",
			Error: err.Error(),
		}
	}

	// 如果更新了密码
	if service.NewPassword != "" {
		if university.CheckPassword(service.OldPassword) == false {
			return serializer.ParamErr("账号或密码错误", nil)
		}

		if err := university.SetPassword(service.NewPassword); err != nil {
			return serializer.Err(
				serializer.CodeEncryptError,
				"密码加密失败",
				err,
			)
		}
	}
	university.Nickname = service.Nickname

	err = model.DB.Model(&university).Update(&university).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeDBError,
			Msg:   "用户信息保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUniversity(university),
	}
}

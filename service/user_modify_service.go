/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:user_modify_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"singo/model"
	"singo/serializer"
)

// UserModifyService 普通用户修改信息的服务
type UserModifyService struct {
	Nickname    string `form:"nickname" json:"nickname" binding:"omitempty,min=2,max=30"`
	OldPassword string `form:"old_password" json:"old_password"`
	NewPassword string `form:"new_password" json:"new_password" binding:"omitempty,min=8,max=40"`
}

// Change 用户修改信息
func (service *UserModifyService) UserModify(ID uint) serializer.Response {
	var user model.User
	//找到用户
	err := model.DB.First(&user, ID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询用户失败",
			Error: err.Error(),
		}
	}

	// 如果更新了密码
	if service.NewPassword != "" {
		if user.CheckPassword(service.OldPassword) == false {
			return serializer.ParamErr("账号或密码错误", nil)
		}

		if err := user.SetPassword(service.NewPassword); err != nil {
			return serializer.Err(
				serializer.CodeEncryptError,
				"密码加密失败",
				err,
			)
		}
	}
	if service.Nickname != "" {
		user.Nickname = service.Nickname
	}

	err = model.DB.Model(&user).Update(&user).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeDBError,
			Msg:   "用户信息保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}

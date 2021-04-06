/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:university_register_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"singo/model"
	"singo/serializer"
)

// UniversityRegisterService 学校用户注册服务
type UniversityRegisterService struct {
	UniversityName  string `form:"university_name" json:"university_name" binding:"required"`
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UniversityRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.University{}).Where("nickname = ? and university_name = ?", service.Nickname, service.UniversityName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}

	count = 0
	model.DB.Model(&model.University{}).Where("user_name = ? and university_name = ?", service.UserName, service.UniversityName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 学校注册
func (service *UniversityRegisterService) Register() serializer.Response {
	university := model.University{
		UniversityName: service.UniversityName,
		Nickname:       service.Nickname,
		UserName:       service.UserName,
		Status:         model.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := university.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := model.DB.Create(&university).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUniversityResponse(university)
}

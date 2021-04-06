/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:user_certification_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"singo/model"
	"singo/serializer"
)

type CertificationService struct {
}

func (service *CertificationService) Certification(user *model.User) serializer.Response {
	var certifications []model.Certification
	if user.CardCode == "" {
		return serializer.Response{
			Data: nil,
		}
	}
	err := model.DB.Where("card_code = ?", user.CardCode).Find(&certifications).Error
	if err != nil {
		return serializer.DBErr("数据库查询失败", err)
	}
	return serializer.BuildListResponse(serializer.BuildCertifications(certifications), uint(len(certifications)))
}

/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:university_student_auth_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
	"time"
)

type StudentAuthService struct {
	Name              string `json:"name" form:"name" binding:"required"`
	Sex               uint   `json:"sex" form:"sex" binding:"required"`                               // 0 男 1女
	Ethnic            string `json:"ethnic" form:"ethnic" binding:"required"`                         // 民族
	Birthday          int64  `json:"birthday" form:"birthday" binding:"required"`                     // 生日
	CardCode          string `json:"card_code" form:"card_code" binding:"required"`                   // 身份证号
	EducationCategory string `json:"education_category" form:"education_category" binding:"required"` // 学历类别
	Level             string `json:"level" form:"level" binding:"required"`                           // 层次
	University        string `json:"university" form:"university" binding:"required"`                 // 学校
	Professional      string `json:"professional" form:"professional" binding:"required"`             // 专业
	LearningFormat    string `json:"learning_format" form:"learning_format" binding:"required"`       // 学习形式
	EducationalSystem string `json:"educational_system" form:"educational_system" binding:"required"` // 学制
	AdmissionDate     string `json:"admission_date" form:"admission_date" binding:"required"`         // 入学日期
	GraduationDate    string `json:"graduation_date" form:"graduation_date" binding:"required"`       // 毕业日期
	Status            string `json:"status" form:"status" binding:"required"`                         // 状态（是否结业）
	StudentAvatar     string `json:"student_avatar" form:"student_avatar" binding:"required"`         // 照片
}

func (service *StudentAuthService) StudentAuth(university *model.University) serializer.Response {
	if university.UniversityName != service.University {
		return serializer.Response{
			Code: 403,
			Msg:  "权限不足，无权颁发其他学校学历证书",
		}
	}

	var message = model.Message{
		MsgType:      model.EducationalQualifications,
		Description:  fmt.Sprintf("%s 请求学生认证", university.UniversityName),
		StudentAcMsg: model.StudentAcMsg{},
		EducationalAcMsg: model.EducationalAcMsg{
			UniversityID:      university.ID,
			Name:              service.Name,
			Sex:               service.Sex,
			Ethnic:            service.Ethnic,
			Birthday:          time.Unix(service.Birthday, 0),
			CardCode:          service.CardCode,
			EducationCategory: service.EducationCategory,
			Level:             service.Level,
			University:        service.University,
			Professional:      service.Professional,
			LearningFormat:    service.LearningFormat,
			EducationalSystem: service.EducationalSystem,
			AdmissionDate:     service.AdmissionDate,
			GraduationDate:    service.GraduationDate,
			Status:            service.Status,
			StudentAvatar:     service.StudentAvatar,
		},
	}

	err := model.DB.Save(&message).Error
	if err != nil {
		return serializer.DBErr("消息保存失败", err)
	}

	//university.Status = model.Authenticating
	//err = model.DB.Save(&university).Error
	//if err != nil {
	//	return serializer.DBErr("消息保存失败", err)
	//}

	sc := model.Scheduler{
		UniversityName:   message.EducationalAcMsg.University,
		UniversityUserID: message.EducationalAcMsg.UniversityID,
		MessageID:        message.ID,
		CertificationID:  0,
		Status:           model.WAIT,
		StudentName:      message.EducationalAcMsg.Name,
	}
	err = model.DB.Create(&sc).Error
	if err != nil {
		return serializer.DBErr("保存失败", err)
	}

	return serializer.Response{
		Code: 0,
		Msg:  "提交成功，请等待审核",
	}
}

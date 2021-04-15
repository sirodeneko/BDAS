/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:university_certificate_list_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"singo/model"
	"singo/serializer"
	"time"
)

type StudentAuthListService struct {
	Limit     int    `json:"limit" form:"limit"`
	Offset    int    `json:"offset" form:"offset"`
	TimeLimit int64  `json:"time_limit" form:"time_limit"`
	Name      string `json:"name" form:"name"`
	Status    int    `json:"status" form:"status"`
}

func (service *StudentAuthListService) StudentAuthList(university *model.University) serializer.Response {
	//var msgs []model.Message
	var scs []model.Scheduler
	var db *gorm.DB
	total := 0

	if service.Limit == 0 {
		service.Limit = 15
	}

	db = model.DB.Where("university_name = ?", university.UniversityName)
	if service.TimeLimit != 0 {
		t := time.Unix(service.TimeLimit, 0)
		y := t.Year()
		m := t.Month()
		y1 := y
		m1 := m + 1
		if m1 == 13 {
			m1 = 1
			y1++
		}
		db = db.Where("created_at between ? and ?", fmt.Sprintf("%d-%d-1 00:00:00", y, m), fmt.Sprintf("%d-%d-1 00:00:00", y1, m1))
	}
	if service.Name != "" {
		db = db.Where("student_name like ? ", fmt.Sprintf("%s%%", service.Name))
	}

	if service.Status != 0 {
		db = db.Where("status = ? ", service.Status)
	}

	err := db.Model(model.Scheduler{}).Count(&total).Error
	if err != nil {
		return serializer.DBErr("数据库链接错误", err)
	}

	err = db.Order("id desc").Limit(service.Limit).Offset(service.Offset).Find(&scs).Error
	if err != nil {
		return serializer.DBErr("数据库链接错误", err)
	}

	//ids:=make([]uint,0,2)
	//for _,i:=range scs{
	//	ids=append(ids, i.MessageID)
	//}
	//
	//err=model.DB.Where("id in (?)",ids).
	//	Unscoped().
	//	Preload("EducationalAcMsg").
	//	Find(&msgs).Error
	//
	//if err != nil {
	//	return serializer.DBErr("数据库链接错误", err)
	//}

	return serializer.BuildListResponse(serializer.BuildSchedulers(scs), uint(total))
}

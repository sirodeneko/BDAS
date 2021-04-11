/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:inbox_list_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

type InboxListService struct {
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
	//Op int `json:"op" form:"op"`
}

func (service *InboxListService) InboxList(c *gin.Context) serializer.Response {
	var inboxes []model.Inbox
	total := 0
	session := sessions.Default(c)
	utype := session.Get("user_type")
	uid := session.Get("user_id")
	if utype == nil || uid == nil {
		return serializer.CheckLogin()
	}
	if service.Limit == 0 {
		service.Limit = 12
	}

	if err := model.DB.Model(model.Inbox{}).Where("user_type = ? AND user_id = ?", utype, uid).Count(&total).Error; err != nil {
		return serializer.DBErr("数据库链接错误", err)
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Offset).Where("user_type = ? AND user_id = ?", utype, uid).Order("id desc").Find(&inboxes).Error; err != nil {
		return serializer.DBErr("数据库链接错误", err)
	}
	//time.Sleep(3*time.Second)

	return serializer.BuildListResponse(serializer.BuildInboxes(inboxes), uint(total))
}

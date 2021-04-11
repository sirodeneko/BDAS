/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:admin_send_inbox_service.go
 * Date:2021/4/12 上午12:21
 * Author:sirodeneko
 */

package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

type AdminSendInboxService struct {
	UserType string `json:"u_type" form:"u_type" binding:"required"`
	UserID   uint   `json:"id" form:"id" binding:"required"`
	Body     string `json:"body" form:"body" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required" `
}

func (service *AdminSendInboxService) AdminSendInbox(c *gin.Context) serializer.Response {
	inbox := model.Inbox{
		UserType: service.UserType,
		UserID:   service.UserID,
		Body:     service.Body,
		Title:    service.Title,
		State:    0,
	}
	err := model.DB.Create(&inbox).Error
	if err != nil {
		return serializer.DBErr("数据库插入失败", err)
	}
	return serializer.Response{}
}

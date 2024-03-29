/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:inbox_looked_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"singo/model"
	"singo/serializer"
)

type InboxLookedService struct {
	ID string `json:"inbox_id" form:"inbox_id"`
}

func (service *InboxLookedService) InboxLooked() serializer.Response {
	model.DB.Model(&model.Inbox{}).Where("id = ?", service.ID).Update("state", 1)
	return serializer.Response{}
}

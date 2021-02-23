package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

type InboxListUnreadService struct {
}

func (service *InboxListUnreadService) InboxListUnread(c *gin.Context) serializer.Response {

	total := 0
	session := sessions.Default(c)
	utype := session.Get("user_type")
	uid := session.Get("user_id")
	if utype == nil || uid == nil {
		return serializer.CheckLogin()
	}

	if err := model.DB.Model(model.Inbox{}).Where("user_type = ? AND user_id = ? AND state = 0", utype, uid).Count(&total).Error; err != nil {
		return serializer.DBErr("数据库链接错误", err)
	}

	return serializer.Response{
		Data: total,
	}
}

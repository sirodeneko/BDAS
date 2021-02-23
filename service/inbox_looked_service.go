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

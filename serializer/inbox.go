/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:inbox.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package serializer

import "singo/model"

// Inbox  各种通知
type Inbox struct {
	ID        uint   `json:"id"`
	CreatedAt int64  `json:"created_at"`
	Body      string `json:"body"`
	Title     string `json:"title"`
	State     int    `json:"state"` //0为未读 1为已读
}

// BuildInbox 序列化消息
func BuildInbox(inbox model.Inbox) Inbox {
	return Inbox{
		ID:        inbox.ID,
		CreatedAt: inbox.CreatedAt.Unix(),
		Body:      inbox.Body,
		Title:     inbox.Title,
		State:     inbox.State,
	}
}

// BuildInboxes 序列化消息列表
func BuildInboxes(items []model.Inbox) []Inbox {
	var inboxes []Inbox

	for _, item := range items {
		inbox := BuildInbox(item)
		inboxes = append(inboxes, inbox)
	}
	return inboxes
}

// BuildInboxResponse 序列化消息响应
func BuildInboxResponse(inbox model.Inbox) Response {
	return Response{
		Code: 0,
		Data: BuildInbox(inbox),
	}
}

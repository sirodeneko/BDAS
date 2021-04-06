/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:inbox.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package model

import (
	"github.com/jinzhu/gorm"
)

// Inbox  消息通知
type Inbox struct {
	gorm.Model
	UserType string `gorm:"index:idx_member"`
	UserID   uint   `gorm:"index:idx_member"`
	Body     string
	Title    string
	State    int //0为未读 1为已读
}

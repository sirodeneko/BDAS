package model

import (
	"github.com/jinzhu/gorm"
)

// Inbox  消息通知
type Inbox struct {
	gorm.Model
	UserType string
	UserID   uint
	Body     string
	Title    string
	State    int //0为未读 1为已读
}

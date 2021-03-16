package model

import "github.com/jinzhu/gorm"

// Scheduler 调度消息记录
type Scheduler struct {
	gorm.Model
	UniversityName   string `gorm:"index"`
	UniversityUserID uint   `gorm:"index"`
	MessageID        uint   `gorm:"index"`
	CertificationID  uint
	Err              string
	Status           int
	StudentName      string
}

const (
	WAIT = iota + 1
	EXECUTING
	FAILED
	SUCCESS
	NOPASS
)

func (s *Scheduler) GetStatus() string {
	switch s.Status {
	case WAIT:
		return "等待"
	case EXECUTING:
		return "运行中"
	case FAILED:
		return "失败"
	case SUCCESS:
		return "成功"
	case NOPASS:
		return "不通过"
	default:
		return "未知"
	}
}

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
}

const (
	WAIT      int = 1
	EXECUTING int = 2
	FAILED    int = 3
	SUCCESS   int = 4
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
	default:
		return "未知"
	}
}

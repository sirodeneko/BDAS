package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Message  各种需要管理员进行处理的通知
type Message struct {
	gorm.Model
	MsgType            string
	Description        string
	StudentAcMsgID     uint
	EducationalAcMsgID uint
	StudentAcMsg       StudentAcMsg
	EducationalAcMsg   EducationalAcMsg
}

type StudentAcMsg struct {
	gorm.Model
	UserId       uint
	Name         string
	CardCode     string
	FrontFaceImg string
	BackFaceImg  string
}

type EducationalAcMsg struct {
	gorm.Model
	UniversityID      uint      // 消息发送者的id
	Name              string    // 姓名
	Sex               uint      // 1男 2女
	Ethnic            string    // 民族
	Birthday          time.Time // 生日
	CardCode          string    // 身份证号
	EducationCategory string    // 学历类别
	Level             string    // 层次
	University        string    // 学校
	Professional      string    // 专业
	LearningFormat    string    // 学习形式
	EducationalSystem string    // 学制
	AdmissionDate     string    // 入学日期
	GraduationDate    string    // 毕业日期
	Status            string    // 状态（是否结业）
	StudentAvatar     string    // 照片
}

const (
	// StudentAccreditation 学生认证请求
	StudentAccreditation string = "student accreditation"
	// EducationalQualifications 学历认证请求
	EducationalQualifications string = "educational qualifications"
)

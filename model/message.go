package model

import "github.com/jinzhu/gorm"

// Message  各种需要管理员进行处理的通知
type Message struct {
	gorm.Model
	MsgType     string
	Description string
	SAM         StudentAcMsg
	EAM         EducationalAcMsg
}

type StudentAcMsg struct {
	UserId       int
	Name         int
	CardCode     string
	FrontFaceImg string
	BackFaceImg  string
}

type EducationalAcMsg struct {
}

const (
	// StudentAccreditation 学生认证请求
	StudentAccreditation string = "student accreditation"
	// EducationalQualifications 学历认证请求
	EducationalQualifications string = "educational qualifications"
)

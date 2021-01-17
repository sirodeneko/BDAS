package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// University 大学管理员模型
type University struct {
	gorm.Model
	UniversityName string // 学校名称
	UserName       string // 账号
	PasswordDigest string // 密码
	Nickname       string // 账号名称
	Status         string // 账号状态
	// Description    string // 描述
}

// GetUniversity 用ID获取学校管理员
func GetUniversity(ID interface{}) (University, error) {
	var university University
	result := DB.First(&university, ID)
	return university, result.Error
}

// SetPassword 设置密码
func (university *University) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	university.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (university *University) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(university.PasswordDigest), []byte(password))
	return err == nil
}

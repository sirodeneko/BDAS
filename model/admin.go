package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"singo/util"
)

// Admin 管理员模型
type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
}
type MyInt string

const (
	AdminType      string = "admin"
	UserType       string = "user"
	UniversityType string = "university"
)

func GetUserWithType(uType string, ID interface{}) (interface{}, error) {
	switch uType {
	case AdminType:
		admin, err := GetAdmin(ID)
		return &admin, err
	case UserType:
		user, err := GetUser(ID)
		return &user, err
	case UniversityType:
		university, err := GetUniversity(ID)
		return &university, err
	default:
		err := errors.New("类型匹配失败")
		util.Log().Warning(uType+"类型匹配失败", err)
		return nil, err
	}
}

// GetAdmin 用ID获取管理员
func GetAdmin(ID interface{}) (Admin, error) {
	var admin Admin
	result := DB.First(&admin, ID)
	return admin, result.Error
}

// SetPassword 设置密码
func (admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	admin.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordDigest), []byte(password))
	return err == nil
}

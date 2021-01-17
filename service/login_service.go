package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// LoginService 统一用户登录的服务
type LoginService struct {
	UType          string `form:"u_type" json:"u_type" binding:"required"`
	UniversityName string `form:"university_name" json:"university_name"`
	UserName       string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password       string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

const (
	AdminType      string = "admin"
	UserType       string = "user"
	UniversityType string = "university"
)

// setSession 设置session
func (service *LoginService) setSession(c *gin.Context, id uint, t string) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", id)
	s.Set("type", t)
	s.Save()
}

func (service *LoginService) adminLogin(c *gin.Context) serializer.Response {

}
func (service *LoginService) userLogin(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 设置session
	service.setSession(c, user.ID, service.UType)

	return serializer.BuildUserResponse(user)
}
func (service *LoginService) universityLogin(c *gin.Context) serializer.Response {

}

// Login 用户登录函数
func (service *LoginService) Login(c *gin.Context) serializer.Response {
	switch service.UType {
	case AdminType:
		return service.adminLogin(c)
	case UserType:
		return service.userLogin(c)
	case UniversityType:
		return service.universityLogin(c)
	default:
		return serializer.ParamErr("账户类型错误", nil)
	}
}

package api

import (
	"singo/model"
	"singo/serializer"
	"singo/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	if user, ok := CurrentUser(c).(*model.User); ok {
		res := serializer.BuildUserResponse(*user)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CheckLogin())
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// Login 统一登入接口
func Login(c *gin.Context) {
	var service service.LoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserModify 用户修改个人信息
func UserModify(c *gin.Context) {
	var service service.UserModifyService
	var ID uint

	if user, ok := CurrentUser(c).(*model.User); ok {
		ID = user.ID
	} else {
		c.JSON(200, serializer.CheckLogin())
	}

	if err := c.ShouldBind(&service); err == nil {
		res := service.UserModify(ID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

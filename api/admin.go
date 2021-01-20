package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
	"singo/service"
)

// AdminRegister 管理员注册接口
func AdminRegister(c *gin.Context) {
	var service service.AdminRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminMe 管理员用户详情
func AdminMe(c *gin.Context) {
	if admin, ok := CurrentUser(c).(*model.Admin); ok {
		res := serializer.BuildAdminResponse(*admin)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CheckLogin())
	}
}

// AdminLogout 管理员用户登出
func AdminLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// AdminModify 管理员修改自己基本信息
func AdminModify(c *gin.Context) {
	var service service.AdminModifyService
	var ID uint

	if user, ok := CurrentUser(c).(*model.Admin); ok {
		ID = user.ID
	} else {
		c.JSON(200, serializer.CheckLogin())
	}

	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminModify(ID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminModifyUser 管理员修改普通用户信息
func AdminModifyUser(c *gin.Context) {
	var service service.AdminModifyUserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminModifyUser()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminModifyUniversity 管理员修改学校用户信息
func AdminModifyUniversity(c *gin.Context) {
	var service service.AdminModifyUniversityService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminModifyUniversity()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

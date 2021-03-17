package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
	"singo/service"
)

// UniversityRegister 学校用户注册接口
func UniversityRegister(c *gin.Context) {
	var service service.UniversityRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UniversityMe 学校用户详情
func UniversityMe(c *gin.Context) {
	if university, ok := CurrentUser(c).(*model.University); ok {
		res := serializer.BuildUniversityResponse(*university)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CheckLogin())
	}
}

// UniversityLogout 学校用户登出
func UniversityLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// UniversityModify 学校基本信息修改
func UniversityModify(c *gin.Context) {
	var service service.UniversityModifyService
	var ID uint

	if user, ok := CurrentUser(c).(*model.University); ok {
		ID = user.ID
	} else {
		c.JSON(200, serializer.CheckLogin())
	}

	if err := c.ShouldBind(&service); err == nil {
		res := service.UniversityModify(ID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// StudentAuth 学校进行学生的学历信息提交
func StudentAuth(c *gin.Context) {
	var service service.StudentAuthService

	var user *model.University

	if u, ok := CurrentUser(c).(*model.University); ok {
		user = u
	} else {
		c.JSON(200, serializer.CheckLogin())
	}

	if err := c.ShouldBind(&service); err == nil {
		res := service.StudentAuth(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// StudentAuthList 获取消息列表,支持查询
func StudentAuthList(c *gin.Context) {
	var service service.StudentAuthListService

	var user *model.University

	if u, ok := CurrentUser(c).(*model.University); ok {
		user = u
	} else {
		c.JSON(200, serializer.CheckLogin())
	}

	if err := c.ShouldBind(&service); err == nil {
		res := service.StudentAuthList(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// StudentAuthMsg 通过msg_id 获取msg信息
func StudentAuthMsg(c *gin.Context) {
	var service service.StudentAuthMsgService

	var user *model.University
	if u, ok := CurrentUser(c).(*model.University); ok {
		user = u
	} else {
		c.JSON(200, serializer.CheckLogin())
	}

	if err := c.ShouldBind(&service); err == nil {
		res := service.StudentAuthMsg(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

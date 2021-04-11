/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:admin.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

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

// MsgList 获取消息列表
func MsgList(c *gin.Context) {
	var service service.MsgListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MsgList()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminAuthUser 认证用户
func AdminAuthUser(c *gin.Context) {
	var service service.AdminAuthUserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminAuthUser()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminACStudent 管理员通过学校的认证请求
func AdminACStudent(c *gin.Context) {
	var service service.AdminACStudentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminACStudent()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 管理员通过用户类型和id查询到user
func AdminGetUser(c *gin.Context) {
	var service service.AdminGetUserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminGetUser()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 管理员获取一条msg
func GetAMsg(c *gin.Context) {
	var service service.GetAMsgService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetAMsg()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 管理员发送通知消息
func AdminSendInbox(c *gin.Context) {
	var service service.AdminSendInboxService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminSendInbox(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

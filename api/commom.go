/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:commom.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// FileUpload 文件上传
func FileUpload(c *gin.Context) {
	var service service.FileUploadService
	f, err := c.FormFile("f1")

	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.FileUpload(c, f)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 文件下载
func FileDownload(c *gin.Context) {
	var service service.FileDownloadService
	FileName := c.Param("filename")
	res := service.FileDownload(c, FileName)
	if res.Code != 0 {
		c.JSON(200, res)
	}
}

// InboxList 获取收件箱邮件
func InboxList(c *gin.Context) {
	var service service.InboxListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InboxList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// InboxListUnread 获取未读邮件的数量
func InboxListUnread(c *gin.Context) {
	var service service.InboxListUnreadService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InboxListUnread(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// InboxLooked 查看邮件
func InboxLooked(c *gin.Context) {
	var service service.InboxLookedService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InboxLooked()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

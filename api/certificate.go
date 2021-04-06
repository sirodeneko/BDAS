/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:certificate.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CertificateWithAddress 根据地址等信息查询
func CertificateWithAddress(c *gin.Context) {
	var service service.CertificateWithAddressService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CertificateWithAddress()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CertificateWithFile 根据文件判断真伪
func CertificateWithFile(c *gin.Context) {
	var service service.CertificateWithFileService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CertificateWithFile()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

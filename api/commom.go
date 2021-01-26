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

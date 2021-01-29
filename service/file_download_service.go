package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"regexp"
	"singo/serializer"
)

type FileDownloadService struct{}

func (service *FileDownloadService) checkFileName(fileName string) bool {
	if fileName == "" {
		return false
	}
	re, _ := regexp.Compile("[\\\\/:*?\"<>|]")
	if re.MatchString(fileName) {
		return false
	}
	return true
}

func (service *FileDownloadService) FileDownload(c *gin.Context, fileName string) serializer.Response {
	if !service.checkFileName(fileName) {
		return serializer.ParamErr("文件名不合法", nil)
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	dst := path.Join("./static/file", fileName)
	c.File(dst)
	return serializer.Response{Code: 0}
}

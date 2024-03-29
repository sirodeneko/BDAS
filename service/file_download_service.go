/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:file_download_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"regexp"
	"singo/serializer"
	"singo/util"
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
	dst := path.Join(util.FileURL, fileName)
	c.File(dst)
	return serializer.Response{Code: 0}
}

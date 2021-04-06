/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:certification_file_download_service.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"regexp"
	"singo/model"
	"singo/serializer"
	"singo/util"
)

type CertificationFileDownloadService struct{}

func (service *CertificationFileDownloadService) checkFileName(fileName string) bool {
	if fileName == "" {
		return false
	}
	re, _ := regexp.Compile("[\\\\/:*?\"<>|]")
	if re.MatchString(fileName) {
		return false
	}
	return true
}

func (service *CertificationFileDownloadService) CertificationFileDownload(c *gin.Context, user *model.User, fileName string) serializer.Response {
	if !service.checkFileName(fileName) {
		return serializer.ParamErr("文件名不合法", nil)
	}
	count := 0
	model.DB.Model(&model.Certification{}).Where("card_code = ? AND url = ?", user.CardCode, fileName).Count(&count)
	if count == 0 {
		return serializer.Response{
			Code: 403,
			Msg:  "非法访问",
		}
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	dst := path.Join(util.SaveURL, fileName)
	c.File(dst)
	return serializer.Response{Code: 0}
}

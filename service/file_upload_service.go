package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"path"
	"singo/serializer"
	"singo/util"
)

type FileUploadService struct {
}

func (service *FileUploadService) FileUpload(c *gin.Context, f *multipart.FileHeader) serializer.Response {
	newFileName := uuid.Must(uuid.NewRandom()).String() + path.Ext(f.Filename)
	dst := path.Join("./static/file", newFileName)

	if f.Size > 5<<20 {
		return serializer.Response{
			Code: 403,
			Msg:  "文件太大",
		}
	}

	err := c.SaveUploadedFile(f, dst)
	if err != nil {
		fmt.Println(err)
		util.Log().Error("%s 文件保存失败", newFileName)
		return serializer.Err(50000, "文件保存失败", err)
	}
	data := make(map[string]string)
	data["fileName"] = newFileName
	return serializer.Response{
		Code: 0,
		Data: data,
	}
}

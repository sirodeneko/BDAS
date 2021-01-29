package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
)

type UserAuthService struct {
	Name         string `json:"name" form:"name" binding:"required"`
	CardCode     string `json:"card_code" form:"card_code" binding:"required"`
	FrontFaceImg string `json:"front_face_img" form:"front_face_img" binding:"required"`
	BackFaceImg  string `json:"back_face_img" form:"back_face_img" binding:"required"`
}

func (service *UserAuthService) UserAuth(id uint) serializer.Response {
	var user model.User
	//找到用户
	err := model.DB.First(&user, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询用户失败",
			Error: err.Error(),
		}
	}

	if user.CardCode != "" {
		return serializer.Response{
			Code: 403,
			Msg:  "已认证通过，请勿重复认证",
		}
	}

	var mssage = model.Message{
		MsgType:     model.StudentAccreditation,
		Description: fmt.Sprintf("%d 用户请求生份认证", id),
		StudentAcMsg: model.StudentAcMsg{
			UserId:       id,
			Name:         service.Name,
			CardCode:     service.CardCode,
			FrontFaceImg: service.FrontFaceImg,
			BackFaceImg:  service.BackFaceImg,
		},
		EducationalAcMsg: model.EducationalAcMsg{},
	}

	err = model.DB.Save(&mssage).Error
	if err != nil {
		return serializer.DBErr("消息保存失败", err)
	}

	user.Status = model.Authenticating
	err = model.DB.Save(&user).Error
	if err != nil {
		return serializer.DBErr("消息保存失败", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "提交成功，请等待审核",
	}
}

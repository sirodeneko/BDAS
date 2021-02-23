package service

import (
	"singo/model"
	"singo/serializer"
)

type AdminAuthUserService struct {
	//UserID   uint   `json:"user_id" form:"user_id" binding:"required"`
	//Name     string `json:"name" form:"name" binding:"required"`
	//CardCode string `json:"card_code" form:"card_code" binding:"required"`
	MsgID uint `json:"msg_id" form:"msg_id" binding:"required"`
}

func (service *AdminAuthUserService) AdminAuthUser() serializer.Response {
	var user model.User
	var message model.Message

	err := model.DB.Preload("StudentAcMsg").First(&message, service.MsgID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "消息查询失败",
			Error: err.Error(),
		}
	}
	if message.MsgType != model.StudentAccreditation {
		return serializer.Response{
			Code: 404,
			Msg:  "非法请求",
		}
	}
	//找到用户
	err = model.DB.First(&user, message.StudentAcMsg.UserId).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询用户失败",
			Error: err.Error(),
		}
	}
	user.Status = model.Active
	user.Name = message.StudentAcMsg.Name
	user.CardCode = message.StudentAcMsg.CardCode
	err = model.DB.Save(&user).Error
	if err != nil {
		return serializer.DBErr("用户信息保存失败", err)
	}
	// 删除message对象 不删除关联对象
	model.DB.Delete(message)
	// 删除其关联的标记，不删除任何对象
	//model.DB.Model(&message).Association("StudentAcMsg").Clear()

	return serializer.Response{
		Msg: "用户激活成功",
	}
}

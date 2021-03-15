package service

import (
	"singo/model"
	"singo/serializer"
)

type AdminAuthUserService struct {
	//UserID   uint   `json:"user_id" form:"user_id" binding:"required"`
	//Name     string `json:"name" form:"name" binding:"required"`
	//CardCode string `json:"card_code" form:"card_code" binding:"required"`
	MsgID uint   `json:"msg_id" form:"msg_id" binding:"required"`
	Op    uint   `json:"op" form:"op"`
	Msg   string `json:"msg" form:"msg"`
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

	if user.CardCode != "" {
		return serializer.Response{
			Code: 200,
			Msg:  "已认证通过，请勿重复认证",
		}
	}

	// 处理不通过请求
	if service.Op != 0 {
		user.Status = model.Inactive
		err = model.DB.Save(&user).Error
		if err != nil {
			return serializer.DBErr("用户信息保存失败", err)
		} else {
			var inbox = model.Inbox{
				UserType: model.UserType,
				UserID:   user.ID,
				Body: "您好，您的身份认证请求经管理员审核<div style=\"color:red;\">不通过</div>，原因如下：<br>" +
					service.Msg +
					"<br>感谢您使用本平台，祝您生活愉快",
				Title: "身份认证不通过",
				State: 0,
			}
			model.DB.Save(&inbox)
			return serializer.Response{
				Code: 0,
				Msg:  "ok",
			}
		}
	}

	// 处理通过请求
	user.Status = model.Active
	user.Name = message.StudentAcMsg.Name
	user.CardCode = message.StudentAcMsg.CardCode
	err = model.DB.Save(&user).Error
	if err != nil {
		return serializer.DBErr("用户信息保存失败", err)
	}

	var inbox = model.Inbox{
		UserType: model.UserType,
		UserID:   user.ID,
		Body:     "您好，您的身份认证请求经管理员审核<div style=\"color:red;\">通过</div> <br>感谢您使用本平台，祝您生活愉快",
		Title:    "身份认证通过",
		State:    0,
	}
	model.DB.Save(&inbox)
	//model.DB.Delete(message.StudentAcMsg)
	// 删除message对象 不删除关联对象
	model.DB.Delete(message)
	// 删除其关联的标记，不删除任何对象
	//model.DB.Model(&message).Association("StudentAcMsg").Clear()

	return serializer.Response{
		Msg: "ok",
	}
}

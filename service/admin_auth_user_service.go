package service

import (
	"singo/model"
	"singo/serializer"
)

type AdminAuthUserService struct {
	UserID   uint   `json:"user_id" form:"user_id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	CardCode string `json:"card_code" form:"card_code" binding:"required"`
}

func (service *AdminAuthUserService) AdminAuthUser() serializer.Response {
	var user model.User
	//找到用户
	err := model.DB.First(&user, service.UserID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询用户失败",
			Error: err.Error(),
		}
	}
	user.Status = model.Active
	user.Name = service.Name
	user.CardCode = service.CardCode
	err = model.DB.Save(&user).Error
	if err != nil {
		return serializer.DBErr("用户信息保存失败", err)
	}
	return serializer.Response{
		Msg: "用户激活成功",
	}
}

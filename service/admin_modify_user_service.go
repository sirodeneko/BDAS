package service

import (
	"singo/model"
	"singo/serializer"
)

// AdminModifyUserService 管理员修改信息的服务
type AdminModifyUserService struct {
	ID          uint   `form:"id" json:"id" binding:"required"`
	Nickname    string `form:"nickname" json:"nickname" binding:"omitempty,min=2,max=30"`
	NewPassword string `form:"new_password" json:"new_password" binding:"omitempty,min=2,max=30"`
	Status      string `json:"status" form:"status" `
	CardCode    string `json:"card_code" form:"card_code"`
}

// Change 用户修改信息
func (service *AdminModifyUserService) AdminModifyUser() serializer.Response {
	var user model.User
	//找到用户
	err := model.DB.First(&user, service.ID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询用户失败",
			Error: err.Error(),
		}
	}

	// 如果更新了密码
	if service.NewPassword != "" {
		if err := user.SetPassword(service.NewPassword); err != nil {
			return serializer.Err(
				serializer.CodeEncryptError,
				"密码加密失败",
				err,
			)
		}
	}
	user.Nickname = service.Nickname
	user.Status = service.Status
	user.CardCode = service.CardCode

	err = model.DB.Model(&user).Update(&user).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeDBError,
			Msg:   "用户信息保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}

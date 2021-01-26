package service

import (
	"singo/model"
	"singo/serializer"
)

// AdminModifyService 管理员修改信息的服务
type AdminModifyService struct {
	Nickname    string `form:"nickname" json:"nickname" binding:"omitempty,min=2,max=30"`
	OldPassword string `form:"old_password" json:"old_password"`
	NewPassword string `form:"new_password" json:"new_password" binding:"omitempty,min=2,max=30"`
}

// Change 用户修改信息
func (service *AdminModifyService) AdminModify(ID uint) serializer.Response {
	var admin model.Admin
	//找到用户
	err := model.DB.First(&admin, ID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询用户失败",
			Error: err.Error(),
		}
	}

	// 如果更新了密码
	if service.NewPassword != "" {
		if admin.CheckPassword(service.OldPassword) == false {
			return serializer.ParamErr("账号或密码错误", nil)
		}

		if err := admin.SetPassword(service.NewPassword); err != nil {
			return serializer.Err(
				serializer.CodeEncryptError,
				"密码加密失败",
				err,
			)
		}
	}
	admin.Nickname = service.Nickname

	err = model.DB.Model(&admin).Update(&admin).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeDBError,
			Msg:   "用户信息保存失败",
			Error: err.Error(),
		}
	}
	model.DB.Find(&admin)
	return serializer.Response{
		Data: serializer.BuildAdmin(admin),
	}
}

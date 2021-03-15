package service

import (
	"singo/model"
	"singo/serializer"
)

type AdminGetUserService struct {
	UType string `form:"u_type" json:"u_type" binding:"required"`
	UID   uint   `form:"uid" json:"uid" binding:"required"`
}

func (service *AdminGetUserService) AdminGetUser() serializer.Response {
	switch service.UType {
	case model.UserType:
		return getUser(service.UID)
	case model.UniversityType:
		return getUniversity(service.UID)
	default:
		return serializer.Response{
			Code:  403,
			Msg:   "类型错误",
			Error: "admin_get_user:类型错误",
		}
	}

}

func getUser(id uint) serializer.Response {
	user, err := model.GetUser(id)
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "未查询到用户",
			Error: err.Error(),
		}
	}

	return serializer.BuildUserResponse(user)

}

func getUniversity(id uint) serializer.Response {
	user, err := model.GetUniversity(id)
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "未查询到用户",
			Error: err.Error(),
		}
	}

	return serializer.BuildUniversityResponse(user)
}

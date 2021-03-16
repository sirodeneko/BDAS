package service

import (
	"singo/model"
	"singo/scheduler"
	"singo/serializer"
	"singo/util"
)

type AdminACStudentService struct {
	MsgID string `json:"msg_id" form:"msg_id" binding:"required"`
	Op    uint   `json:"op" form:"op"`
	Msg   string `json:"msg" form:"msg"`
}

// var jsonKey string

func (service *AdminACStudentService) AdminACStudent() serializer.Response {
	var message model.Message

	err := model.DB.Preload("EducationalAcMsg").First(&message, service.MsgID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "消息查询失败",
			Error: err.Error(),
		}
	}
	if message.MsgType != model.EducationalQualifications {
		return serializer.Response{
			Code: 404,
			Msg:  "非法请求",
		}
	}

	var sc model.Scheduler

	// 查询被软删除的记录
	err = model.DB.Unscoped().Where("message_id = ? ", message.ID).First(&sc).Error
	if err != nil {
		util.Log().Error("数据库查询失败: %v", err)
		return serializer.DBErr("数据库查询失败", err)
	}

	// 处理不通过请求
	if service.Op != 0 {
		sc.Status = model.NOPASS
		err = model.DB.Save(&sc).Error

		if err != nil {
			return serializer.DBErr("用户信息保存失败", err)
		} else {
			var inbox = model.Inbox{
				UserType: model.UniversityType,
				UserID:   message.EducationalAcMsg.UniversityID,
				Body: "您好，您的学生学历认证请求经管理员审核<div style=\"color:red;\">不通过</div>，原因如下：<br>" +
					service.Msg +
					"<br>感谢您使用本平台，祝您生活愉快",
				Title: "学历认证不通过",
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

	sc.Status = model.EXECUTING
	model.DB.Save(&sc)

	err = scheduler.Submit(&message)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "提交失败",
			Error: err.Error(),
		}
	}

	model.DB.Delete(message)

	return serializer.Response{
		Msg: "提交成功",
	}
}

//
//func model2PictureInfo(msg *model.EducationalAcMsg) util.PictureInfo {
//	return util.PictureInfo{
//		CreatedAt:         msg.CreatedAt.Unix(),
//		Name:              msg.Name,
//		Sex:               msg.Sex,
//		Ethnic:            msg.Ethnic,
//		Birthday:          msg.Birthday.Unix(),
//		CardCode:          msg.CardCode,
//		EducationCategory: msg.EducationCategory,
//		Level:             msg.Level,
//		University:        msg.University,
//		Professional:      msg.Professional,
//		LearningFormat:    msg.LearningFormat,
//		EducationalSystem: msg.EducationalSystem,
//		AdmissionDate:     msg.AdmissionDate,
//		GraduationDate:    msg.GraduationDate,
//		Status:            msg.Status,
//		StudentAvatar:     msg.StudentAvatar,
//	}
//}
//
//func getJsonKey() string {
//	if jsonKey == "" {
//		jsonKey = os.Getenv("JSON_KEY")
//	}
//	return jsonKey
//}

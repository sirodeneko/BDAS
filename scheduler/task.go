package scheduler

import (
	"encoding/json"
	"singo/model"
	"singo/util"
	"singo/vnt"
)

func caFileAndChain(args interface{}) {

	var sc model.Scheduler

	message, _ := args.(*model.Message)

	// 查询被软删除的记录
	err := model.DB.Unscoped().Where("message_id = ? ", message.ID).First(&sc).Error
	if err != nil {
		util.Log().Error("sc(message_id=%d)任务失败，数据库查询失败: %v", message.ID, err)
		return
	}
	sc.Status = model.EXECUTING
	model.DB.Save(&sc)

	pictureInfo := model2PictureInfo(&message.EducationalAcMsg)

	// 写入数据库
	certification := model.Certification{
		CardCode:     pictureInfo.CardCode,
		Level:        pictureInfo.Level,
		Professional: pictureInfo.Professional,
		Name:         pictureInfo.Name,
	}
	// 先保存主要数据
	err = model.DB.Save(&certification).Error
	if err != nil {
		util.Log().Error("sc(id=%d)任务失败，数据库查询失败: %v", sc.ID, err)
		sc.Status = model.FAILED
		sc.Err = err.Error()
		model.DB.Save(&sc)
		return
	}

	pictureInfo.FileID = certification.ID

	// 将信息写入sc
	sc.CertificationID = certification.ID

	url, err := pictureInfo.CreateImg()
	if err != nil {
		util.Log().Error("sc(id=%d)任务失败，认证文件生成失败: %v", sc.ID, err)
		sc.Status = model.FAILED
		sc.Err = err.Error()
		model.DB.Save(&sc)
		return
	}

	pictureInfo.FileUrl = url
	pictureInfo.FileHash = util.FileSHA256(util.SaveURL + url)

	pictureInfoByte, _ := json.Marshal(pictureInfo)
	pictureInfCrypt := util.AesEncrypt(string(pictureInfoByte), util.GetJsonKey())

	// 获取nonce
	nonce := vnt.GetNonce()
	// 签名
	signHex := vnt.Sign([]byte(pictureInfCrypt), nonce)
	// 广播上链
	txAddress, err := vnt.SendRawTransaction(signHex)
	if err != nil {
		util.Log().Error("sc(id=%d)任务失败，数据上链失败: %v", sc.ID, err)
		sc.Status = model.FAILED
		sc.Err = err.Error()
		model.DB.Save(&sc)
		return
	}

	// 写入数据库
	certification.Address = txAddress
	certification.Url = url
	err = model.DB.Save(certification).Error
	if err != nil {
		util.Log().Error("sc(id=%d)任务失败，数据上链成功但是保存失败: %v", sc.ID, err)
		sc.Status = model.FAILED
		sc.Err = err.Error()
		model.DB.Save(&sc)
		return
	}
	sc.Status = model.SUCCESS
	model.DB.Save(&sc)
	util.Log().Info("sc(id=%d)任务成功，数据上链成功: %v", sc.ID, err)
	// 删除msg
	//model.DB.Delete(message.EducationalAcMsg)
	//model.DB.Delete(message)
}

func caFileAndChainPanic(args interface{}) {}

func model2PictureInfo(msg *model.EducationalAcMsg) util.PictureInfo {
	return util.PictureInfo{
		CreatedAt:         msg.CreatedAt.Unix(),
		Name:              msg.Name,
		Sex:               msg.Sex,
		Ethnic:            msg.Ethnic,
		Birthday:          msg.Birthday.Unix(),
		CardCode:          msg.CardCode,
		EducationCategory: msg.EducationCategory,
		Level:             msg.Level,
		University:        msg.University,
		Professional:      msg.Professional,
		LearningFormat:    msg.LearningFormat,
		EducationalSystem: msg.EducationalSystem,
		AdmissionDate:     msg.AdmissionDate,
		GraduationDate:    msg.GraduationDate,
		Status:            msg.Status,
		StudentAvatar:     msg.StudentAvatar,
	}
}

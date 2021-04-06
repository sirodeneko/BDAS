/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:certification.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package serializer

import (
	"singo/model"
	"singo/util"
)

type Certification struct {
	ID           uint   `json:"id"`
	CreatedAt    int64  `json:"created_at"`
	CardCode     string `json:"card_code"`
	Address      string `json:"address"`
	Level        string `json:"level"`        // 层次 :本科
	Professional string `json:"professional"` // 专业 :xxx专业
}

// BuildCertification 序列化认证文件信息
func BuildCertification(certification model.Certification) Certification {
	return Certification{
		ID:           certification.ID,
		CreatedAt:    certification.CreatedAt.Unix(),
		CardCode:     util.HiddenCharacters(certification.CardCode),
		Address:      certification.Address,
		Level:        certification.Level,
		Professional: certification.Professional,
	}
}

// BuildCertifications 序列化认证文件信息列表
func BuildCertifications(items []model.Certification) []Certification {
	var certifications []Certification

	for _, item := range items {
		certification := BuildCertification(item)
		certifications = append(certifications, certification)
	}
	return certifications
}

// BuildCertificationResponse 序列化认证文件信息响应
func BuildCertificationResponse(certification model.Certification) Response {
	return Response{
		Data: BuildCertification(certification),
	}
}

/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:university.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package serializer

import "singo/model"

type University struct {
	ID             uint   `json:"id"`
	UniversityName string `json:"university_name"`
	UserName       string `json:"user_name"`
	Nickname       string `json:"nickname"`
	Status         string `json:"status"`
	CreatedAt      int64  `json:"created_at"`
}

// BuildUniversity 序列学校管理员
func BuildUniversity(university model.University) University {
	return University{
		ID:             university.ID,
		UniversityName: university.UniversityName,
		UserName:       university.UserName,
		Nickname:       university.Nickname,
		Status:         university.Status,
		CreatedAt:      university.CreatedAt.Unix(),
	}
}

// BuildUniversityResponse 序列化学校管理员响应
func BuildUniversityResponse(university model.University) Response {
	return Response{
		Data: BuildUniversity(university),
	}
}

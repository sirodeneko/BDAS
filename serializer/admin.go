/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:admin.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package serializer

import "singo/model"

type Admin struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

// BuildAdmin 序列管理员
func BuildAdmin(admin model.Admin) Admin {
	return Admin{
		ID:        admin.ID,
		UserName:  admin.UserName,
		Nickname:  admin.Nickname,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildAdminResponse(admin model.Admin) Response {
	return Response{
		Data: BuildAdmin(admin),
	}
}

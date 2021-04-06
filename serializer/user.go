/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:user.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package serializer

import (
	"singo/model"
	"singo/util"
)

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	CardCode  string `json:"card_code"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		CardCode:  util.HiddenCharacters(user.CardCode),
		Nickname:  user.Nickname,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}

package main

import (
	"singo/conf"
	"singo/model"
	"singo/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	test := model.Message{
		MsgType:      "sdada",
		StudentAcMsg: model.StudentAcMsg{UserId: 123, Name: "sfdsdf"},
	}
	model.DB.Create(&test)
	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}

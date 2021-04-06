/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:main.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package main

import (
	"singo/conf"
	"singo/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}

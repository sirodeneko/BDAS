/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:migration.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Admin{})
	DB.AutoMigrate(&Message{}, &StudentAcMsg{}, &EducationalAcMsg{})
	DB.AutoMigrate(&University{})
	DB.AutoMigrate(&Inbox{})
	DB.AutoMigrate(&Certification{})
	DB.AutoMigrate(&Scheduler{})
}

/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:certification.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package model

import (
	"github.com/jinzhu/gorm"
)

type Certification struct {
	gorm.Model
	CardCode     string `gorm:"index"`
	Name         string
	Url          string
	Address      string
	Level        string // 层次 :本科
	Professional string // 专业 :xxx专业
}

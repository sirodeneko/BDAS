package model

import (
	"github.com/jinzhu/gorm"
)

type Certification struct {
	gorm.Model
	CardCode     string
	Name         string
	Url          string
	Address      string
	Level        string // 层次 :本科
	Professional string // 专业 :xxx专业
}

package model

import (
	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Cnumber string `gorm:"unique;not null;primary_key"` // 将课程号设置为唯一且不为空
	Name    string `gorm:"unique;not null"`
}

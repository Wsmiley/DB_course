package model

type Student struct {
	Id      int    `gorm:"AUTO_INCREMENT"` // 设置id为自动递增
	Name    string `gorm:"unique;not null"`
	Snumber string `gorm:"unique;not null;primary_key"` // 将成员号设置为唯一且不为空
	Sex     string `gorm:"not null"`
	Dept    string `gorm:"not null"`
	Class   string `gorm:"not null"`
}

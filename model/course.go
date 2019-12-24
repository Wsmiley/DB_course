package model

type Course struct {
	Cnumber  string `gorm:"unique;not null;primary_key"` // 将课程号设置为唯一且不为空
	Name     string `gorm:"unique;not null"`
	Deptment string `gorm:"not null"`
}

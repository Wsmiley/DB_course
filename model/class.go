package model

type Class struct {
	ClassId   string `gorm:"unique;not null;primary_key"`
	ClassName string `gorm:"unique;not null"`
	Deptment  string `gorm:"not null"`
}

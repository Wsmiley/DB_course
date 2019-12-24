package model

type Score struct {
	Id      int     `gorm:"unique;not null;primary_key"`    // 成绩记录号
	Cnumber string  `gorm:";ForeignKey:CourseId;not null"`  //课程号
	Snumber string  `gorm:";ForeignKey:StudentId;not null"` //学生学号
	Score   float64 `gorm:"not null"`                       //成绩
}

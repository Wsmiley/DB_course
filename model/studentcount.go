package model

type StudentCount struct {
	Username string `gorm:";ForeignKey:StudentId;unique;not null"`
	Password string `gorm:"NOT NULL"`
	Name     string
}

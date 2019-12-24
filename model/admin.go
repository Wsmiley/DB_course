package model

//管理员
type Admins struct {
	Username string `gorm:"unique;NOT NULL"`
	Password string `gorm:"NOT NULL"`
	Name     string
}

package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;varchar(300)" json:"username" binding:"required,username"`
	Password string `gorm:"text" json:"password" binding:"required,min=6"`
}

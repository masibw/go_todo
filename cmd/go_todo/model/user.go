package model

import "time"

type User struct {
	Id        int       ` gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:id" json:"id"`
	Email     string    `binding:"required" gorm:"unique;not null;column:email" json:"email"`
	Password  string    `binding:"required" gorm:"column:password" json:"password"`
	CreatedAt time.Time ` gorm:"column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
	LastLogin time.Time `  gorm:"column:last_login" sql:"DEFAULT:current_timestamp" json:"last_login"`
}

package model

import "time"

type Todo struct {
	Id        int       ` gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:id" json:"id"`
	UserId    int       `gorm:"not null;column:user_id" json:"user_id"`
	Content   string    `binding:"required" gorm:"not null;column:content" json:"content"`
	Done bool `gorm:"not null; column:done" json:"done"`
	CreatedAt time.Time ` gorm:"column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
	UpdatedAt time.Time ` gorm:"column:updated_at" sql:"DEFAULT:current_timestamp" json:"updated_at"`
}

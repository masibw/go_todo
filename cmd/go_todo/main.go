package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/masibw/go_todo/pkg/db"
	"net/http"
	"time"
)

type User struct {
	Id        int       `binding:"required" gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:id"`
	Email     string    `binding:"required" gorm:"unique;not null;column:email"`
	Password  string    `binding:"required" gorm:"column:password"`
	CreatedAt time.Time ` binding:"required" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	LastLogin time.Time ` binding:"required" gorm:"column:last_login" sql:"DEFAULT:current_timestamp"`
}
type Todo struct {
	Id        int       `binding:"required" gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:id"`
	UserId    int       `binding:"required" gorm:"PRIMARY_KEY;column:user_id"`
	Content   string    `binding:"required" gorm:"not null;column:content"`
	CreatedAt time.Time `binding:"required" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `binding:"required" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
}


func main() {
	db := db.GormConnect()
	defer db.Close()
	r := gin.Default()

	db.Create(&User{Email: "test2@example.com", Password: "testpass"})
	db.Create(&Todo{Content: "このTodoアプリを完成させる", UserId: 1})
	var user User
	db.First(&user, 3)
	var todo Todo
	db.First(&todo, 1)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
			"todo": todo,
		})
	})
	r.Run()
}

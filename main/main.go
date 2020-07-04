package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct{
	Id int `binding:"required" gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:id"`
	Email string `binding:"required" gorm:"unique;not null;column:email"`
	Password string `binding:"required" gorm:"column:password"`
	CreatedAt time.Time ` binding:"required" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	LastLogin time.Time ` binding:"required" gorm:"column:last_login" sql:"DEFAULT:current_timestamp"`
}
type Todo struct{
	Id int `binding:"required" gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:id"`
	UserId int `binding:"required" gorm:"PRIMARY_KEY;column:user_id"`
	Content string `binding:"required" gorm:"not null;column:content"`
	CreatedAt time.Time `binding:"required"　gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `binding:"required"　gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
}

// DBの初期化
func dbInit() {
	db := gormConnect()
	// コネクション解放
	defer db.Close()
}

func gormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")
	DBHOST := os.Getenv("DBHOST")
	DBNAME := os.Getenv("DBNAME")

	CONNECT := DBUSER + ":" + DBPASS + "@("+DBHOST+")/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}


func main(){
	db:=gormConnect()
	defer db.Close()
	r:=gin.Default()

	db.Create(&User{Email:"test2@example.com",Password:"testpass"})
	db.Create(&Todo{Content:"このTodoアプリを完成させる",UserId:1})
	var user User
	db.First(&user,3)
	var todo Todo
	db.First(&todo,1)
	r.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"user":user,
			"todo":todo,
		})
	})
	r.Run()
}
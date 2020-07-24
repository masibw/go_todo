package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/masibw/go_todo/pkg/db"
	"net/http"
	"time"
)

type User struct {
	Id        int       ` gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:id"`
	Email     string    `binding:"required" gorm:"unique;not null;column:email" json:"email"`
	Password  string    `binding:"required" gorm:"column:password" json:"password"`
	CreatedAt time.Time ` gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	LastLogin time.Time `  gorm:"column:last_login" sql:"DEFAULT:current_timestamp"`
}
type Todo struct {
	Id        int       ` gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:id"`
	UserId    int       `binding:"required" gorm:"not null;column:user_id" json:"user_id"`
	Content   string    `binding:"required" gorm:"not null;column:content" json:"content"`
	CreatedAt time.Time ` gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time ` gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
}


func main() {
	db := db.GormConnect()
	defer db.Close()
	r := gin.Default()

	//全てのuserを返す
	r.GET("/api/1.0/users",func(c *gin.Context){
		var users []User
		if err := db.Find(&users).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"users":users,
		})
	})

	//特定のuserを返す
	r.GET("/api/1.0/users/:id",func(c *gin.Context){
		id:=c.Param("id")
		var user User
		if err := db.First(&user,id).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"user":user,
		})
	})

	//userの新規作成
	r.POST("/api/1.0/users",func(c *gin.Context){
		var user User
		if err:=c.ShouldBindJSON(&user); err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if err := db.Create(&user).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"response":"Successfully created user",
		})
	})


	//特定のuserを削除
	r.DELETE("/api/1.0/users/:id",func(c *gin.Context){
		id:=c.Param("id")
		if err := db.Where("id = ?",id).Delete(User{}).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"response":"Successfully deleted user",
		})
	})


	//全てのtodoを返す
	r.GET("/api/1.0/todos",func(c *gin.Context){
		var todos []Todo
		if err := db.Find(&todos).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"todos":todos,
		})
	})

	//特定のtodoを返す
	r.GET("/api/1.0/todos/:id",func(c *gin.Context){
		id:=c.Param("id")
		var todo Todo
		if err := db.First(&todo,id).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"todo":todo,
		})
	})

	//todoの新規作成
	r.POST("/api/1.0/todos",func(c *gin.Context){
		var todo Todo
		if err:=c.ShouldBindJSON(&todo); err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if err := db.Create(&todo).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"response":"Successfully created todo",
		})
	})

	//todoのcontentを更新
	r.PUT("/api/1.0/todos/:id",func(c *gin.Context){
		id:=c.Param("id")
		var content map[string]string
		var todo Todo
		if err:=c.ShouldBindJSON(&content); err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		if err := db.First(&todo,id).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if err := db.Model(&todo).Update("Content",content["content"]).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"response":"Successfully updated todo",
		})
	})

	//特定のtodoの削除
	r.DELETE("/api/1.0/todos/:id",func(c *gin.Context){
		id:=c.Param("id")
		if err := db.Where("id = ?",id).Delete(Todo{}).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"response":"Successfully deleted todo",
		})
	})

	r.Run()
}

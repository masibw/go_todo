package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/masibw/go_todo/pkg/db"
	//"github.com/masibw/go_todo/pkg/utility"
	"local.packages/utility"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Id        int       ` gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:id" json:"id"`
	Email     string    `binding:"required" gorm:"unique;not null;column:email" json:"email"`
	Password  string    `binding:"required" gorm:"column:password" json:"password"`
	CreatedAt time.Time ` gorm:"column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
	LastLogin time.Time `  gorm:"column:last_login" sql:"DEFAULT:current_timestamp" json:"last_login"`
}
type Todo struct {
	Id        int       ` gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:id" json:"id"`
	UserId    int       `gorm:"not null;column:user_id" json:"user_id"`
	Content   string    `binding:"required" gorm:"not null;column:content" json:"content"`
	Done bool `gorm:"not null; column:done" json:"done"`
	CreatedAt time.Time ` gorm:"column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
	UpdatedAt time.Time ` gorm:"column:updated_at" sql:"DEFAULT:current_timestamp" json:"updated_at"`
}


func main() {
	db := db.GormConnect()
	defer db.Close()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	r.Use(cors.New(config))


	db.Create(&User{Email:"example.com",Password:"password"})

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
	r.GET("/api/1.0/users/:user_id",func(c *gin.Context){
		userId:=c.Param("user_id")
		var user User
		if err := db.First(&user,userId).Error; err != nil{
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
		var err error
		if err:=c.ShouldBindJSON(&user); err!=nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.Password,err = utility.Encode(user.Password,10)
		if err != nil {
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


	r.POST("/api/1.0/login",func(c *gin.Context){
		var user User
		if err:=c.ShouldBindJSON(&user); err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		var dbUser User
		if err :=db.Where("email = ?",user.Email).Find(&dbUser).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password),[]byte(user.Password)); err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"response":"Successfully logged in",
		})

	})


	//特定のuserを削除
	r.DELETE("/api/1.0/users/:user_id",func(c *gin.Context){
		userId:=c.Param("user_id")
		if err := db.Where("id = ?",userId).Delete(User{}).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"response":"Successfully deleted user",
		})
	})


	//特定のuserの全てのtodoを返す
	r.GET("/api/1.0/users/:user_id/todos",func(c *gin.Context){
		var todos []Todo
		userId:=c.Param("user_id")
		if err := db.Where("user_id = ?",userId).Find(&todos).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"todos":todos,
		})
	})

	//特定のtodoを返す
	r.GET("/api/1.0/users/:user_id/todos/:todo_id",func(c *gin.Context){
		todoId:=c.Param("todo_id")
		var todo Todo
		if err := db.First(&todo,todoId).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"todo":todo,
		})
	})

	//todoの新規作成
	r.POST("/api/1.0/users/:user_id/todos",func(c *gin.Context){
		var todo Todo
		userId:=c.Param("user_id")
		if err:=c.ShouldBindJSON(&todo); err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		todo.UserId,_=strconv.Atoi(userId)
		if err := db.Create(&todo).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"response":"Successfully created todo",
		})
	})

	//todoのcontentを更新
	r.PUT("/api/1.0/users/:user_id/todos/:todo_id",func(c *gin.Context){
		todoId:=c.Param("todo_id")
		var content map[string]string
		var todo Todo
		if err:=c.ShouldBindJSON(&content); err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		if err := db.First(&todo,todoId).Error; err != nil{
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
	r.DELETE("/api/1.0/users/:user_id/todos/:todo_id",func(c *gin.Context){
		todoId:=c.Param("todo_id")
		if err := db.Where("id = ?",todoId).Delete(Todo{}).Error; err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"response":"Successfully deleted todo",
		})
	})

	r.Run()
}

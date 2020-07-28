package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/masibw/go_todo/pkg/db"
	"golang.org/x/crypto/bcrypt"
	//"local.packages/handler"
	"github.com/masibw/go_todo/cmd/go_todo/infrastructure/api/handler"
	"github.com/masibw/go_todo/cmd/go_todo/infrastructure/persistance"
	."github.com/masibw/go_todo/cmd/go_todo/model"
	//. "local.packages/model"
	//"local.packages/persistance"
	"net/http"
)



func main() {
	db := db.GormConnect()
	defer db.Close()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	r.Use(cors.New(config))

	todoRepo := persistance.NewTodoRepository(db)
	userRepo := persistance.NewUserRepository(db)

	tH:=handler.NewTodoHandler(todoRepo)
	uH:=handler.NewUserHandler(userRepo)



	db.Create(&User{Email:"example.com",Password:"password"})

	//全てのuserを返す
	r.GET("/api/1.0/users",uH.FindUsers)

	//特定のuserを返す
	r.GET("/api/1.0/users/:user_id",uH.FindUser)

	//userの新規作成
	r.POST("/api/1.0/users",uH.CreateUser)


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
	r.DELETE("/api/1.0/users/:user_id",uH.DeleteUser)


	//特定のuserの全てのtodoを返す
	r.GET("/api/1.0/users/:user_id/todos",tH.FindTodos)

	//特定のtodoを返す
	r.GET("/api/1.0/users/:user_id/todos/:todo_id",tH.FindTodo)

	//todoの新規作成
	r.POST("/api/1.0/users/:user_id/todos",tH.CreateTodo)

	//todoのcontentを更新
	r.PUT("/api/1.0/users/:user_id/todos/:todo_id",tH.UpdateTodo)

	//特定のtodoの削除
	r.DELETE("/api/1.0/users/:user_id/todos/:todo_id",tH.DeleteTodo)

	r.Run()
}

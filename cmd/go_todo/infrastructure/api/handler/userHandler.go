package handler

import (

	//"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/masibw/go_todo/cmd/go_todo/model"
	//"github.com/masibw/go_todo/cmd/go_todo/usecase/repository"
	"github.com/masibw/go_todo/pkg/utility"
	"github.com/masibw/go_todo/cmd/go_todo/model"
	"github.com/masibw/go_todo/cmd/go_todo/usecase/repository"

	//"local.packages/model"
	//"local.packages/repository"
	"net/http"
)

type userHandler struct{
	userRepository repository.UserRepository
}

type UserHandler interface{
	CreateUser(c *gin.Context)
	FindUsers(c *gin.Context)
	FindUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewUserHandler(ur repository.UserRepository)UserHandler{
	return &userHandler{userRepository: ur}
}

func (uH *userHandler) CreateUser(c *gin.Context){
	user := &model.User{}
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
	if err := uH.userRepository.Create(user); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"response":"Successfully created user",
	})
}


func (uH *userHandler) FindUsers(c *gin.Context){
	 users,err := uH.userRepository.FindAll()
	 if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"users":users,
	})
}
func(uH *userHandler) FindUser(c *gin.Context){
	userId:=c.Param("user_id")
	user, err := uH.userRepository.Find(userId)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"user":user,
	})
}
func(uH *userHandler) DeleteUser(c *gin.Context){
	userId:=c.Param("user_id")
	if err := uH.userRepository.Delete(userId); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"response":"Successfully deleted user",
	})
}
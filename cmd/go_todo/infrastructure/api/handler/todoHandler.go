package handler

import (
	"github.com/gin-gonic/gin"
	//"github.com/masibw/go_todo/cmd/go_todo/model"
	"strconv"
	"local.packages/model"
	//"github.com/masibw/go_todo/cmd/go_todo/usecase/repository"
	"net/http"
	"local.packages/repository"
)

type todoHandler struct{
	todoRepository repository.TodoRepository
}

type TodoHandler interface{
	CreateTodo(c *gin.Context)
	FindTodos(c *gin.Context)
	FindTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
	UpdateTodo(C *gin.Context)
}

func NewTodoHandler(tr repository.TodoRepository)TodoHandler {
	return &todoHandler{todoRepository: tr}
}

func(tH *todoHandler)CreateTodo(c *gin.Context){
	 todo :=&model.Todo{}
	userId:=c.Param("user_id")
	if err:=c.ShouldBindJSON(&todo); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	todo.UserId,_=strconv.Atoi(userId)
	if err := tH.todoRepository.Create(todo); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"response":"Successfully created todo",
	})
}


func(tH *todoHandler)FindTodos(c *gin.Context){
	userId:=c.Param("user_id")
	todos, err := tH.todoRepository.FindAll(userId)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"todos":todos,
	})
}

func(tH *todoHandler)FindTodo(c *gin.Context){
	todoId:=c.Param("todo_id")
	todo,err := tH.todoRepository.Find(todoId)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"todo":todo,
	})
}

func(tH *todoHandler)DeleteTodo(c *gin.Context){
	todoId:=c.Param("todo_id")
	if err := tH.todoRepository.Delete(todoId); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"response":"Successfully deleted todo",
	})
}

func(tH *todoHandler)UpdateTodo(c *gin.Context){
	todoId:=c.Param("todo_id")
	var content map[string]string
	if err:=c.ShouldBindJSON(&content); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err := tH.todoRepository.Update(todoId,content["content"]); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"response":"Successfully updated todo",
	})
}
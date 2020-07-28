package repository

import (
	//"github.com/masibw/go_todo/cmd/go_todo/model"
	"local.packages/model"
)

type TodoRepository interface {
	Create(todo *model.Todo) error
	FindAll(userId string) ([]*model.Todo, error)
	Find(todoId string) (*model.Todo, error)
	Delete(todoId string) error
	Update(todoId ,content string)error
}
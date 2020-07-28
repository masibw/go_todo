package repository

import (
	"github.com/masibw/go_todo/cmd/go_todo/model"
	//"local.packages/model"
)

type UserRepository interface{
	Create(user *model.User)error
 	FindAll()([]*model.User, error)
 	Find(userId string)(*model.User, error)
	Delete(userId string)error
}

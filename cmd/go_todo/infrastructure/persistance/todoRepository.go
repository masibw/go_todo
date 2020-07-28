package persistance


import(
	//"github.com/masibw/go_todo/cmd/go_todo/model"
	//"github.com/masibw/go_todo/cmd/go_todo/model"
	//"local.packages/handler"
	//"github.com/masibw/go_todo/cmd/go_todo/infrastructure/api/handler"
	//"github.com/masibw/go_todo/cmd/go_todo/infrastructure/persistance"
	//."github.com/masibw/go_todo/cmd/go_todo/model"
	// "local.packages/model"
	"github.com/jinzhu/gorm"
	"github.com/masibw/go_todo/cmd/go_todo/model"
	"github.com/masibw/go_todo/cmd/go_todo/usecase/repository"

	//"github.com/masibw/go_todo/cmd/go_todo/usecase/repository"
	//"local.packages/repository"
)

type todoRepository struct{
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository{
	return &todoRepository{db}
}

func (tr *todoRepository)Create(todo *model.Todo)error{
	if err := tr.db.Create(&todo).Error; err != nil{
		return err
	}
	return nil
}

func (tr *todoRepository)FindAll(userId string)([]*model.Todo, error){
	todos :=[]*model.Todo{}
	if err := tr. db.Where("user_id = ?",userId).Find(&todos).Error; err != nil{
		return nil,err
	}
	return todos,nil
}

func (tr *todoRepository)Find(todoId string)(*model.Todo, error){
	todo :=model.Todo{}
	if err := tr.db.First(&todo,todoId).Error; err != nil{
		return nil,err
	}
	return &todo,nil
}

func (tr *todoRepository)Delete(todoId string)error{
	if err := tr.db.Where("id = ?",todoId).Delete(model.Todo{}).Error; err != nil{
		return err
	}
	return nil
}

func (tr *todoRepository)Update(todoId , content string)error{
	todo:= &model.Todo{}
	var err error
	if todo,err = tr.Find(todoId); err != nil{
		return err
	}
	if err := tr.db.Model(todo).Update("Content",content).Error; err != nil{
		return err
	}
	return nil
}

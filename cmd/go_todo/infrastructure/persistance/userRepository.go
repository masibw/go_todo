package persistance


import(
	"fmt"
	"github.com/masibw/go_todo/cmd/go_todo/model"

	//"github.com/masibw/go_todo/cmd/go_todo/usecase/repository"
	//"github.com/masibw/go_todo/cmd/go_todo/model"
	//"local.packages/model"
	//"local.packages/repository"
	"github.com/jinzhu/gorm"
	"github.com/masibw/go_todo/cmd/go_todo/usecase/repository"
)

type userRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository{
	return &userRepository{db}
}

func (ur *userRepository)Create(user *model.User)error{
	if user == nil{
		fmt.Print("user nil")
	}
	if err := ur.db.Create(&user).Error; err != nil{
		return err
	}
	return nil
}

func (ur *userRepository)FindAll()([]*model.User,error){
	users :=  []*model.User{}
	if err := ur.db.Find(&users).Error; err != nil{
		return nil,err
	}
	return users,nil
}

func(ur *userRepository)Find(userId string)(*model.User,error){
	user := model.User{}
	if err := ur.db.First(&user,userId).Error; err != nil{
		return nil,err
	}
	return &user,nil
}

func (ur *userRepository)Delete(userId string)error{
	if err := ur.db.Where("id = ?",userId).Delete(model.User{}).Error; err != nil{
		return err
	}
	return nil
}

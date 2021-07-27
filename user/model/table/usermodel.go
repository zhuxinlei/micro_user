package table

import (
	"github.com/jinzhu/gorm"
	"github.com/zhuxinlei/micro_user/user/model/input"
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) UserModel {
	return UserModel{
		db: db,
	}
}

func (u UserModel) Insert(input input.User) (err error) {
	err = u.db.Debug().Create(&input).Error
	return

}
func (u UserModel) GetUser(user input.User) (input input.User, err error) {
	err = u.db.Debug().Where(user).Take(&input).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

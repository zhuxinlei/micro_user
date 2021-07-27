package model

import (
	"github.com/zhuxinlei/micro_user/user/model/input"
)

type User interface {
	Insert(user input.User) error
	GetUser(user input.User) (input.User, error)
}

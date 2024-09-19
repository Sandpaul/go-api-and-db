package user

import (
	"acme/model"
)

type Repository interface {
	GetUsers() ([]model.User, error)
	GetUser(id int) (model.User, error)
	AddUser(user model.User) (id int, err error)
	UpdateUserName(id int, user *model.User) (model.User, error)
	DeleteUser(id int) error
	Close()
}

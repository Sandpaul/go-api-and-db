package db

import (
	"acme/model"
	"errors"
	"slices"
)

var users []model.User
var count int = 3

func init() {
	ResetUsers()
}

func ResetUsers() {
	users = []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}
}

func GetUsers() ([]model.User, error) {
	return users, nil
}

func GetUser(id int) (model.User, error) {
	var user model.User

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return user, errors.New("user id not found")
}

func AddUser(user model.User) (id int, err error) {
	count++
	user.ID = count

	users = append(users, user)

	return count, nil
}

func DeleteUser(id int) error {

	for index, user := range users {
		if user.ID == id {
			users = slices.Delete(users, index, index+1)
			return nil
		}
	}

	return errors.New("user id not found to delete")
}

func UpdateUserName(id int, updatedUser model.User) (model.User, error) {
	for index, user := range users {
		if user.ID == id {
			users[index].Name = updatedUser.Name
			return user, nil
		}
	}

	return model.User{}, errors.New("user id not found to update")
}

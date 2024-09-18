package service

import (
	"acme/db"
	"acme/model"
	"acme/postgres"
	"errors"
	"fmt"
)

func GetUsers() ([]model.User, error) {
	users, err := postgres.GetUsers()

	if err != nil {
		fmt.Println("Error getting users from DB:", err)
		return nil, errors.New("there was an error getting the users from the database")
	}

	return users, nil
}

func DeleteUser(id int) error {
	err := postgres.DeleteUser(id)

	if err != nil {
		fmt.Println("Error deleting user from DB:", err)
		return errors.New("could not delete user")
	}

	return nil
}

func GetSingleUser(id int) (model.User, error) {
	user, err := postgres.GetUser(id)

	if err != nil {
		fmt.Println("Error retrieving user from DB:", err)
		return model.User{}, errors.New("could not retrieve user")
	}

	return user, nil
}

func CreateUser(user model.User) (id int, err error) {
	id, err = postgres.AddUser(user)

	if err != nil {
		fmt.Println("Error adding user to DB:", err)
		return 0, errors.New("could not create user")
	}

	return id, nil
}

func UpdateUserName(id int, user model.User) (model.User, error) {
	user, err := db.UpdateUserName(id, user)

	if err != nil {
		fmt.Println("Error updating user name in DB:", err)
		return model.User{}, errors.New("could not update user name")
	}

	return user, nil
}

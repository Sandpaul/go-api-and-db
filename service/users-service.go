package service

import (
	"acme/model"
	"acme/repository/user"
	"errors"
	"fmt"
)

type UserService struct {
	repository user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	users, err := s.repository.GetUsers()

	if err != nil {
		fmt.Println("Error getting users from DB:", err)
		return nil, errors.New("there was an error getting the users from the database")
	}

	return users, nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.repository.DeleteUser(id)

	if err != nil {
		fmt.Println("Error deleting user from DB:", err)
		return errors.New("could not delete user")
	}

	return nil
}

func (s *UserService) GetSingleUser(id int) (model.User, error) {
	user, err := s.repository.GetUser(id)

	if err != nil {
		fmt.Println("Error retrieving user from DB:", err)
		return model.User{}, errors.New("could not retrieve user")
	}

	return user, nil
}

func (s *UserService) CreateUser(user model.User) (id int, err error) {
	id, err = s.repository.AddUser(user)

	if err != nil {
		fmt.Println("Error adding user to DB:", err)
		return 0, errors.New("could not create user")
	}

	return id, nil
}

func (s *UserService) UpdateUserName(id int, user model.User) (model.User, error) {
	updatedUser, err := s.repository.UpdateUserName(id, &user)

	if err != nil {
		fmt.Println("Error updating user name in DB:", err)
		return model.User{}, errors.New("could not update user name")
	}

	return updatedUser, nil
}

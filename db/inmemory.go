package db

import (
	"slices"
)

type User struct {
	ID   int    `json:"id`
	Name string `json:"name"`
}

var users []User
var count int = 3

func init() {
	ResetUsers()
}

func ResetUsers() {
	users = []User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}
}

func GetUsers() []User {
	return users
}

func GetUser(id int) User {
	var user User

	for _, user := range users{
		if user.ID == id {
			return user
		}
	}

	return user
} 

func AddUser(user User) (id int) {
	count++
	user.ID = count

	users = append(users, user)

	return count
}

func DeleteUser(id int) {
	index := -1

	for i, user := range users {
		if user.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		users = slices.Delete(users, index, index+1)
	}
}

package db

type User struct {
	ID   int    `json:"id`
	Name string `json:"name"`
}

var users []User

func init() {
	users = []User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}
}

func GetUsers() []User {
	return users
}

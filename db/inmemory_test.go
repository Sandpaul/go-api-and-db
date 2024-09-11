package db

import (
	"acme/model"
	"slices"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	ResetUsers()

	DeleteUser(2)

	expectedUsers := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 3, Name: "User 3"},
	}

	if !slices.Equal(users, expectedUsers) {
		t.Errorf("Expected users %v, but got %v", expectedUsers, users)
	}
}

func TestDeleteUserInvalidId(t *testing.T) {
	ResetUsers()

	DeleteUser(7)

	expectedUsers := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	if !slices.Equal(users, expectedUsers) {
		t.Errorf("Expected users %v, but got %v", expectedUsers, users)
	}
}

func TestUpdateUserName(t *testing.T) {
	ResetUsers()

	user := model.User{
		Name: "Ralph",
	}

	UpdateUserName(1, user)

	expectedUsers := []model.User{
		{ID: 1, Name: "Ralph"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	if !slices.Equal(users, expectedUsers) {
		t.Errorf("Expected users %v, but got %v", expectedUsers, users)
	}
}

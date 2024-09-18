package db

import (
	"acme/model"
	"slices"
	"testing"
)

func TestGetUsers(t *testing.T) {
	ResetUsers()

	expected_users := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	actual_users, err := GetUsers()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !slices.Equal(expected_users, actual_users) {
		t.Errorf("Expected users: %v but got: %v", expected_users, actual_users)
	}
}

func TestGetUser(t *testing.T) {
	ResetUsers()

	expectedUser := model.User{ID: 2, Name: "User 2"}

	user, err := GetUser(2)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user != expectedUser {
		t.Errorf("expected %v, got %v", expectedUser, user)
	}

	_, err = GetUser(999)

	if err == nil {
		t.Errorf("expected an error for non-existing user ID, got nil")
	}
}

func TestAddUserAddsNewUser(t *testing.T) {
	ResetUsers()

	expectedUsers := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
		{ID: 4, Name: "User 4"},
	}

	newUser := model.User{Name: "User 4"}

	newUserID, err := AddUser(newUser)

	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	expectedUserID := 4

	if newUserID != expectedUserID {
		t.Errorf("expected new user ID to be %v, but got %v", expectedUserID, newUserID)
	}

	users, err := GetUsers()

	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if !slices.Equal(users, expectedUsers) {
		t.Errorf("expected users %v, but got %v", expectedUsers, users)
	}
}

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

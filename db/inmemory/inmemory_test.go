package inmemory

import (
	"acme/model"
	"acme/repository/user"
	"reflect"
	"testing"
)

func TestGetUsers(t *testing.T) {

	repo := user.NewInMemoryUserRepository()

	expected_users := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	actual_users, err := repo.GetUsers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(expected_users, actual_users) {
		t.Errorf("Expected users: %v but got: %v", expected_users, actual_users)
	}
}

func TestGetUser(t *testing.T) {

	repo := user.NewInMemoryUserRepository()

	expectedUser := model.User{ID: 2, Name: "User 2"}

	user, err := repo.GetUser(2)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user != expectedUser {
		t.Errorf("expected %v, got %v", expectedUser, user)
	}

	_, err = repo.GetUser(999)

	if err == nil {
		t.Errorf("expected an error for non-existing user ID, got nil")
	}
}

func TestAddUserAddsNewUser(t *testing.T) {

	repo := user.NewInMemoryUserRepository()

	expectedUsers := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
		{ID: 4, Name: "User 4"},
	}

	newUser := model.User{Name: "User 4"}

	newUserID, err := repo.AddUser(newUser)

	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	expectedUserID := 4

	if newUserID != expectedUserID {
		t.Errorf("expected new user ID to be %v, but got %v", expectedUserID, newUserID)
	}

	users, err := repo.GetUsers()

	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("expected users %v, but got %v", expectedUsers, users)
	}
}

func TestDeleteUser(t *testing.T) {

	repo := user.NewInMemoryUserRepository()

	repo.DeleteUser(2)

	expectedUsers := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 3, Name: "User 3"},
	}

	users, err := repo.GetUsers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("Expected users %v, but got %v", expectedUsers, users)
	}
}

func TestDeleteUserInvalidId(t *testing.T) {

	repo := user.NewInMemoryUserRepository()
	
	err := repo.DeleteUser(7)
	
	if err == nil {
		t.Errorf("Expected an error when deleting user with invalid ID, but got nil")
	}

	expectedUsers := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	users, err := repo.GetUsers()
	if err != nil {
		t.Fatalf("error getting users")
	}

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("Expected users %v, but got %v", expectedUsers, users)
	}
}

func TestUpdateUserName(t *testing.T) {

	repo := user.NewInMemoryUserRepository()

	user := model.User{
		Name: "Ralph",
	}

	repo.UpdateUserName(1, &user)

	expectedUsers := []model.User{
		{ID: 1, Name: "Ralph"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	users, err := repo.GetUsers()
	if err != nil {
		t.Fatalf("error getting users")
	}

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("Expected users %v, but got %v", expectedUsers, users)
	}
}

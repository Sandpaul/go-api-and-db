package db

import (
	"slices"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	ResetUsers()

	result := DeleteUser(2)

	expectedUsers := []User{
			{ID: 1, Name: "User 1"},
			{ID: 3, Name: "User 3"},
		}

	if !slices.Equal(users, expectedUsers) {
		t.Errorf("Expected users %v, but got %v", expectedUsers, users)
	}

	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestDeleteUserInvalidId(t *testing.T) {
	ResetUsers()

	result := DeleteUser(7)

	expectedUsers := []User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	if !slices.Equal(users, expectedUsers) {
		t.Errorf("Expected users %v, but got %v", expectedUsers, users)
	}

	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
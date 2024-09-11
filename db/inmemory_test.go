package db

import (
	"slices"
	"testing"
)

func TestDeleteUser(t *testing.T) {

	DeleteUser(2)

	expectedUsers := []User{
			{ID: 1, Name: "User 1"},
			{ID: 3, Name: "User 3"},
		}

	if !slices.Equal(users, expectedUsers) {
		t.Errorf("Expected users %v, but got %v", expectedUsers, users)
	}
}
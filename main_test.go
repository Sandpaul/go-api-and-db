package main

import (
	"acme/api"
	"acme/model"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(rootHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello, World!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestRootHandlerWithServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer server.Close()

	resp, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello, World!"
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if string(bodyBytes) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(bodyBytes), expected)
	}
}

func TestGetUsersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.GetUsers)

	expected := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual []model.User
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetUsersHandlerWithServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(api.GetUsers))
	defer server.Close()

	resp, err := http.Get(server.URL + "/api/users")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var actual []model.User
	if err := json.Unmarshal(bodyBytes, &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetSingleUserHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := http.NewServeMux()
	router.HandleFunc("GET /api/users/{id}", api.GetSingleUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}
}

func TestGetSingleUserNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := http.NewServeMux()
	router.HandleFunc("GET /api/users/{id}", api.GetSingleUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestCreateUserHandler(t *testing.T) {
	userData := `{"Name": "New User"}`
	req, err := http.NewRequest("POST", "/api/users", io.NopCloser(strings.NewReader(userData)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.CreateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	expected := "User created successfully: 4"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, wanted %v", rr.Body.String(), expected)
	}
}

func TestDeleteUserHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := http.NewServeMux()
	router.HandleFunc("DELETE /api/users/{id}", api.DeleteUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteUserHandlerNotFound(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/users/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := http.NewServeMux()
	router.HandleFunc("DELETE /api/users/{id}", api.DeleteUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

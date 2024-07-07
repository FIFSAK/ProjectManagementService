package handlers

import (
	"ProjectManagementService/internal/models"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllUsersHandler(t *testing.T) {
	mockUserModel := &models.MockUserModel{
		MockGetUsers: func() ([]*models.User, error) {
			return []*models.User{
				{ID: 1, Name: "Test User", Email: "test@example.com", Role: "admin"},
			}, nil
		},
	}

	handler := NewUserHandler(mockUserModel)
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.GetAllUsersHandler).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":1,"name":"Test User","email":"test@example.com","registration_date":"","role":"admin"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateUserHandler(t *testing.T) {
	mockUserModel := &models.MockUserModel{
		MockCreateUser: func(name string, email string, role string) error {
			return nil
		},
	}

	handler := NewUserHandler(mockUserModel)
	userJson := `{"name":"Test User","email":"test@example.com","role":"admin"}`
	req, err := http.NewRequest("POST", "/users", strings.NewReader(userJson))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.CreateUserHandler).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestGetUserHandler(t *testing.T) {
	mockUserModel := &models.MockUserModel{
		MockGetUserById: func(id int) (*models.User, error) {
			return &models.User{ID: 1, Name: "Test User", Email: "test@example.com", Role: "admin"}, nil
		},
	}

	handler := NewUserHandler(mockUserModel)
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handler.GetUserHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":1,"name":"Test User","email":"test@example.com","registration_date":"","role":"admin"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
func TestUpdateUserHandler(t *testing.T) {
	mockUserModel := &models.MockUserModel{
		MockGetUserById: func(id int) (*models.User, error) {
			return &models.User{ID: 1, Name: "Old Name", Email: "old@example.com", Role: "user"}, nil
		},
		MockUpdateUser: func(id int, name string, email string, role string) error {
			if id != 1 || name != "New Name" || email != "new@example.com" || role != "admin" {
				t.Errorf("Unexpected input: %v, %v, %v, %v", id, name, email, role)
			}
			return nil
		},
	}

	handler := NewUserHandler(mockUserModel)
	userJson := `{"name":"New Name","email":"new@example.com","role":"admin"}`
	req, err := http.NewRequest("PUT", "/users/1", strings.NewReader(userJson))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handler.UpdateUserHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteUserHandler(t *testing.T) {
	mockUserModel := &models.MockUserModel{
		MockDeleteUser: func(id int) (int, error) {
			if id != 1 {
				t.Errorf("Unexpected id: %v", id)
			}
			return 1, nil
		},
	}

	handler := NewUserHandler(mockUserModel)
	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handler.DeleteUserHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

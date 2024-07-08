package handlers

import (
	"ProjectManagementService/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserModel models.UserModel
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewUserHandler(userModel models.UserModel) *UserHandler {
	return &UserHandler{
		UserModel: userModel,
	}
}

// @Summary Health check
// @Tags health check\

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// @Summary Get all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
// @Failure 404 {string} string "No users found"
// @Failure 500 {string} string "Internal server error"
func (uh *UserHandler) GetAllUsersHandler(writer http.ResponseWriter, request *http.Request) {
	users, err := uh.UserModel.GetUsers()
	if len(users) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonUsers)
	if err != nil {
		return
	}
}

// @Summary Create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserInput true "User object"
// @Success 201 {string} string "User created"
// @Router /users [post]
// @Failure 400 {string} string "Missing required fields"
// @Failure 500 {string} string "Internal server error"
func (uh *UserHandler) CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if user.Role == "" || user.Name == "" || user.Email == "" {
		http.Error(writer, "missing required fields", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = uh.UserModel.CreateUser(user.Name, user.Email, user.Role)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

// @Summary Get user by id
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
func (uh *UserHandler) GetUserHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := uh.UserModel.GetUserById(id)
	if user == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonUser)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body UserInput true "User object"
// @Success 200 {string} string "User updated"
// @Router /users/{id} [put]
// @Failure 400 {string} string "Missing required fields"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
func (uh *UserHandler) UpdateUserHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := uh.UserModel.GetUserById(id)
	if user == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = uh.UserModel.UpdateUser(user.ID, user.Name, user.Email, user.Role)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

// @Summary Delete user
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {string} string "User deleted"
// @Router /users/{id} [delete]
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
func (uh *UserHandler) DeleteUserHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	deleteId, err := uh.UserModel.DeleteUser(id)
	if deleteId == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

// @Summary Get user tasks
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} models.Task
// @Router /users/{id}/tasks [get]
// @Failure 404 {string} string "No tasks found"
// @Failure 500 {string} string "Internal server error"
func (uh *UserHandler) GetUserTasksHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	user_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	tasks, err := uh.UserModel.GetUserTasks(user_id)
	if len(tasks) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonTasks)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Search user
// @Tags users
// @Produce json
// @Param email query string false "User email"
// @Param name query string false "User name"
// @Success 200 {array} models.User
// @Router /users/search [get]
// @Failure 400 {string} string "Missing email or name parameter"
// @Failure 404 {string} string "No users found"
// @Failure 500 {string} string "Internal server error"
func (uh *UserHandler) SearchUserHandler(writer http.ResponseWriter, request *http.Request) {
	email := request.URL.Query().Get("email")
	name := request.URL.Query().Get("name")
	if email == "" && name == "" {
		http.Error(writer, "missing email or name parameter", http.StatusBadRequest)
		return
	}
	var (
		users []*models.User
		err   error
	)
	if email != "" {
		users, err = uh.UserModel.SearchUserByEmail(email)

	} else {
		users, err = uh.UserModel.SearchUserByName(name)
	}
	if len(users) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonUsers, err := json.Marshal(users)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonUsers)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

}

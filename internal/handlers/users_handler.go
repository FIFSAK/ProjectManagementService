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

func NewUserHandler(userModel models.UserModel) *UserHandler {
	return &UserHandler{
		UserModel: userModel,
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

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

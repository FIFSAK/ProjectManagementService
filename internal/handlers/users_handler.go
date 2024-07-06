package handlers

import (
	"ProjectManagementService/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetAllUsersHandler(userModel *models.UserModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		users, err := userModel.GetUsers()
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
}

func CreateUserHandler(userModel *models.UserModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var user models.User
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		err = userModel.CreateUser(user.Name, user.Email, user.Role)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	}
}

func GetUserHandler(userModel *models.UserModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := userModel.GetUserById(id)
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
}

func UpdateUserHandler(userModel *models.UserModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := userModel.GetUserById(id)
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
		err = userModel.UpdateUser(user.ID, user.Name, user.Email, user.Role)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func DeleteUserHandler(userModel *models.UserModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		deleteId, err := userModel.DeleteUser(id)
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

}

func GetUserTasksHandler(userModel *models.UserModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		user_id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		tasks, err := userModel.GetUserTasks(user_id)
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
}

func SearchUserHandler(userModel *models.UserModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
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
			users, err = userModel.SearchUserByEmail(email)

		} else {
			users, err = userModel.SearchUserByName(name)
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
}

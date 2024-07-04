package handlers

import (
	"ProjectManagementService/internal/models"
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetAllUsersHandler(userModel *models.UserModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		users, err := userModel.GetUsers()
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
		_, err = writer.Write(jsonUsers)
		if err != nil {
			return
		}

	}
}

func CreateUserHandler(writer http.ResponseWriter, request *http.Request) {

}

func GetUserHandler(writer http.ResponseWriter, request *http.Request) {

}

func UpdateUserHandler(writer http.ResponseWriter, request *http.Request) {

}

func DeleteUserHandler(writer http.ResponseWriter, request *http.Request) {

}

func GetUserTasksHandler(writer http.ResponseWriter, request *http.Request) {

}

func SearchUserHandler(writer http.ResponseWriter, request *http.Request) {

}

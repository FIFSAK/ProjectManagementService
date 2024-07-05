package main

import (
	"ProjectManagementService/internal/handlers"
	"ProjectManagementService/internal/models"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRouter(router *mux.Router, userModel *models.UserModel, taskModel *models.TaskModel, projectModel *models.ProjectModel) {
	router.HandleFunc("/health-check", handlers.HealthCheck).Methods(http.MethodGet)

	usersRouter := router.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("", handlers.GetAllUsersHandler(userModel)).Methods(http.MethodGet)
	usersRouter.HandleFunc("", handlers.CreateUserHandler(userModel)).Methods(http.MethodPost)
	usersRouter.HandleFunc("/{id:[0-9]+}", handlers.GetUserHandler(userModel)).Methods(http.MethodGet)
	usersRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateUserHandler(userModel)).Methods(http.MethodPut)
	usersRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteUserHandler(userModel)).Methods(http.MethodDelete)
	usersRouter.HandleFunc("/{id:[0-9]+}/tasks", handlers.GetUserTasksHandler(userModel)).Methods(http.MethodGet)
	usersRouter.HandleFunc("/search", handlers.SearchUserHandler(userModel)).Methods(http.MethodGet)

	tasksRouter := router.PathPrefix("/tasks").Subrouter()

	tasksRouter.HandleFunc("", handlers.GetAllTasksHandler(taskModel)).Methods(http.MethodGet)
	tasksRouter.HandleFunc("", handlers.CreateTaskHandler(taskModel)).Methods(http.MethodPost)
	tasksRouter.HandleFunc("/{id:[0-9]+}", handlers.GetTaskHandler(taskModel)).Methods(http.MethodGet)
	tasksRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateTaskHandler(taskModel)).Methods(http.MethodPut)
	tasksRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteTaskHandler(taskModel)).Methods(http.MethodDelete)
	tasksRouter.HandleFunc("/search", handlers.SearchTasksHandler(taskModel)).Methods(http.MethodGet)

	projectsRouter := router.PathPrefix("/projects").Subrouter()

	projectsRouter.HandleFunc("", handlers.GetAllProjectsHandler(projectModel)).Methods(http.MethodGet)
	projectsRouter.HandleFunc("", handlers.CreateProjectHandler(projectModel)).Methods(http.MethodPost)
	projectsRouter.HandleFunc("/{id:[0-9]+}", handlers.GetProjectHandler(projectModel)).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateProjectHandler(projectModel)).Methods(http.MethodPut)
	projectsRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteProjectHandler(projectModel)).Methods(http.MethodDelete)
	projectsRouter.HandleFunc("/{id:[0-9]+}/tasks", handlers.GetProjectTasksHandler(projectModel)).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/search", handlers.SearchProjectsHandler(projectModel)).Methods(http.MethodGet)
}

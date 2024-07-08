package main

import (
	_ "ProjectManagementService/docs"
	"ProjectManagementService/internal/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func SetupRouter(router *mux.Router, userHandler *handlers.UserHandler, taskHandler *handlers.TaskHandler, projectHandler *handlers.ProjectHandler) {
	router.HandleFunc("/health-check", handlers.HealthCheck).Methods(http.MethodGet)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	usersRouter := router.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("", userHandler.GetAllUsersHandler).Methods(http.MethodGet)
	usersRouter.HandleFunc("", userHandler.CreateUserHandler).Methods(http.MethodPost)
	usersRouter.HandleFunc("/{id:[0-9]+}", userHandler.GetUserHandler).Methods(http.MethodGet)
	usersRouter.HandleFunc("/{id:[0-9]+}", userHandler.UpdateUserHandler).Methods(http.MethodPut)
	usersRouter.HandleFunc("/{id:[0-9]+}", userHandler.DeleteUserHandler).Methods(http.MethodDelete)
	usersRouter.HandleFunc("/{id:[0-9]+}/tasks", userHandler.GetUserTasksHandler).Methods(http.MethodGet)
	usersRouter.HandleFunc("/search", userHandler.SearchUserHandler).Methods(http.MethodGet)

	tasksRouter := router.PathPrefix("/tasks").Subrouter()

	tasksRouter.HandleFunc("", taskHandler.GetAllTasksHandler).Methods(http.MethodGet)
	tasksRouter.HandleFunc("", taskHandler.CreateTaskHandler).Methods(http.MethodPost)
	tasksRouter.HandleFunc("/{id:[0-9]+}", taskHandler.GetTaskHandler).Methods(http.MethodGet)
	tasksRouter.HandleFunc("/{id:[0-9]+}", taskHandler.UpdateTaskHandler).Methods(http.MethodPut)
	tasksRouter.HandleFunc("/{id:[0-9]+}", taskHandler.DeleteTaskHandler).Methods(http.MethodDelete)
	tasksRouter.HandleFunc("/search", taskHandler.SearchTasksHandler).Methods(http.MethodGet)

	projectsRouter := router.PathPrefix("/projects").Subrouter()

	projectsRouter.HandleFunc("", projectHandler.GetAllProjectsHandler).Methods(http.MethodGet)
	projectsRouter.HandleFunc("", projectHandler.CreateProjectHandler).Methods(http.MethodPost)
	projectsRouter.HandleFunc("/{id:[0-9]+}", projectHandler.GetProjectHandler).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/{id:[0-9]+}", projectHandler.UpdateProjectHandler).Methods(http.MethodPut)
	projectsRouter.HandleFunc("/{id:[0-9]+}", projectHandler.DeleteProjectHandler).Methods(http.MethodDelete)
	projectsRouter.HandleFunc("/{id:[0-9]+}/tasks", projectHandler.GetProjectTasksHandler).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/search", projectHandler.SearchProjectsHandler).Methods(http.MethodGet)
}

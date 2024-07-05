package handlers

import (
	"ProjectManagementService/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllProjectsHandler(projectModel *models.ProjectModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		projects, err := projectModel.GetProjects()
		if err != nil {
			http.Error(writer, "could not get projects: "+err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(projects)
		if err != nil {
			http.Error(writer, "could not encode projects: "+err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func CreateProjectHandler(projectModel *models.ProjectModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var project models.Project
		err := json.NewDecoder(request.Body).Decode(&project)
		if err != nil {
			http.Error(writer, "could not decode project: "+err.Error(), http.StatusBadRequest)
			return
		}
		err = projectModel.CreateProject(project.Title, project.Description, project.ManagerID)
		if err != nil {
			http.Error(writer, "could not create project: "+err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	}

}

func GetProjectHandler(projectModel *models.ProjectModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		project, err := projectModel.GetProjectByID(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(project)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func UpdateProjectHandler(projectModel *models.ProjectModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		project, err := projectModel.GetProjectByID(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewDecoder(request.Body).Decode(&project)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		err = projectModel.UpdateProject(project.ID, project.Title, project.Description, project.ManagerID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

}

func DeleteProjectHandler(projectModel *models.ProjectModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		err = projectModel.DeleteProject(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

}

func GetProjectTasksHandler(projectModel *models.ProjectModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		tasks, err := projectModel.GetProjectTasks(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(tasks)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func SearchProjectsHandler(projectModel *models.ProjectModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		title := request.URL.Query().Get("title")
		manager := request.URL.Query().Get("manager")
		if title != "" {
			projects, err := projectModel.SearchProjectsByTitle(title)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			writer.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(writer).Encode(projects)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if manager != "" {
			managerID, err := strconv.Atoi(manager)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			writer.Header().Set("Content-Type", "application/json")
			projects, err := projectModel.SearchProjectsByManagerID(managerID)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(writer).Encode(projects)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(writer, "invalid search parameters", http.StatusBadRequest)
			return
		}
	}

}

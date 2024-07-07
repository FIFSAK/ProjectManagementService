package handlers

import (
	"ProjectManagementService/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProjectHandler struct {
	ProjectModel models.ProjectModel
}

func NewProjectHandler(projectModel models.ProjectModel) *ProjectHandler {
	return &ProjectHandler{
		ProjectModel: projectModel,
	}

}

func (ph *ProjectHandler) GetAllProjectsHandler(writer http.ResponseWriter, request *http.Request) {
	projects, err := ph.ProjectModel.GetProjects()
	if err != nil {
		http.Error(writer, "could not get projects: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if len(projects) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(projects)
	if err != nil {
		http.Error(writer, "could not encode projects: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func (ph *ProjectHandler) CreateProjectHandler(writer http.ResponseWriter, request *http.Request) {
	var project models.Project
	err := json.NewDecoder(request.Body).Decode(&project)
	if err != nil {
		http.Error(writer, "could not decode project: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = ph.ProjectModel.CreateProject(project.Title, project.Description, project.ManagerID)
	if err != nil {
		http.Error(writer, "could not create project: "+err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (ph *ProjectHandler) GetProjectHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	project, err := ph.ProjectModel.GetProjectByID(id)
	if project == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(project)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ph *ProjectHandler) UpdateProjectHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	project, err := ph.ProjectModel.GetProjectByID(id)
	if project == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&project)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = ph.ProjectModel.UpdateProject(project.ID, project.Title, project.Description, project.ManagerID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (ph *ProjectHandler) DeleteProjectHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	deletedId, err := ph.ProjectModel.DeleteProject(id)
	if deletedId == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (ph *ProjectHandler) GetProjectTasksHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	tasks, err := ph.ProjectModel.GetProjectTasks(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(tasks) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(tasks)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ph *ProjectHandler) SearchProjectsHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Query().Get("title")
	manager := request.URL.Query().Get("manager")
	var (
		projects []models.Project
		err      error
	)
	if title != "" {
		projects, err = ph.ProjectModel.SearchProjectsByTitle(title)
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
		projects, err = ph.ProjectModel.SearchProjectsByManagerID(managerID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(writer, "invalid search parameters", http.StatusBadRequest)
		return
	}
	if len(projects) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(projects)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

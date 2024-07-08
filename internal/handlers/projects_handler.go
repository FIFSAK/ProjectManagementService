package handlers

import (
	"ProjectManagementService/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProjectInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ManagerID   int    `json:"manager_id"`
}

type ProjectHandler struct {
	ProjectModel models.ProjectModel
}

func NewProjectHandler(projectModel models.ProjectModel) *ProjectHandler {
	return &ProjectHandler{
		ProjectModel: projectModel,
	}

}

// @Summary Get all projects
// @Tags projects
// @Produce json
// @Success 200 {array} models.Project
// @Router /projects [get]
// @Failure 404 {string} string "No projects found"
// @Failure 500 {string} string "Internal server error"
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

// @Summary Create a project
// @Tags projects
// @Accept json
// @Produce json
// @Param project body ProjectInput true "Project information"
// @Success 201 {string} string "Project created"
// @Router /projects [post]
// @Failure 400 {string} string "Could not decode project"
// @Failure 500 {string} string "Internal server error"
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

// @Summary Get project by ID
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} models.Project
// @Router /projects/{id} [get]
// @Failure 404 {string} string "Project not found"
// @Failure 500 {string} string "Internal server error"
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

// @Summary Update a project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param project body ProjectInput true "Project information"
// @Success 200 {string} string "Project updated"
// @Router /projects/{id} [put]
// @Failure 400 {string} string "Could not decode project"
// @Failure 404 {string} string "Project not found"
// @Failure 500 {string} string "Internal server error"
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

// @Summary Delete a project
// @Tags projects
// @Param id path int true "Project ID"
// @Success 200 {string} string "Project deleted"
// @Router /projects/{id} [delete]
// @Failure 404 {string} string "Project not found"
// @Failure 500 {string} string "Internal server error"
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

// @Summary Get all tasks for a project
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {array} models.Task
// @Router /projects/{id}/tasks [get]
// @Failure 404 {string} string "No tasks found"
// @Failure 500 {string} string "Internal server error"
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

// @Summary Search projects
// @Tags projects
// @Produce json
// @Param title query string false "Project title"
// @Param manager query string false "Project manager ID"
// @Success 200 {array} models.Project
// @Router /projects/search [get]
// @Failure 400 {string} string "Invalid search parameters"
// @Failure 404 {string} string "No projects found"
// @Failure 500 {string} string "Internal server error"
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

package handlers

import (
	"ProjectManagementService/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TaskInput struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	Priority          string `json:"priority"`
	Status            string `json:"status"`
	ResponsibleUserID int    `json:"responsible_user_id"`
	ProjectID         int    `json:"project_id"`
}

type TaskHandler struct {
	TaskModel models.TaskModel
}

func NewTaskHandler(taskModel models.TaskModel) *TaskHandler {
	return &TaskHandler{
		TaskModel: taskModel,
	}
}

// @Summary Get all tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
// @Failure 404 {string} string "No tasks found"
// @Failure 500 {string} string "Internal server error"
func (th *TaskHandler) GetAllTasksHandler(writer http.ResponseWriter, request *http.Request) {
	tasks, err := th.TaskModel.GetTasks()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(tasks) == 0 {
		writer.WriteHeader(http.StatusNotFound)
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

// @Summary Create a task
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body TaskInput true "Task"
// @Success 201 {string} string "Task created"
// @Router /tasks [post]
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
func (th *TaskHandler) CreateTaskHandler(writer http.ResponseWriter, request *http.Request) {
	var task models.Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, "data reading error: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = th.TaskModel.CreateTask(task.Title, task.Description, task.Priority, task.Status, task.ResponsibleUserID, task.ProjectID)
	if err != nil {
		http.Error(writer, "error creating task: "+err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)

}

// @Summary Get a task by ID
// @Tags tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [get]
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal server error"
func (th *TaskHandler) GetTaskHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	task, err := th.TaskModel.GetTaskById(id)
	if task == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonTask, err := json.Marshal(task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonTask)
}

// @Summary Update a task
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body TaskInput true "Task"
// @Success 200 {string} string "Task updated"
// @Router /tasks/{id} [put]
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal server error"
func (th *TaskHandler) UpdateTaskHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	task, err := th.TaskModel.GetTaskById(id)
	if task == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = th.TaskModel.UpdateTask(task.ID, task.Title, task.Description, task.Priority, task.Status, task.ResponsibleUserID, task.ProjectID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)

}

// @Summary Delete a task
// @Tags tasks
// @Param id path int true "Task ID"
// @Success 200 {string} string "Task deleted"
// @Router /tasks/{id} [delete]
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal server error"
// @Param id path int true "Task ID"
// @Success 200 {string} string "Task deleted"
// @Router /tasks/{id} [delete]
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal server error"
func (th *TaskHandler) DeleteTaskHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	deletedId, err := th.TaskModel.DeleteTask(id)
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

// @Summary Search tasks
// @Tags tasks
// @Produce json
// @Param title query string false "Task title"
// @Param status query string false "Task status"
// @Param priority query string false "Task priority"
// @Param assignee query string false "Task assignee"
// @Param project query string false "Task project"
// @Success 200 {array} models.Task
// @Router /tasks/search [get]
// @Failure 400 {string} string "No search parameters provided"
// @Failure 404 {string} string "No tasks found"
// @Failure 500 {string} string "Internal server error"
func (th *TaskHandler) SearchTasksHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Query().Get("title")
	status := request.URL.Query().Get("status")

	priority := request.URL.Query().Get("priority")

	assignee := request.URL.Query().Get("assignee")

	project := request.URL.Query().Get("project")

	if title == "" && status == "" && priority == "" && assignee == "" && project == "" {
		http.Error(writer, "No search parameters provided", http.StatusBadRequest)
		return
	}

	var (
		tasks []*models.Task
		err   error
	)
	if title != "" {
		tasks, err = th.TaskModel.SearchTaskByTitle(title)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	if status != "" {
		var statusEnumValue models.StatusEnum
		switch status {
		case string(models.New):
			statusEnumValue = models.New
		case string(models.InProgress):
			statusEnumValue = models.InProgress
		case string(models.Done):
			statusEnumValue = models.Done
		default:
			http.Error(writer, "Invalid status", http.StatusBadRequest)
			return
		}
		tasks, err = th.TaskModel.SearchTaskByStatus(statusEnumValue)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	if priority != "" {
		var priorityEnumValue models.PriorityEnum
		switch priority {
		case string(models.Low):
			priorityEnumValue = models.Low
		case string(models.Medium):
			priorityEnumValue = models.Medium
		case string(models.High):
			priorityEnumValue = models.High
		}
		tasks, err = th.TaskModel.SearchTaskByPriority(priorityEnumValue)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	if assignee != "" {
		assigneeId, err := strconv.Atoi(assignee)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		tasks, err = th.TaskModel.SearchTaskByResponsibleUserID(assigneeId)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	if project != "" {
		projectId, err := strconv.Atoi(project)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		tasks, err = th.TaskModel.SearchTaskByProjectID(projectId)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if len(tasks) == 0 {
		writer.WriteHeader(http.StatusNotFound)
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

}

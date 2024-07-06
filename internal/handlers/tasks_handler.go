package handlers

import (
	"ProjectManagementService/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllTasksHandler(taskModel *models.TaskModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tasks, err := taskModel.GetTasks()
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
}

func CreateTaskHandler(taskModel *models.TaskModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var task models.Task
		err := json.NewDecoder(request.Body).Decode(&task)
		if err != nil {
			http.Error(writer, "data reading error: "+err.Error(), http.StatusBadRequest)
			return
		}
		err = taskModel.CreateTask(task.Title, task.Description, task.Priority, task.Status, task.ResponsibleUserID, task.ProjectID)
		if err != nil {
			http.Error(writer, "error creating task: "+err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)

	}

}

func GetTaskHandler(taskModel *models.TaskModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		task, err := taskModel.GetTaskById(id)
		if task == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if task == nil {
			writer.WriteHeader(http.StatusNotFound)
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

}

func UpdateTaskHandler(taskModel *models.TaskModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		task, err := taskModel.GetTaskById(id)
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
		err = taskModel.UpdateTask(task.ID, task.Title, task.Description, task.Priority, task.Status, task.ResponsibleUserID, task.ProjectID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)

	}

}

func DeleteTaskHandler(taskModel *models.TaskModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		deletedId, err := taskModel.DeleteTask(id)
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

}

func SearchTasksHandler(taskModel *models.TaskModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
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
			tasks, err = taskModel.SearchTaskByTitle(title)
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
			tasks, err = taskModel.SearchTaskByStatus(statusEnumValue)
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
			tasks, err = taskModel.SearchTaskByPriority(priorityEnumValue)
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
			tasks, err = taskModel.SearchTaskByResponsibleUserID(assigneeId)
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
			tasks, err = taskModel.SearchTaskByProjectID(projectId)
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

}

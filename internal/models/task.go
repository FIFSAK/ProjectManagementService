package models

import "database/sql"

type PriorityEnum string
type StatusEnum string

const (
	Low    PriorityEnum = "low"
	Medium PriorityEnum = "medium"
	High   PriorityEnum = "high"
)

const (
	New        StatusEnum = "new"
	InProgress StatusEnum = "in_progress"
	Done       StatusEnum = "done"
)

type Task struct {
	ID                int          `json:"id"`
	Title             string       `json:"title"`
	Description       string       `json:"description"`
	Priority          PriorityEnum `json:"priority"`
	Status            StatusEnum   `json:"status"`
	ResponsibleUserID int          `json:"responsible_user_id"`
	ProjectID         int          `json:"project_id"`
	CreationDate      string       `json:"creation_date"`
	CompletionDate    string       `json:"completion_date"`
}

type TaskModel struct {
	DB *sql.DB
}

func NewTaskModel(db *sql.DB) *TaskModel {
	return &TaskModel{DB: db}
}

func (m *TaskModel) GetTasks() ([]*Task, error) {
	rows, err := m.DB.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	tasks := make([]*Task, 0)
	for rows.Next() {
		task := &Task{}
		var completionDate sql.NullString
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.ResponsibleUserID, &task.ProjectID, &task.CreationDate, &completionDate)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			task.CompletionDate = completionDate.String
		}
		tasks = append(tasks, task)
	}
	return tasks, nil

}

func (m *TaskModel) CreateTask(title, description string, priority PriorityEnum, status StatusEnum, responsibleUserID, projectID int) error {
	_, err := m.DB.Exec("INSERT INTO tasks (title, description, priority, status, responsible_user_id, project_id) VALUES ($1, $2, $3, $4, $5, $6)", title, description, priority, status, responsibleUserID, projectID)
	if err != nil {
		return err
	}
	return nil
}

func (m *TaskModel) GetTaskById(id int) (*Task, error) {
	row := m.DB.QueryRow("SELECT * FROM tasks WHERE id = $1", id)
	task := &Task{}
	var completionDate sql.NullString
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.ResponsibleUserID, &task.ProjectID, &task.CreationDate, &completionDate)
	if err != nil {
		return nil, err
	}
	if completionDate.Valid {
		task.CompletionDate = completionDate.String
	}
	return task, nil
}

func (m *TaskModel) UpdateTask(id int, title, description string, priority PriorityEnum, status StatusEnum, responsibleUserID, projectID int) error {
	_, err := m.DB.Exec("UPDATE tasks SET title = $1, description = $2, priority = $3, status = $4, responsible_user_id = $5, project_id = $6 WHERE id = $7", title, description, priority, status, responsibleUserID, projectID, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *TaskModel) DeleteTask(id int) error {
	_, err := m.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (m *TaskModel) SearchTaskByTitle(title string) ([]*Task, error) {
	rows, err := m.DB.Query("SELECT * FROM tasks WHERE title = $1", title)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	tasks := make([]*Task, 0)
	for rows.Next() {
		task := &Task{}
		var completionDate sql.NullString
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.ResponsibleUserID, &task.ProjectID, &task.CreationDate, &completionDate)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			task.CompletionDate = completionDate.String
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *TaskModel) SearchTaskByStatus(status StatusEnum) ([]*Task, error) {
	rows, err := m.DB.Query("SELECT * FROM tasks WHERE status = $1", status)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	tasks := make([]*Task, 0)
	for rows.Next() {
		task := &Task{}
		var completionDate sql.NullString
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.ResponsibleUserID, &task.ProjectID, &task.CreationDate, &completionDate)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			task.CompletionDate = completionDate.String
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *TaskModel) SearchTaskByPriority(priority PriorityEnum) ([]*Task, error) {
	rows, err := m.DB.Query("SELECT * FROM tasks WHERE priority = $1", priority)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	tasks := make([]*Task, 0)
	for rows.Next() {
		task := &Task{}
		var completionDate sql.NullString
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.ResponsibleUserID, &task.ProjectID, &task.CreationDate, &completionDate)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			task.CompletionDate = completionDate.String
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *TaskModel) SearchTaskByResponsibleUserID(responsibleUserID int) ([]*Task, error) {
	rows, err := m.DB.Query("SELECT * FROM tasks WHERE responsible_user_id = $1", responsibleUserID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	tasks := make([]*Task, 0)
	for rows.Next() {
		task := &Task{}
		var completionDate sql.NullString
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.ResponsibleUserID, &task.ProjectID, &task.CreationDate, &completionDate)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			task.CompletionDate = completionDate.String
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *TaskModel) SearchTaskByProjectID(projectID int) ([]*Task, error) {
	rows, err := m.DB.Query("SELECT * FROM tasks WHERE project_id = $1", projectID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	tasks := make([]*Task, 0)
	for rows.Next() {
		task := &Task{}
		var completionDate sql.NullString
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.ResponsibleUserID, &task.ProjectID, &task.CreationDate, &completionDate)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			task.CompletionDate = completionDate.String
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

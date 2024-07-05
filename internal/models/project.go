package models

import "database/sql"

type Project struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	CreationDate   string `json:"creation_date"`
	CompletionDate string `json:"completion_date"`
	ManagerID      int    `json:"manager_id"`
}

type ProjectModel struct {
	DB *sql.DB
}

func NewProjectModel(db *sql.DB) *ProjectModel {
	return &ProjectModel{DB: db}
}

func (pm *ProjectModel) GetProjects() ([]Project, error) {
	rows, err := pm.DB.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	projects := make([]Project, 0)
	for rows.Next() {
		project := Project{}
		var completionDate sql.NullString
		err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.CreationDate, &completionDate, &project.ManagerID)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			project.CompletionDate = completionDate.String
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (pm *ProjectModel) CreateProject(title, description string, managerID int) error {
	var id int
	err := pm.DB.QueryRow("INSERT INTO projects (title, description, manager_id) VALUES ($1, $2, $3) RETURNING id", title, description, managerID).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (pm *ProjectModel) GetProjectByID(id int) (*Project, error) {
	project := Project{}
	var completionDate sql.NullString
	err := pm.DB.QueryRow("SELECT * FROM projects WHERE id = $1", id).Scan(&project.ID, &project.Title, &project.Description, &project.CreationDate, &completionDate, &project.ManagerID)
	if err != nil {
		return nil, err
	}
	if completionDate.Valid {
		project.CompletionDate = completionDate.String
	}
	return &project, nil
}

func (pm *ProjectModel) UpdateProject(id int, title, description string, managerID int) error {
	_, err := pm.DB.Exec("UPDATE projects SET title = $1, description = $2, manager_id = $3 WHERE id = $4", title, description, managerID, id)
	if err != nil {
		return err
	}
	return nil
}

func (pm *ProjectModel) DeleteProject(id int) error {
	_, err := pm.DB.Exec("DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (pm *ProjectModel) GetProjectTasks(id int) ([]Task, error) {
	rows, err := pm.DB.Query("SELECT * FROM tasks WHERE project_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	tasks := make([]Task, 0)
	for rows.Next() {
		task := Task{}
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

func (pm *ProjectModel) SearchProjectsByTitle(title string) ([]Project, error) {
	rows, err := pm.DB.Query("SELECT * FROM projects WHERE title = $1", title)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	projects := make([]Project, 0)
	for rows.Next() {
		project := Project{}
		var completionDate sql.NullString
		err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.CreationDate, &completionDate, &project.ManagerID)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			project.CompletionDate = completionDate.String
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (pm *ProjectModel) SearchProjectsByManagerID(managerID int) ([]Project, error) {
	rows, err := pm.DB.Query("SELECT * FROM projects WHERE manager_id = $1", managerID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	projects := make([]Project, 0)
	for rows.Next() {
		project := Project{}
		var completionDate sql.NullString
		err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.CreationDate, &completionDate, &project.ManagerID)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			project.CompletionDate = completionDate.String
		}
		projects = append(projects, project)
	}
	return projects, nil
}

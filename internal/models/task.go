package models

type priorityEnum string
type statusEnum string

const (
	Low    priorityEnum = "low"
	Medium priorityEnum = "medium"
	High   priorityEnum = "high"
)

const (
	New        statusEnum = "new"
	InProgress statusEnum = "in_progress"
	Done       statusEnum = "done"
)

type Task struct {
	ID                int          `json:"id"`
	Title             string       `json:"title"`
	Description       string       `json:"description"`
	Priority          priorityEnum `json:"priority"`
	Status            statusEnum   `json:"status"`
	ResponsibleUserID int          `json:"responsible_user_id"`
	ProjectID         int          `json:"project_id"`
	CreationDate      string       `json:"creation_date"`
	CompletionDate    string       `json:"completion_date"`
}

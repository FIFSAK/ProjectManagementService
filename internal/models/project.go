package models

type Project struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	CreationDate   string `json:"creation_date"`
	CompletionDate string `json:"completion_date"`
	ManagerID      int    `json:"manager_id"`
}

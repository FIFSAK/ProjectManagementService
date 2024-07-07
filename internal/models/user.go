package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	RegistrationDate string `json:"registration_date"`
	Role             string `json:"role"`
}

type UserModel interface {
	GetUsers() ([]*User, error)
	CreateUser(name string, email string, role string) error
	GetUserById(id int) (*User, error)
	UpdateUser(id int, name string, email string, role string) error
	DeleteUser(id int) (int, error)
	SearchUserByEmail(email string) ([]*User, error)
	SearchUserByName(name string) ([]*User, error)
	GetUserTasks(id int) ([]*Task, error)
}

type UserModelImpl struct {
	DB *sql.DB
}

func (m *UserModelImpl) Error(s string) error {
	return fmt.Errorf(s)
}

func NewUserModel(db *sql.DB) *UserModelImpl {
	return &UserModelImpl{DB: db}
}

func (m *UserModelImpl) GetUsers() ([]*User, error) {
	rows, err := m.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	users := make([]*User, 0)
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.RegistrationDate, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}
func (m *UserModelImpl) CreateUser(name string, email string, role string) error {
	users, err := m.SearchUserByEmail(email)
	if err != nil {
		return err
	}
	if len(users) > 0 {
		return m.Error("User with this email already exists")
	}

	_, err = m.DB.Exec("INSERT INTO users (name, email, role) VALUES ($1, $2, $3)", name, email, role)
	if err != nil {
		return err
	}
	return nil

}

func (m *UserModelImpl) GetUserById(id int) (*User, error) {
	user := &User{}
	err := m.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.RegistrationDate, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (m *UserModelImpl) UpdateUser(id int, name string, email string, role string) error {
	_, err := m.DB.Exec("UPDATE users SET name = $1, email = $2, role = $3 WHERE id = $4", name, email, role, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModelImpl) DeleteUser(id int) (int, error) {
	row := m.DB.QueryRow("DELETE FROM users WHERE id = $1 RETURNING id", id)
	var deletedId int
	err := row.Scan(&deletedId)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *UserModelImpl) SearchUserByEmail(email string) ([]*User, error) {
	rows, err := m.DB.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	users := make([]*User, 0)
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.RegistrationDate, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (m *UserModelImpl) SearchUserByName(name string) ([]*User, error) {
	rows, err := m.DB.Query("SELECT * FROM users WHERE name = $1", name)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	users := make([]*User, 0)
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.RegistrationDate, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (m *UserModelImpl) GetUserTasks(id int) ([]*Task, error) {
	rows, err := m.DB.Query("SELECT * FROM tasks WHERE responsible_user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

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
		} else {
			task.CompletionDate = ""
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

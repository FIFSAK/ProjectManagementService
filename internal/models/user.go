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

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Error(s string) error {
	return fmt.Errorf(s)
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{DB: db}
}

func (m *UserModel) GetUsers() ([]*User, error) {
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
func (m *UserModel) CreateUser(name string, email string, role string) error {
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

func (m *UserModel) GetUserById(id int) (*User, error) {
	user := &User{}
	err := m.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.RegistrationDate, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (m *UserModel) UpdateUser(id int, name string, email string, role string) error {
	_, err := m.DB.Exec("UPDATE users SET name = $1, email = $2, role = $3 WHERE id = $4", name, email, role, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) DeleteUser(id int) error {
	_, err := m.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) SearchUserByEmail(email string) ([]*User, error) {
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

func (m *UserModel) SearchUserByName(name string) ([]*User, error) {
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

func (m *UserModel) GetUserTasks(id int) ([]*Task, error) {
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
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.ResponsibleUserID, &task.ProjectID, &task.CreationDate, &task.CompletionDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

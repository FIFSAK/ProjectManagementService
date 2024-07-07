package models

type MockUserModel struct {
	MockGetUsers          func() ([]*User, error)
	MockCreateUser        func(name string, email string, role string) error
	MockGetUserById       func(id int) (*User, error)
	MockUpdateUser        func(id int, name string, email string, role string) error
	MockDeleteUser        func(id int) (int, error)
	MockSearchUserByEmail func(email string) ([]*User, error)
	MockSearchUserByName  func(name string) ([]*User, error)
	MockGetUserTasks      func(id int) ([]*Task, error)
}

func (m *MockUserModel) GetUsers() ([]*User, error) {
	if m.MockGetUsers != nil {
		return m.MockGetUsers()
	}
	return nil, nil
}

func (m *MockUserModel) CreateUser(name string, email string, role string) error {
	if m.MockCreateUser != nil {
		return m.MockCreateUser(name, email, role)
	}
	return nil
}

func (m *MockUserModel) GetUserById(id int) (*User, error) {
	if m.MockGetUserById != nil {
		return m.MockGetUserById(id)
	}
	return nil, nil
}

func (m *MockUserModel) UpdateUser(id int, name string, email string, role string) error {
	if m.MockUpdateUser != nil {
		return m.MockUpdateUser(id, name, email, role)
	}
	return nil
}

func (m *MockUserModel) DeleteUser(id int) (int, error) {
	if m.MockDeleteUser != nil {
		return m.MockDeleteUser(id)
	}
	return 0, nil
}

func (m *MockUserModel) SearchUserByEmail(email string) ([]*User, error) {
	if m.MockSearchUserByEmail != nil {
		return m.MockSearchUserByEmail(email)
	}
	return nil, nil
}

func (m *MockUserModel) SearchUserByName(name string) ([]*User, error) {
	if m.MockSearchUserByName != nil {
		return m.MockSearchUserByName(name)
	}
	return nil, nil
}

func (m *MockUserModel) GetUserTasks(id int) ([]*Task, error) {
	if m.MockGetUserTasks != nil {
		return m.MockGetUserTasks(id)
	}
	return nil, nil
}

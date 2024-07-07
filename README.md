# Project Management Service

## Routers

### Health Check

- **Endpoint:** `GET /health-check`
    - **Response:** `OK`

### Get Users
- **Endpoint:** `GET /users`
    - **Body:**
    - **Response:**
      ```json
      [
      {
      "id": 1,
      "name": "John Doe", 
      "email": "example@mail.com",
      "role": "admin",
      "registration_date": "2021-09-01T00:00:00Z"
      }
      ]
      ```

### Create User

- **Endpoint:** `POST /users`
    - **Body:**
      ```json
      { 
      "name": "John Doe", 
      "email": "example@mail.com",
      "role": "admin"
      }
      ```

### Get User
- **Endpoint:** `GET /users/{ID}`
    - **Body:**
    - **Response:**
      ```json
      {
      "id": 1,
      "name": "John Doe", 
      "email": "example@mail.ru",
      "registration_date": "2021-09-01T00:00:00Z"
      }
      ```
    

### Update User

- **Endpoint:** `PUT /users/{ID}`
    - **Body:**
      ```json
      { 
      "name": "New John Doe", 
      "email": "example@mail.com",
      "role": "admin"
      }
      ```
### Delete User
- **Endpoint:** `DELETE /users/{ID}`

### Get User's Tasks
- **Endpoint:** `GET /users/{ID}/tasks`


### Search User

- **Endpoint:** `PUT /users/search?name=John Doe` | ?email={email}

### Get Tasks

- **Endpoint:** `GET /tasks`

### Create Task
- **Endpoint:** `POST /tasks`
    - **Body:**
      ```json
      { 
      "title": "Task 1", 
      "description": "Task 1 description",
      "priority": "high medium low",
      "status": "new done in_progress",
      "responsible_user_id": 1,
      "project_id": 1
      }
      ```

### Get Task
- **Endpoint:** `GET /tasks/{ID}`
    - **Body:**
    - **Response:**
      ```json
      {
      "id": 1,
      "title": "Task 1", 
      "description": "Task 1 description",
      "priority": "high medium low",
      "status": "new done in_progress",
      "responsible_user_id": 1,
      "project_id": 1,
      "creation_date": "2021-09-01T00:00:00Z",
      "completion_date": ""
      }
      ```

### Update Task
- **Endpoint:** `PUT /tasks/{ID}`
    - **Body:**
      ```json
      { 
      "title": "New Task 1", 
      "description": "Task 1 description",
      "priority": "high medium low",
      "status": "new done in_progress",
      "responsible_user_id": 1,
      "project_id": 1
      }
      ```
### Delete Task
- **Endpoint:** `DELETE /tasks/{ID}`

### Search Task
- **Endpoint:** `GET /tasks/search?title=Task 1` | ?priority={priority} | ?status={status} | ?assignee={responsible_user_id} | ?project_id={project_id}

### Get Projects
- **Endpoint:** `DELETE /projects`

### Create Project
- **Endpoint:** `POST /projects`
    - **Body:**
      ```json
      { 
      "title": "Project 1", 
      "description": "Project 1 description",
      "manager_id": 1
      }
      ```
### Get Project
- **Endpoint:** `GET /projects/{ID}`
    - **Body:**
    - **Response:**
      ```json
      {
      "id": 1,
      "title": "Project 1", 
      "description": "Project 1 description",
      "manager_id": 1,
      "creation_date": "2021-09-01T00:00:00Z",
      "completion_date": ""
      }
      ```
### Update Project
- **Endpoint:** `PUT /projects/{ID}`
    - **Body:**
      ```json
      { 
      "title": "New Project 1", 
      "description": "Project 1 description",
      "manager_id": 1
      }
      ```
### Delete Project
- **Endpoint:** `DELETE /projects/{ID}`

### Search Project
- **Endpoint:** `GET /projects/search?title=Project 1` | ?manager={user_id}

## Models Structure

```sql
Users {
    id: int,
    name: string,
    email: string,
    role: string,
    registration_date: date,
}
Tasks {
    id: int,
    title: string,
    description: string,
    priority: string,
    status: string,
    responsible_user_id: int,
    project_id: int,
    creation_date: date,
    completion_date: date,
}
Projects {
    id: int,
    title: string,
    description: string,
    manager_id: int,
    creation_date: date,
    completion_date: date,
}
```

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/FIFSAK/ProjectManagementService
   cd ProjectManagementService
   ```
replace func initializeDB() in dbconn.go with the following code:
```go
connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	os.Getenv("user"), os.Getenv("password"), os.Getenv("host"),
	os.Getenv("port"), os.Getenv("dbname"))
db, err := sql.Open("postgres", connStr)
if err != nil {
	return nil, err
}
if err := db.Ping(); err != nil {
	return nil, err
}

migrationUp(db)

return db, nil
```
and in migrationUp function marked block code with the following code:
```go
m, err := migrate.NewWithDatabaseInstance(
	"file:///usr/src/app/migrations",
	"postgres", driver)
```


2. **Build the Docker images:**
   ```bash
   make build
   ```
3. **Start the Docker containers:**
   ```bash
   make up
   ```
4. **Check the health of the server:**
   Open your browser and go to http://localhost:8080/health-check to ensure the server is running properly.

6. **Stop the Docker containers:**
   ```bash
   make down
   ```

**LINK: https://projectmanagementservice.onrender.com**

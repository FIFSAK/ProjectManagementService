# Todo List

## Routers

### Health Check

- **Endpoint:** `GET /health-check`
    - **Response:** `OK`

### Create Task

- **Endpoint:** `POST /api/todo-list/tasks`
    - **Body:**
      ```json
      { 
      "title": "Купить книгу", 
      "activeAt": "2023-08-04" 
      }
      ```

### Update Task

- **Endpoint:** `PUT /api/todo-list/tasks/{ID}`
    - **Body:**
      ```json
      { 
      "title": "Купить книгу - Высоконагруженные приложения", 
      "activeAt": "2023-08-05" 
      }
      ```
### Delete Task

- **Endpoint:** `DELETE /api/todo-list/tasks/{ID}`
    - **Response:** `status: 204`

### Mark Task as Done

- **Endpoint:** `PUT /api/todo-list/tasks/{ID}/done`
    - **Response:** `status: 204`

### Get Tasks

- **Endpoint:** `GET /api/todo-list/tasks?status=active` status is optional
    - **Response:**
      ```json
      [
          {
          "id": "65f19340848f4be025160391",
          "title": "Купить книгу - Высоконагруженные приложения",
          "activeAt": "2023-08-05"
          },
          {
          "id": "75f19340848f4be025160392",
          "title": "Купить квартиру",
          "activeAt": "2023-08-05"
          },
          {
          "id": "45f19340848f4be025160394",
          "title": "Купить машину",
          "activeAt": "2023-08-05"
          }
      ]
      ```

## Models Structure

```sql
Task
{
    id: int,
    status: string,
    title: string,
    activeAt: string,
}
```

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/FIFSAK/Hl-task1
   cd Hl-task1
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

**LINK: https://todo-list-e4d2.onrender.com/**

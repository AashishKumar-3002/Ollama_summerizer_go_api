# FealtyX Student API
FealtyX Student API is a RESTful API developed in Go for managing student data and generating summaries using Ollama.

## Features

- CRUD operations for student records
- Persistent data storage with in-memory data storage
- Student summary generation via Ollama
- Concurrent-safe operations
- Input validation

## Prerequisites

- Go 1.16 or higher
- PostgreSQL installed and running locally
- Ollama installed and running locally

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/AashishKumar-3002/fealtyx-student-api.git
    cd fealtyx-student-api
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Set up PostgreSQL:
    - Install PostgreSQL:
        ```bash
        sudo apt-get update
        sudo apt-get install postgresql postgresql-contrib
        ```
    - Start PostgreSQL service:
        ```bash
        sudo service postgresql start
        ```
    - Create a database and user:
        ```bash
        sudo -u postgres psql
        CREATE DATABASE fealtyx;
        CREATE USER fealtyx_user WITH ENCRYPTED PASSWORD 'yourpassword';
        GRANT ALL PRIVILEGES ON DATABASE fealtyx TO fealtyx_user;
        \q
        ```

4. Create a `.env` file in the root directory with the following content:
    ```env
    DB_CONNECTION_STRING=postgres://fealtyx_user:yourpassword@localhost:5432/fealtyx?sslmode=disable
    ```
    Replace `yourpassword` with the password you set for the `fealtyx_user`.

## Usage

1. Start the server:
    ```bash
    go run cmd/api/main.go
    ```
    The server will start on [http://localhost:8080](http://localhost:8080).

2. Use the following endpoints:

    - Create a student: `POST /students`
    - Get all students: `GET /students`
    - Get a student by ID: `GET /students/{id}`
    - Update a student: `PUT /students/{id}`
    - Delete a student: `DELETE /students/{id}`
    - Generate a student summary: `GET /students/{id}/summary`

## API Examples

### Create a student

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe","age":20,"email":"john@example.com"}' http://localhost:8080/students
```

### Get all students

```bash
curl http://localhost:8080/students
```

### Get a student by ID

```bash
curl http://localhost:8080/students/1
```

### Update a student

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name":"John Doe","age":21,"email":"john.doe@example.com"}' http://localhost:8080/students/1
```

### Delete a student

```bash
curl -X DELETE http://localhost:8080/students/1
```

### Generate a student summary

```bash
curl http://localhost:8080/students/1/summary
```

## Installation instructions for Ollama

To integrate Ollama with your Go project, follow these steps:

1. **Install** Ollama on your localhost by following: https://github.com/ollama/ollama/blob/main/README.md#quickstart
2. **Install Llama3** language model for Ollama. For this project, we are using the 1b parameter model. Hence, our model name is `llama3_1b`. To install the model, run the following command:
    ```bash
    ollama install llama3_1b
    ```
3. **Make API Requests** to localhost and pass the Student object to generate the summary. Perform prompt engineering to get the summary for the Student.

Note: By default, Ollama listens on port 11434. This project is configured to use `OLLAMA_PORT=12345` If you need to use a different port, set the `OLLAMA_PORT` environment variable in the `.env` file to your desired port number (e.g., `OLLAMA_PORT=12345`).
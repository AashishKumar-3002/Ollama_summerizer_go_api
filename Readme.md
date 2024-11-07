# FealtyX Student API

FealtyX Student API is a RESTful API developed in Go for managing student data and generating summaries using Ollama. It supports both in-memory storage and PostgreSQL database storage.

## Features

- CRUD operations for student records
- Flexible storage options: in-memory or PostgreSQL database
- Student summary generation via Ollama
- Concurrent-safe operations
- Input validation

## Deployed Base URL

The API is deployed and accessible at:

https://ollama-summerizer-go-api.onrender.com

You can use this base URL to make requests to the API endpoints described below.

## Prerequisites

- Go 1.23.2 or higher
- PostgreSQL (optional, for database storage)
- Ollama installed and running locally

## Installation

1. Clone the repository:
   
   git clone https://github.com/AashishKumar-3002/FealtyX.git
   cd FealtyX
   

2. Install dependencies:
   
   go mod tidy
   

3. Set up PostgreSQL (optional):
   - Install and start PostgreSQL
   - Create a database and user for the application

4. Create a .env file in the root directory with the following content:
   
   DATABASE_URL="postgres://username:password@localhost:5432/dbname?sslmode=disable"
   PORT=8080
   OLLAMA_PORT=12345
   
   Replace the DATABASE_URL with your PostgreSQL connection string. If left empty, the application will use in-memory storage.

## Usage

1. Start the server:
   
   go run main.go
   
   The server will start on the port specified in the .env file (default: 8080).

2. Use the following endpoints:
   - Create a student: POST /students
   - Get all students: GET /students
   - Get a student by ID: GET /students/{id}
   - Update a student: PUT /students/{id}
   - Delete a student: DELETE /students/{id}
   - Generate a student summary: GET /students/{id}/summary

### API Examples

Replace http://localhost:8080 with https://ollama-summerizer-go-api.onrender.com when using the deployed version.

- Create a student
  
  curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe","age":20,"email":"john@example.com"}' https://ollama-summerizer-go-api.onrender.com/students
  

- Get all students
  
  curl https://ollama-summerizer-go-api.onrender.com/students
  

- Get a student by ID
  
  curl https://ollama-summerizer-go-api.onrender.com/students/1
  

- Update a student
  
  curl -X PUT -H "Content-Type: application/json" -d '{"name":"John Doe","age":21,"email":"john.doe@example.com"}' https://ollama-summerizer-go-api.onrender.com/students/1
  

- Delete a student
  
  curl -X DELETE https://ollama-summerizer-go-api.onrender.com/students/1
  

- Generate a student summary
  
  curl https://ollama-summerizer-go-api.onrender.com/students/1/summary
  

## Ollama Integration

This project uses Ollama for generating student summaries. To set up Ollama:

1. Install Ollama on your localhost by following the instructions at: https://github.com/ollama/ollama#quickstart

2. Install the Llama3 language model (1b parameter version):
   
   ollama install llama3.2:1b
   

3. Ensure Ollama is running on the port specified in the .env file (default: 12345).

## Testing

To run the tests, use the following command:


go test ./...


## License

This project is licensed under the Apache License 2.0. See the LICENSE file for details.
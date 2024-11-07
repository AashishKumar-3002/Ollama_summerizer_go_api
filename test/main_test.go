package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/AashishKumar-3002/FealtyX/internal/database"
	"github.com/AashishKumar-3002/FealtyX/internal/handlers"
	"github.com/AashishKumar-3002/FealtyX/internal/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var db *sql.DB
var h *handlers.Handler

func TestMain(m *testing.M) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Set up test database
	var err error
	db, err = database.Connect(os.Getenv("TEST_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	h = handlers.NewHandler(db)

	// Run tests
	code := m.Run()

	// Clean up
	db.Exec("DROP TABLE IF EXISTS students")

	os.Exit(code)
}

func TestCreateStudent(t *testing.T) {
	student := models.Student{Name: "John Doe", Age: 20, Email: "john@example.com"}
	body, _ := json.Marshal(student)

	req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/students", h.CreateStudent).Methods("POST")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var createdStudent models.Student
	json.Unmarshal(rr.Body.Bytes(), &createdStudent)
	if createdStudent.Name != student.Name {
		t.Errorf("handler returned unexpected body: got %v want %v", createdStudent.Name, student.Name)
	}
}

func TestGetAllStudents(t *testing.T) {
	req, _ := http.NewRequest("GET", "/students", nil)
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/students", h.GetAllStudents).Methods("GET")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var students []models.Student
	json.Unmarshal(rr.Body.Bytes(), &students)
	if len(students) == 0 {
		t.Errorf("handler returned no students")
	}
}

func TestConcurrency(t *testing.T) {
    numRequests := 100
    done := make(chan bool)

    for i := 0; i < numRequests; i++ {
        go func(i int) {
            student := models.Student{
                Name:  fmt.Sprintf("Test User %d", i),
                Age:   25,
                Email: fmt.Sprintf("test%d@example.com", i),
            }
            body, err := json.Marshal(student)
            if err != nil {
                t.Errorf("failed to marshal student: %v", err)
                done <- false
                return
            }

            req, err := http.NewRequest("POST", "/students", bytes.NewBuffer(body))
            if err != nil {
                t.Errorf("failed to create request: %v", err)
                done <- false
                return
            }
            rr := httptest.NewRecorder()

            r := mux.NewRouter()
            r.HandleFunc("/students", h.CreateStudent).Methods("POST")
            r.ServeHTTP(rr, req)

            if rr.Code != http.StatusOK {
                t.Errorf("failed to create student: %v", rr.Body.String())
                done <- false
                return
            }

            done <- true
        }(i)
    }

    successCount := 0
    for i := 0; i < numRequests; i++ {
        if <-done {
            successCount++
        } else {
            t.Errorf("failed to complete request %d", i)
        }
    }

    // Check if all students were created
    req, err := http.NewRequest("GET", "/students", nil)
    if err != nil {
        t.Fatalf("failed to create request: %v", err)
    }
    rr := httptest.NewRecorder()

    r := mux.NewRouter()
    r.HandleFunc("/students", h.GetAllStudents).Methods("GET")
    r.ServeHTTP(rr, req)

    var students []models.Student
    err = json.Unmarshal(rr.Body.Bytes(), &students)
    if err != nil {
        t.Fatalf("failed to unmarshal response: %v", err)
    }
    if len(students) != successCount {
        t.Errorf("expected %d students, got %d", successCount, len(students))
    }
}
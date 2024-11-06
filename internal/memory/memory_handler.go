package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AashishKumar-3002/FealtyX/internal/ai"
	"github.com/AashishKumar-3002/FealtyX/internal/models"
	"github.com/AashishKumar-3002/FealtyX/internal/storage"
	"github.com/gorilla/mux"
)

type API struct {
    storage *storage.Storage
}

func NewAPI(storage *storage.Storage) *API {
    return &API{storage: storage}
}

func (a *API) CreateStudent(w http.ResponseWriter, r *http.Request) {
    var student models.Student
    if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := validateStudent(student); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    createdStudent, err := a.storage.Create(student)
    if err != nil {
        http.Error(w, "Failed to create student", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdStudent)
}

func (a *API) GetAllStudents(w http.ResponseWriter, r *http.Request) {
    students := a.storage.GetAll()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(students)
}

func (a *API) GetStudentByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    student, err := a.storage.GetByID(id)
    if err != nil {
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(student)
}

func (a *API) UpdateStudent(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    var student models.Student
    if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := validateStudent(student); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    updatedStudent, err := a.storage.Update(id, student)
    if err != nil {
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedStudent)
}

func (a *API) DeleteStudent(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    if err := a.storage.Delete(id); err != nil {
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (a *API) GenerateStudentSummary(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    student, err := a.storage.GetByID(id)
    if err != nil {
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    summary, err := ai.GenerateStudentSummary(student)
    if err != nil {
        log.Printf("Error generating summary: %v", err)
        http.Error(w, "Failed to generate summary", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"summary": summary})
}

func validateStudent(student models.Student) error {
    if student.Name == "" {
        return fmt.Errorf("name is required")
    }
    if student.Age <= 0 {
        return fmt.Errorf("age must be positive")
    }
    if student.Email == "" {
        return fmt.Errorf("email is required")
    }
    // You can add more validation rules here
    return nil
}

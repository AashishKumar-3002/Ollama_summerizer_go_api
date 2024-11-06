package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AashishKumar-3002/FealtyX/internal/database"
	"github.com/AashishKumar-3002/FealtyX/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	db, err := database.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create router
	r := mux.NewRouter()

	// Initialize handlers
	h := handlers.NewHandler(db)

	// Define routes
	r.HandleFunc("/students", h.CreateStudent).Methods("POST")
	r.HandleFunc("/students", h.GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{id}", h.GetStudent).Methods("GET")
	r.HandleFunc("/students/{id}", h.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", h.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}/summary", h.GetStudentSummary).Methods("GET")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

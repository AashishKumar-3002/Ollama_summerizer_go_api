package models

import (
	"database/sql"
	"errors"

	"github.com/AashishKumar-3002/FealtyX/pkg/validator"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required,gte=1,lte=150"`
	Email string `json:"email" validate:"required,email"`
}

func (s *Student) Validate() error {
	return validator.Validate(s)
}

func (s *Student) Create(db *sql.DB) error {
	return db.QueryRow("INSERT INTO students (name, age, email) VALUES ($1, $2, $3) RETURNING id",
		s.Name, s.Age, s.Email).Scan(&s.ID)
}

func GetAllStudents(db *sql.DB) ([]Student, error) {
	rows, err := db.Query("SELECT id, name, age, email FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Age, &s.Email); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}

func GetStudent(db *sql.DB, id int) (*Student, error) {
	var s Student
	err := db.QueryRow("SELECT id, name, age, email FROM students WHERE id = $1", id).
		Scan(&s.ID, &s.Name, &s.Age, &s.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("student not found")
	}
	return &s, err
}

func (s *Student) Update(db *sql.DB, id int) error {
	res, err := db.Exec("UPDATE students SET name = $1, age = $2, email = $3 WHERE id = $4",
		s.Name, s.Age, s.Email, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("student not found")
	}
	s.ID = id
	return nil
}

func DeleteStudent(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM students WHERE id = $1", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("student not found")
	}
	return nil
}

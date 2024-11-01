package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/train-do/Golang-embed-dan-pengenalan-golang-web/model"
)

func CreateStudent(tx *sql.Tx, student *model.Student) (int, error) {
	query := `INSERT INTO "Student" (user_id, name, entry_year, phone_number, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err := tx.QueryRow(query, student.UserId, student.Name, student.EntryYear, student.PhoneNumber, time.Now()).Scan(&student.Id)
	if err != nil {
		fmt.Println("Create Student :", err)
		tx.Rollback()
		return -1, err
	}
	return student.Id, nil
}

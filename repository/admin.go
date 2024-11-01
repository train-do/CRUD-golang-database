package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/train-do/Golang-embed-dan-pengenalan-golang-web/model"
)

// type RepoAdmin struct {
// 	Admin model.Admin
// }

//	func (r *RepoAdmin) Create(tx *sql.Tx) (int, error) {
//		query := `INSERT INTO "Admin" (user_id, name, createdat) VALUES ($1, $2, $3);`
//		err := tx.QueryRow(query, r.User_id, r.Name, time.Now()).Scan(&r.Id)
//		if err != nil {
//			fmt.Println("Create Admin :", err)
//			tx.Rollback()
//			return -1, err
//		}
//		return r.Id, nil
//	}
func CreateAdmin(tx *sql.Tx, admin *model.Admin) (int, error) {
	query := `INSERT INTO "Admin" (user_id, name, created_at) VALUES ($1, $2, $3) RETURNING id;`
	err := tx.QueryRow(query, admin.User_id, admin.Name, time.Now()).Scan(&admin.Id)
	if err != nil {
		fmt.Println("Create Admin Error:", err)
		tx.Rollback()
		return -1, err
	}
	return admin.Id, nil
}

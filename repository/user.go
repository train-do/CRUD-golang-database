package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/train-do/Golang-embed-dan-pengenalan-golang-web/model"
)

type RepoUser struct {
	User model.User
}

func (r *RepoUser) Create(tx *sql.Tx) (int, error) {
	query := `INSERT INTO "User" (username, password, createdat) VALUES ($1, $2, $3);`
	err := tx.QueryRow(query, r.User.Username, r.User.Password, time.Now()).Scan(&r.User.Id)
	if err != nil {
		fmt.Println("Create User Error: ", err)
		tx.Rollback()
		return -1, err
	}
	return r.User.Id, nil
}
func Login(tx *sql.Tx, user model.User) (int, error) {
	query := fmt.Sprintf(`select id from "User" where username = '%s' and password = '%s';`, user.Username, user.Password)
	var id int
	err := tx.QueryRow(query).Scan(&id)
	if err != nil {
		fmt.Println("Login User Error: ", err)
		tx.Rollback()
		return -1, err
	}
	return id, nil
}
func CreateUser(tx *sql.Tx, user *model.User) (int, error) {
	// fmt.Printf("Repo : %+v\n", user)
	query := `INSERT INTO "User" (username, password, created_at) VALUES ($1, $2, $3) RETURNING id;`
	err := tx.QueryRow(query, user.Username, user.Password, time.Now()).Scan(&user.Id)
	if err != nil {
		fmt.Println("Create User Error: ", err)
		tx.Rollback()
		return -1, err
	}
	return user.Id, nil
}

//	func ReadAllUser(db *sql.DB) ([]model.User, error) {
//		query := fmt.Sprintf(`select * from "User";`)
//		var allUser []model.User = []model.User{}
//		rows, err := db.Query(query)
//		if err != nil {
//			fmt.Println("Read All User Error: ", err)
//			return []model.User{}, err
//		}
//		for rows.Next() {
//			var user model.User
//			rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
//			if err != nil {
//				fmt.Println("Read User Error: ", err)
//			}
//			allUser = append(allUser, user)
//		}
//		return allUser, nil
//	}
func ReadAllUser(db *sql.DB) ([]model.User, error) {
	query := `select * from "User";`
	var allUser []model.User = []model.User{}
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Read All User Error: ", err)
		return []model.User{}, err
	}
	for rows.Next() {
		var user model.User
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
		if err != nil {
			fmt.Println("Read User Error: ", err)
		}
		allUser = append(allUser, user)
	}
	return allUser, nil
}

// func ReadAllUser(db *sql.DB) ([]model.AllUser, error) {
// 	query := fmt.Sprintf(`select * from "User" u
// 	left join "Admin" a on u.id = a.user_id
// 	left join "Mentor" m on u.id = m.user_id
// 	left join "Student" s on u.id = s.user_id ;`)
// 	var users []model.AllUser = []model.AllUser{}
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		fmt.Println("Read All User Error: ", err)
// 		return []model.AllUser{}, err
// 	}
// 	for rows.Next() {
// 		var user model.AllUser
// 		var adminField, mentorField, studentField sql.NullString // Gunakan sql.NullString untuk menangani nilai NULL

// 		err := rows.Scan(&user.User.Id, &user.User.Username, &user.User.Password, &user.User.CreatedAt, &user.Admin.Id, &user.Admin.User_id, &user.Admin.Name, &user.Admin.CreatedAt, &user.Mentor.Id, &user.Mentor.UserId, &user.Mentor.Name, &user.Mentor.Experience, &user.Mentor.Degree, &user.Mentor.PhoneNumber, &user.Mentor.CreatedAt, &user.Student.Id, &user.Student.UserId, &user.Student.Name, &user.Student.EntryYear, &user.Student.PhoneNumber, &user.Student.CreatedAt)
// 		users = append(users, user)
// 	}
// 	return users, nil
// }

package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
}

//	type User struct {
//		Id        sql.NullInt64
//		Username  sql.NullString
//		Password  sql.NullString
//		CreatedAt sql.NullTime
//	}
// type AllUser struct {
// 	Id       int     `json:"id,omitempty"`
// 	Username string  `json:"username,omitempty"`
// 	Password string  `json:"password,omitempty"`
// 	Admin    Admin   `json:"admin,omitempty"`
// 	Mentor   Mentor  `json:"mentor,omitempty"`
// 	Student  Student `json:"student,omitempty"`
// }

func (u User) Connect() (*sql.DB, error) {
	connStr := "user=postgres dbname=Ojek sslmode=disable password=superUser host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, err
}

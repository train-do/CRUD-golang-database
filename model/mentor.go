package model

import "time"

type Mentor struct {
	Id          int       `json:"id,omitempty"`
	UserId      int       `json:"userId,omitempty"`
	Name        string    `json:"name,omitempty"`
	Experience  string    `json:"experience,omitempty"`
	Degree      string    `json:"degree,omitempty"`
	PhoneNumber string    `json:"PhoneNumber,omitempty"`
	CreatedAt   time.Time `json:"CreatedAt,omitempty"`
}

// type Mentor struct {
// 	Id          sql.NullInt64  `json:"id,omitempty"`
// 	UserId      sql.NullInt64  `json:"userId,omitempty"`
// 	Name        sql.NullString `json:"name,omitempty"`
// 	Experience  sql.NullString `json:"experience,omitempty"`
// 	Degree      sql.NullString `json:"degree,omitempty"`
// 	PhoneNumber sql.NullString `json:"PhoneNumber,omitempty"`
// 	CreatedAt   sql.NullTime   `json:"CreatedAt,omitempty"`
// }

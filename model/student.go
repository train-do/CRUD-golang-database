package model

import (
	"time"
)

type Student struct {
	Id          int       `json:"id,omitempty"`
	UserId      int       `json:"userId,omitempty"`
	Name        string    `json:"name,omitempty"`
	EntryYear   int       `json:"EntryYear,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	CreatedAt   time.Time `json:"CreatedAt,omitempty"`
}

// type Student struct {
// 	Id          sql.NullInt64  `json:"id,omitempty"`
// 	UserId      sql.NullInt64  `json:"userId,omitempty"`
// 	Name        sql.NullString `json:"name,omitempty"`
// 	EntryYear   sql.NullInt64  `json:"EntryYear,omitempty"`
// 	PhoneNumber sql.NullString `json:"phoneNumber,omitempty"`
// 	CreatedAt   sql.NullTime   `json:"CreatedAt,omitempty"`
// }

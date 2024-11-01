package model

import (
	"time"
)

type Admin struct {
	Id        int       `json:"id,omitempty"`
	User_id   int       `json:"userId,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

// type Admin struct {
// 	Id        sql.NullInt64  `json:"id,omitempty"`
// 	User_id   sql.NullInt64  `json:"userId,omitempty"`
// 	Name      sql.NullString `json:"name,omitempty"`
// 	CreatedAt sql.NullTime   `json:"CreatedAt,omitempty"`
// }

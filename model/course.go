package model

import "time"

type Course struct {
	Id        int
	MentorId  int
	Title     string
	AdminId   int
	CreatedAt time.Time
}

package model

import "time"

type Schedule struct {
	ClassId   int
	StartTime time.Time
	EndTime   time.Time
	AdminId   int
	CreateAt  time.Time
}

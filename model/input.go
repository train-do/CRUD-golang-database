package model

// type Login struct {
// 	username string
// 	password string
// }
// type RegisterMentor struct {
// 	name string
// 	experience string
// 	degree string
// 	phoneNumber string
// }
// type RegisterStudent struct {
// 	name string

// 	phoneNumber string
// }
type Input struct {
	Login           User
	Register        User
	RegisterMentor  Mentor
	RegisterStudent Student
	UpdateMentor    Mentor
	Course          Course
	Schedule        Schedule
}

var InputReq Input

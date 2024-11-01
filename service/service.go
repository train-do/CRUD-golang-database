package service

import (
	"database/sql"
	"fmt"

	"github.com/train-do/Golang-embed-dan-pengenalan-golang-web/model"
	"github.com/train-do/Golang-embed-dan-pengenalan-golang-web/repository"
	"github.com/train-do/Golang-embed-dan-pengenalan-golang-web/utils"
)

var User_Id int

func init() {
	utils.DecodeFromJSON("input.json", &model.InputReq)
	// utils.DecodeInput()
}
func RunningApp(db *sql.DB) {
	// fmt.Printf("%+v |||||||||||\n", model.InputReq)
	for {
		fmt.Println("1. POST /login\n2. POST /addMentor\n3. POST /addStudent\n4. GET /allUser\n5. PUT /editMentor\n6. POST /addCourse\n7. POST /addSchedule\n8. DELETE /deleteSchedule\n0. Logout")
		var input string
		fmt.Print("Select endpoint : ")
		fmt.Scanln(&input)
		switch input {
		case "1":
			Login(db)
		case "2":
			registerMentor(db)
		case "3":
			registerStudent(db)
		case "4":
			fetchAllUser(db)
		case "5":
			editMentor(db)
		case "6":
		case "7":
		case "8":
		case "9":
		case "0":
		}
	}
}
func Login(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("error db begin register :", err)
	}
	User_Id, err = repository.Login(tx, model.InputReq.Login)
	if err != nil {
		fmt.Println("error user not found :", err)
		fmt.Println("Create User Error :", err)
		response := model.Response{
			StatusCode: 400,
			Message:    "Invalid User",
			Data:       "",
		}
		utils.EncodeToJSON(response)
	}
	response := model.Response{
		StatusCode: 200,
		Message:    "Success Login",
		Data:       "",
	}
	utils.EncodeToJSON(response)
}
func registerMentor(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("error db begin register :", err)
	}
	repoUser := repository.RepoUser{
		User: model.InputReq.Register,
	}
	fmt.Printf("%+v, %v", &repoUser, tx)
	user_id, err := repository.CreateInterface(repoUser, tx)
	// user_id, err := repository.CreateUser(tx, &model.InputReq.Register)
	// if err != nil {
	// 	fmt.Println("Create User Error :", err)
	// 	response := model.Response{
	// 		StatusCode: 500,
	// 		Message:    "Internal Server Error",
	// 		Data:       "",
	// 	}
	// 	utils.EncodeToJSON(response)
	// 	return
	// }
	// model.InputReq.RegisterMentor.UserId = user_id
	// fmt.Printf("Service : %+v, %d\n", model.InputReq.Register, user_id)
	// mentorId, err := repository.CreateMentor(tx, &model.InputReq.RegisterMentor)
	// if err != nil {
	// 	fmt.Println("Create Mentor Error :", err)
	// 	response := model.Response{
	// 		StatusCode: 500,
	// 		Message:    "Internal Server Error",
	// 		Data:       "",
	// 	}
	// 	utils.EncodeToJSON(response)
	// 	return
	// }
	tx.Commit()
	response := model.Response{
		StatusCode: 201,
		Message:    "Succes Create Mentor",
		Data:       "",
	}
	utils.EncodeToJSON(response)
	fmt.Println(user_id)
}
func registerStudent(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("error db begin register :", err)
	}
	user_id, err := repository.CreateUser(tx, &model.InputReq.Register)
	if err != nil {
		fmt.Println("Create User Error :", err)
		response := model.Response{
			StatusCode: 500,
			Message:    "Internal Server Error",
			Data:       "",
		}
		utils.EncodeToJSON(response)
		return
	}
	model.InputReq.RegisterStudent.UserId = user_id
	fmt.Printf("Service : %+v, %d\n", model.InputReq.Register, user_id)
	mentorId, err := repository.CreateStudent(tx, &model.InputReq.RegisterStudent)
	if err != nil {
		fmt.Println("Create Admin Error :", err)
		response := model.Response{
			StatusCode: 500,
			Message:    "Internal Server Error",
			Data:       "",
		}
		utils.EncodeToJSON(response)
		return
	}
	tx.Commit()
	response := model.Response{
		StatusCode: 201,
		Message:    "Succes Create Student",
		Data:       "",
	}
	utils.EncodeToJSON(response)
	fmt.Println(user_id, mentorId)
}
func fetchAllUser(db *sql.DB) {
	users, err := repository.ReadAllUser(db)
	if err != nil {
		fmt.Println("Update Mentor Error :", err)
		response := model.Response{
			StatusCode: 500,
			Message:    "Internal Server Error",
			Data:       users,
		}
		utils.EncodeToJSON(response)
		return
	}
	response := model.Response{
		StatusCode: 201,
		Message:    "Succes Update Mentor",
		Data:       users,
	}
	utils.EncodeToJSON(response)
	// dataJSON, err := json.MarshalIndent(users, "", " ")
	// if err != nil {
	// 	fmt.Println("Error marshalling JSON:", err)
	// }
	// fmt.Println(string(dataJSON))
}
func editMentor(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("error db begin register :", err)
	}
	mentor, err := repository.UpdateMentor(tx, &model.InputReq.UpdateMentor)
	if err != nil {
		fmt.Println("Update Mentor Error :", err)
		response := model.Response{
			StatusCode: 500,
			Message:    "Internal Server Error",
			Data:       mentor,
		}
		utils.EncodeToJSON(response)
		return
	}
	tx.Commit()
	response := model.Response{
		StatusCode: 201,
		Message:    "Succes Update Mentor",
		Data:       mentor,
	}
	utils.EncodeToJSON(response)
	// fmt.Println(string(dataJSON), "+++++++")
}

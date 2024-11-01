package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	_ "github.com/lib/pq"
	"github.com/train-do/Golang-embed-dan-pengenalan-golang-web/database"
	"github.com/train-do/Golang-embed-dan-pengenalan-golang-web/service"
)

func main() {
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	service.RunningApp(db)
	// service.Register(db)
	log.Println("Starting server on :8080")
}
func init() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

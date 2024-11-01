package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func EncodeToJSON(data interface{}) {
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}

// func DecodeInput() {
// 	file, err := os.Open("input.json")
// 	if err != nil {
// 		// fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

//		decoder := json.NewDecoder(file)
//		if err := decoder.Decode(&model.InputReq); err != nil {
//			// fmt.Println("Error decoding JSON:", err)
//			return
//		}
//	}
func DecodeFromJSON(path string, req interface{}) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Printf("%v\n%v\n", reflect.TypeOf(req).Name(), req)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(req); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

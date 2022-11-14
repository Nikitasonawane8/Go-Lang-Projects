package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	empJsonData := `{"Name":"Nikita Sonawane","Age":22,"Address":"Pune, Maharashtra"}`
	empBytes := []byte(empJsonData)
	var emp Response
	json.Unmarshal(empBytes, &emp)
	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
	fmt.Println(emp.Address)
}

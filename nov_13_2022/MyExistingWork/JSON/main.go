package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string
	Age     int
	Address string
}

func main() {
	emp := Employee{Name: "Nikita Sonawane", Age: 22, Address: "Pune, Maharashtra"}
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(empData))
}

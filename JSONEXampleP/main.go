package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	fmt.Println("JSON Example")

	//creating employee instance
	emp := Employee{
		Name: "Anu",
		Age:  22,
		Dept: "Dev",
	}
	fmt.Println(emp)
	
	//Marshalling -> encoding data into JSON format
	jsonData,err:= json.Marshal(emp)
	//Give func(){} inplace of emp will give a panic
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(jsonData))

	//Unmarshalling -> data from JSON to struct
	var newEmp Employee
	err = json.Unmarshal(jsonData, &newEmp)
	// err = json.Unmarshal([]byte("incorrect json"), &newEmp) // Incorrect JSON string
	if err!=nil{
		panic(err)
	} 	
	fmt.Println(newEmp)
}

type Employee struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Dept string `json:"dept"`
}
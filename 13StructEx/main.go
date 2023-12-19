package main

import "fmt"

func main() {
	fmt.Println("This is an example for structure!")
	// Declare a new person with name and age fields
	person1 := User{"Anu",21,"anu@go.dev",true}
	fmt.Println(person1)
	// Access the fields of the struct using dot notation
	fmt.Println("Name: ", person1.Name)
	fmt.Println("Age : ", person1.Age)
	fmt.Println("Email: ",person1.Email)
	fmt.Println("Status: ", person1.Status)
	// Change the value of one field
	person1.Age = 22
	fmt.Println("After changing Age, here's the updated info:")
	fmt.Printf("Details are: %+v \n",person1) //%+v gives values along with naming conventions


	//For custom type
	pers1 := Person{"Anu",21}
	pers2 := Person{"Ann",22}
	team := Team{"Stars",[]Person{pers1,pers2}}
	fmt.Println("\nTeam details:\n", team)
}
type User struct{
	Name string
	Age int
	Email string
	Status bool
}

//Custom types Ex
type Person struct{
	Name string
	Age int
}
type Team struct{
	Name string
	Members []Person
}
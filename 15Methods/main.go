package main

import "fmt"

func main() {
	fmt.Println("This is an example for structure!")
	// Declare a new person with name and age fields
	person1 := User{"Anu",21,"anu@go.dev",true}
	fmt.Println(person1)
	person1.Age = 22
	fmt.Println("After changing Age, here's the updated info:")
	fmt.Printf("Details are: %+v \n",person1) //%+v gives values along with naming conventions
	fmt.Printf("Name is %v and Email is %v \n",person1.Name,person1.Email)
	person1.GetStatus()
	person1.NewMail()
	fmt.Printf("Name is %v and Email is %v \n",person1.Name,person1.Email)
	person1.AnotherMail()
	fmt.Printf("Name is %v and Email is %v \n",person1.Name,person1.Email)


}
type User struct{
	Name string
	Age int
	Email string
	Status bool
}

func(u User) GetStatus(){
	if u.Status == true {
		fmt.Printf("%v is online.\n",u.Name)
	}else{
		fmt.Printf("%v is offline.\n",u.Name)
	}
}

func(u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("The mail is:",u.Email)
}

func(u *User) AnotherMail(){
	//This is a pointer reciever
	(*u).Email = "test1@go.dev"
	fmt.Println("The another mail is:", (*u).Email)
}
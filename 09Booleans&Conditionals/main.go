package main

import "fmt"

func main(){
	fmt.Println("Booleans and Conditionals")
	fmt.Println("Enter your age: ")
	var age int
	fmt.Scanln(&age)
	if age<0 || age>100{
		fmt.Println("Invalid age")
	}else if age < 18{
		fmt.Println("You're not eligible to vote!")
	}else{
		fmt.Println("You can Vote")
	}

	names := []string{"ann","amy","anu","ash","ari"}
	for i,v := range names{
		if i == 1{
			fmt.Println("Continuing at pos",i)
		}else if i == 3{
			fmt.Printf("Breaking at %v \n",i)
			break
		}else{
			fmt.Printf("The value at %v is %v \n",i,v)
		}
	}

	rougueValue:=1

	for rougueValue < 5{
		if rougueValue == 3{
			goto lco
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

	lco : 
	fmt.Println("Jumped here!")
}
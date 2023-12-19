package main 

import "fmt"

func main(){
	defer fmt.Println("World")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i :=1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
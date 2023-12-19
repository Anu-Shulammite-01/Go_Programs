package main

import "fmt"

func main(){
	fmt.Println("Loops")

	x:=0
	for x<=5{
		fmt.Println("The value of x is:",x)
		x++
	}

	for i:=0;i<3;i++{
		fmt.Println("Hi",i)
	}

	names := []string{"Anu","Mehraj","Harshini","Sravani"}
	for index,value := range names{
		fmt.Println("The position at index:",index)
		fmt.Println("The value is",value)
	} 
	for i := range names {		
		if i == 2{
			break
		}
		fmt.Println("The index is at:",i)
	}
	for _,v := range names{
		fmt.Println("The value is:",v)
	}
}
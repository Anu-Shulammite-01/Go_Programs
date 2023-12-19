package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

func main(){
	fmt.Println("This is User Input")
	var i,j int
	fmt.Println("Enter two numbers: ")
	fmt.Scanln(&i,&j)
	if(i>j){
		fmt.Printf("%d is greater than %d\n",i,j)
	}else{
		fmt.Printf("%d is lesser than %d\n",i,j)
	}

	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a string:")
	str,_ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n") //Removes the extra line

	fmt.Println("You entered:", str)

	//save formatted strings
	str2 := fmt.Sprintf("The inputted string is %v",str)
	fmt.Println(str2)
}
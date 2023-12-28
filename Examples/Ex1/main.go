package main   

import "fmt"

func main(){
	var list = []string{}
	fmt.Println("Enter the elements into the list")
	var input string
	for {
		fmt.Scan(&input)
		if input == "no"{
			break
		}else{
			list = append(list, input)
		}
	}
	fmt.Printf("\nThe entered list is %v\n", list)
}
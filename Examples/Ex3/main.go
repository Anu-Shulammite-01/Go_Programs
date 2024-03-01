package main

import "fmt"

func reverseInt(n int) int{
	newInt:=0

	for n>0{
		remainder:=n%10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}

func main(){
	fmt.Println("Enter  a number: ")
	var num int
	fmt.Scan(&num)
	reversedNum := reverseInt(num)
	if num==reversedNum {
		fmt.Println("It's a palindrome")
	}else{
		fmt.Println("Not a palindrome")
	}
}
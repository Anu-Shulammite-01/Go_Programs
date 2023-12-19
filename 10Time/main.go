package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Time Handling")
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("01-02-2006"))

	createDate := time.Date(2001,time.December,9,6,20,0,0,time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
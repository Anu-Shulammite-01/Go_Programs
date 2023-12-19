package main

import (
	"fmt"
	// "log"
)

func main(){
	defer func(){
		r := recover();
		if r!= nil{
			fmt.Println("Recovered from", r)
		}
	}()

	panic("Oops!") // This is a panic call
	// defer functions will execute

	// log.Fatalln("Panicking")
	// In Fatalln, no defer functions will execute
	
}
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
	Execute()

}
func runner1(){
	defer wg.Done()
	defer fmt.Println("Done")
	for i := 0; i < 4; i++ {
		fmt.Println("I am the 1st Runner")	
	}
}
func runner2(){
	defer wg.Done()
	for i := 0; i < 4; i++ {
		fmt.Println("I am the 2nd Runner")	
	}
}
func Execute(){
	//Adding two tasks to wait group, one for each goroutine
	wg.Add(2) //two goroutines are about to launch
	go runner1()
	go runner2()
	wg.Wait()
}
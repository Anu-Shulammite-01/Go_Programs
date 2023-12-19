package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Channel Example")
	// Create a channel of type int with buffer size 3
	ch := make(chan int,2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int,wg *sync.WaitGroup){
		defer wg.Done()
		val,isChannelOpen := <-ch
		fmt.Println(isChannelOpen)
		fmt.Println(val)
	}(ch,wg)
	go func(ch chan int,wg *sync.WaitGroup){
		defer wg.Done()
		ch <- 0
		//ch <- 6
		close(ch) 
	}(ch,wg)

	wg.Wait()
}
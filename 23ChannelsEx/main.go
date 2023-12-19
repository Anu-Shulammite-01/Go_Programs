package main 

import "fmt"

func main(){
	fmt.Println("Channels in Golang")

	myChannel := make(chan int)
	fmt.Println(myChannel)
	go myFunc(myChannel)
	myChannel <- 23

	//buffered channel
	myChannel1 := make(chan string,2)
	myChannel1 <- "Anu"
	myChannel1 <- "Shulammite"
	fmt.Println(<-myChannel1)
	fmt.Println(<-myChannel1)
	
	fmt.Println("End")
}
// Function to receive data from the channel
func myFunc(ch chan int){
	data := <-ch
	fmt.Println(data+34)
}
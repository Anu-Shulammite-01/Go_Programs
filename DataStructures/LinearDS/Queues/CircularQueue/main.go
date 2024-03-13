package main 

import "fmt"

type CircularQueue struct{
	items []int
	head int
	tail int
}

func(c *CircularQueue) Enqueue(i int){
	if (len(c.items)+1-c.head == c.tail) {
		fmt.Println( "Error: Queue is Full")
		return;
	}
	c.items = append(c.items, i)
	c.tail++
}
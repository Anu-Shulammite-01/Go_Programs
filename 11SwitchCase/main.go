package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	diceNum := rand.Intn(6) + 1
	println("You rolled a", diceNum, "!")

	switch diceNum {
	case 1:
		fmt.Println("You can enter!")
	case 2:
		fmt.Println("You can move 2 spots.")
	case 3:
		fmt.Println("You can move 3 spots.")
	case 4:
		fmt.Println("You can move 4 spots.")
	case 5:
		fmt.Println("You can move 5 spots.")
	case 6:
		fmt.Println("Woahh! You can move 6 spots and can roll the dice again.")
	default:
		fmt.Println("What was that??")
	}
}
package main

import (
	"fmt"
)

func main() {
	menu := map[int]string{
		1: "Pizza",
		2: "Sandwich",
	}
	fmt.Println(menu)
	// Accessing elements in a Map
	fmt.Println("Which item you wanna check 1 or 2?")
	var m int
	fmt.Scanln(&m)
	item, ok := menu[m]
	if ok {
		fmt.Printf("Item %s is available.\n", item)
	} else {
		fmt.Println("This item is not available.")
	}
	// Adding new items to the Map
	menu[3] = "Salad"
	fmt.Println(menu)
	// Removing an Item from the Map
	delete(menu, 2)
	fmt.Println(menu)

	//Another method for maps
	languages := make(map[string]string)
	languages["J"] = "Java"
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	for key, value := range languages {
		fmt.Println(key + ": " + value)
	}
	for _,value := range languages{
		fmt.Println(value)
	}
	for key:=range languages{
		fmt.Println(key)
	}	
}
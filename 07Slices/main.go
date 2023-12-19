package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	fmt.Println("Slices!")

	var vegetables = []string{"Potato","Capsicum","Carrot"}
	fmt.Println(vegetables)
	vegetables = append(vegetables,"Mushroom","Brinjal")
	fmt.Println(vegetables)

	//sorting
	sort.Strings(vegetables)
	fmt.Println(vegetables)

	//removing a value
	var index int = 2
	vegetables = append(vegetables[:index],vegetables[index+1:]... ) //"..." operator passes the elements of vegetables[index+1:]as seperate elements for append
	fmt.Println(vegetables)

	var str string = "Anu"
	fmt.Println(strings.Split(str, ""))
}
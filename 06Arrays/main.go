package main

import "fmt"

func main(){
	fmt.Println("Arrays")

	var fruits[3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Pomegranate"

	fmt.Println(fruits)
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	var vegetables = [3]string{"Potato","Capsicum","Carrot"}
	fmt.Println(vegetables)
	fmt.Println(vegetables[1:2])

	//copy of array into another array using value
	languages := [3]string{"Java","Go","Python"}
	copyArr := languages
	fmt.Println("Array: ",copyArr)
	languages[0] = "Flutter"
	fmt.Println(languages)
	fmt.Println(copyArr) //copy array is not updated

	//copy of array into another array using reference
	languages1 := [3]string{"Scalar","Go","Ruby"}
	copyArr1 := &languages1
	fmt.Println("Array: ",*copyArr1)
	languages1[0] = "JS"
	fmt.Println(languages1)
	fmt.Println(*copyArr1)
}
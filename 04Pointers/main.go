package main

import "fmt"

func main(){
	fmt.Println("These are pointers")
	name := "anu"
	fmt.Println("The memory location of name is:",&name)
	p := &name
	fmt.Println("The memory location is:",p)
	fmt.Println("The value of the memory location is:",*p)

	var myNum int = 2
	ptr := &myNum
	fmt.Println("The number location is:",ptr)
	fmt.Println("The number is:",*ptr)
	fmt.Println("Now we will change the value of the number")
	*ptr = 3
	fmt.Println("The number is now:",*ptr)
	*ptr =  *ptr * 2
	fmt.Println(*ptr)

	//passing pointers to functions
	add(&myNum,10)
	fmt.Println(myNum,"is added with 10.")

	var n1 string = "Anu"
	UName(&n1)
	fmt.Println("Updated User Name is : ",n1)

}
func add(n *int,m int){
	*n += m
}
func UName(s *string){
	*s = "Hello " + *s
}
package main 

import "fmt"

func main(){
	fmt.Println("This is main function")
	greet()

	res := adder(3,4)
	fmt.Printf("The result of adding 3 and 4 is: %d\n", res)//%d represents corresponding argument must be an integer!

	proRes,msg := proAdder(1,3,56,7)
	fmt.Println("Pro Result is: ", proRes)
	fmt.Println("Message:", msg)

	//Anonymous Functions
	func(){
		fmt.Println("Byee")
	}()

	value1 := func(ele string){
		fmt.Println("See you again",ele)
	}
	value1("Anu")


}
func init(){
	fmt.Println("Hello!")
	//Always executed before main and in the order they were created
}

func init(){
	fmt.Println("Init function called")
	//It is used to initialize the global variables
}

func adder(num1,num2 int) int{
	return num1 + num2
}

func proAdder(values ...int)(int,string){
	//... are the variadic functions
	sum := 0
	for _,v := range values{
		sum += v
	}
	return sum,"Done"
}

func greet(){
	fmt.Println("Good Morning!")
}
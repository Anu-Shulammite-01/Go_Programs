package main

import "fmt"
const Var11 = "This is Go" // Var11 is public coz it started with Capital letter

func main()  {
	fmt.Println("Started")
	var num1 int = 10
	fmt.Println(num1)
	num2:=15
	fmt.Println(num2)
	fmt.Println("The Sum is: ",num1+num2)
	fmt.Printf("The type is: %T \n",num2)

	var str1 string = "Anu"
	fmt.Println(str1)
	fmt.Printf("The type is: %T \n",str1)

	var boolean bool = true
	fmt.Println(boolean)
	fmt.Printf("The type is: %T \n",boolean)

	var unit uint8 = 255
	fmt.Println(unit)
	fmt.Printf("The type is: %T \n",unit)

	var fl1 float32 = 245.898937998435
	fmt.Println(fl1)
	fmt.Printf("The type is: %T \n",fl1)

	var fl2 float64 = 245.898937998435
	fmt.Println(fl2)
	fmt.Printf("The type is: %T \n",fl2)

	//default value
	var variable int
	fmt.Println(variable)
	fmt.Printf("The type is: %T \n",variable)

	//no type
	var var1 = "Anu"
	fmt.Println(var1)
	fmt.Printf("The type is: %T \n",var1)

	//constant
	fmt.Println(Var11)

	//multiple declaration
	var a,b,c int = 1,2,3
	var d = a*b*c
	fmt.Println(d)

	var x,y = 6,"Hello"
	z,q := 7,"World"

	fmt.Println(x,z,"GOLang")
	fmt.Println(y,q)
	fmt.Println(3,4,"Goo")


}
package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("This is a program on conversions")
	var str1 string = "23"
	fmt.Printf("The type of str1 is %T \n",str1)
	num,_ := strconv.ParseInt(str1,8,64)
	fmt.Printf("The type of str1 is %T \n",num)
	num1,_ := strconv.Atoi("34") // used for converting string to int
	fmt.Printf("The type of 34 is %T \n",num1)
	str2:= strconv.Itoa(234) // used for converting int to string
	fmt.Printf("The type of 234 is %T \n",str2)
	num2,_ := strconv.ParseFloat(str1,64)
	fmt.Printf("The type of str1 is %T \n",num2)
	num3,_ := strconv.ParseBool("true")
	fmt.Printf("The type of str1 is %T \n",num3)

	q := strconv.Quote("Hello, 世界") //returns go string literals
	fmt.Printf("The type is %T \n",q)

}
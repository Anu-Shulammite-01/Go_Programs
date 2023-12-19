package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("Welcome to Files!")
	content:="This is the content inside the files.."

	file,err:=os.Create("./FileEx.txt")
	// if err != nil{
	// 	panic(err)
	// }
	checkNilError(err)

	length,err:=io.WriteString(file,content)
	checkNilError(err)
	
	fmt.Println("The length is:",length)
	defer file.Close()
	readFile("./FileEx.txt")
}
func readFile(fileName string){
	databyte,err:=os.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("The read data is:\n",string(databyte))
}
func checkNilError(err error){
	if err != nil{
		panic(err)
	}
}
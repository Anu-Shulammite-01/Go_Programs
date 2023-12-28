package main

import (
	"fmt"
	"sort"
)

type students struct{
	name string
	age int
}

func main(){
	var stuDetails = []students{}
	stuDetails = append(stuDetails, students{"John", 21})
	stuDetails = append(stuDetails, students{"Jane", 19})
	fmt.Println("Do you wanna add the details of students?(yes/no):")
	var input string 
	fmt.Scanln(&input)
	for input != "no"{
		fmt.Println("Enter the details of students:")
		var name string
		var age int
		fmt.Printf("Name : ")
		fmt.Scanln(&name)
		fmt.Printf("Age : ")
		fmt.Scanln(&age)
		stuDetails = append(stuDetails, students{name, age})
		fmt.Println("\nDo you want to continue? (yes/no)")
		fmt.Scanln(&input)
	}
	fmt.Println(stuDetails)
	sort.Slice(stuDetails, func(i, j int) bool {
		return stuDetails[i].age < stuDetails[j].age
	})
	fmt.Println("\nSorted Details by Age are as follows:")
	for _, value := range stuDetails{
		fmt.Println(value.name,"\t",value.age)
	}
	
}

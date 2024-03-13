package main

import "fmt"

type Stack struct{
	item []int
}

func(s *Stack) Push(i int)bool{
	s.item = append(s.item,i)
	return true
}

func(s *Stack)  Pop() (int,bool){
	l := len(s.item)-1
	if l < 0 {
		return 0, false
	}else{
		pop := s.item[l]
		s.item  = s.item[:l]
		return pop, true
	}
}

func(s *Stack)  Peek() (int, bool){
	l := len(s.item) - 1
	if l == -1 {
		return 0 ,false
	}else{
		return  s.item[l],true
	}
}

func(s *Stack) Display(){
	fmt.Println(s.item)
}

func main(){
	stack := &Stack{}
	fmt.Println("Do you wanna access stack?(Yes/No): ")
	var yn string
	fmt.Scanln(&yn)
	for yn!= "No"{
		fmt.Println("The following are the choices:")
		fmt.Println("1.Push")
		fmt.Println("2.Pop")
		fmt.Println("3.Peek")
		fmt.Println("4.Display")
		var c int
		fmt.Println("Enter your choice : ")
		fmt.Scanln(&c)
		switch c{
		case 1:
			var num1 int
			fmt.Println("Enter a number to push into Stack: ")
			fmt.Scanln(&num1)
			stack.Push(num1)
		case 2:
			v, b:= stack.Pop()
			if b==true{
				fmt.Println("Popped element is ", v)
			}else{
				fmt.Println("Stack is empty.")
			}
		case 3:
			v,b:= stack.Peek()
			if b==true{
				fmt.Println("Topmost element is ",v)
			}else{
				fmt.Println("Stack is empty.")
			}
			
		case 4:
			stack.Display()
			
		default:
			fmt.Println("Invalid Choice! Try Again.")
				
		}
		fmt.Println("\nDo you want to continue? (Yes/No): ")
		fmt.Scanln(&yn)

	}
}
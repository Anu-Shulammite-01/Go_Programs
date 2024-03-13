package main

import "fmt"

type Queue struct{
	items []int
}

func(q *Queue)  Enqueue(i int){
	q.items = append(q.items,i)
}

func(q *Queue) Dequeue() (int,bool){
	if len(q.items) == 0 {
		fmt.Println( "The queue is empty")
		return -1 , false
	}else{
		pop := q.items[0]
		q.items = q.items[1:]
		return pop, true
	}
}

func(q *Queue) Peek() (int, bool){
	if len(q.items) == 0{
		return 0 , false
	}else{
		return  q.items[0],true
	}
}

func(q *Queue) Display(){
	fmt.Println(q.items)
}

func main(){
	queue := &Queue{}
	fmt.Println("Do you wanna access queue?(Yes/No): ")
	var yn string
	fmt.Scanln(&yn)
	for yn!= "No"{
		fmt.Println("The following are the choices:")
		fmt.Println("1.Enqueue")
		fmt.Println("2.Dequeue")
		fmt.Println("3.Peek")
		fmt.Println("4.Display")
		var c int
		fmt.Println("Enter your choice : ")
		fmt.Scanln(&c)
		switch c{
		case 1:
			var num1 int
			fmt.Println("Enter a number to push into Queue: ")
			fmt.Scanln(&num1)
			queue.Enqueue(num1)
		case 2:
			v, b:= queue.Dequeue()
			if b==true{
				fmt.Println("Dequeued element is ", v)
			}else{
				fmt.Println("Queue is empty.")
			}
		case 3:
			v,b:= queue.Peek()
			if b==true{
				fmt.Println("Topmost element is ",v)
			}else{
				fmt.Println("Queue is empty.")
			}
			
		case 4:
			queue.Display()
			
		default:
			fmt.Println("Invalid Choice! Try Again.")
				
		}
		fmt.Println("\nDo you want to continue? (Yes/No): ")
		fmt.Scanln(&yn)

	}
}
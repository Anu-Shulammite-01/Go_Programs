package main 

import "fmt"

//implementing priority queue using Array

type PriorityQueue struct{
	items []Item	
}
type Item struct {
	value int
	priority int
}

func(p *PriorityQueue) Enqueue(i Item){
	p.items =  append(p.items, i)
}

func(p  *PriorityQueue) Dequeue()(Item , bool){
	if len(p.items) == 0{
		return  Item{}, false
	}
	
	highestPriorityIndex := 0
	for i,item := range p.items{
		if item.priority > p.items[highestPriorityIndex].priority{
			highestPriorityIndex = i
		}
	}
	highestPriorityItem := p.items[highestPriorityIndex]
	p.items = append(p.items[:highestPriorityIndex],p.items[highestPriorityIndex+1:]...)
	return highestPriorityItem, true
}

func(p *PriorityQueue) Peek()(int,bool){
	if !(len(p.items)>0){
		return 0, false
	}
	highestPriorityIndex := 0
	for i, item := range p.items{
		if item.priority > p.items[highestPriorityIndex].priority{
			highestPriorityIndex = i
		}
	}
	highestPriorityItem := p.items[highestPriorityIndex]
	return highestPriorityItem.value,true
}

func(p *PriorityQueue) Display(){
	fmt.Println(p.items)
}

func main(){
	q := &PriorityQueue{}
	fmt.Println("Do you wanna access Priority queue?(Yes/No): ")
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
			var  pri1 int
			fmt.Println("Enter a number to push into PQueue: ")
			fmt.Scanln(&num1)
			fmt.Println("Enter priority for that number: ")
			fmt.Scanln(&pri1)
			q.Enqueue(Item{num1,pri1})
		case 2:
			v, b:= q.Dequeue()
			if b==true{
				fmt.Println("Dequeued element is ", v)
			}else{
				fmt.Println("PQueue is empty.")
			}
		case 3:
			v,b:= q.Peek()
			if b==true{
				fmt.Println("Topmost element is ",v)
			}else{
				fmt.Println("PQueue is empty.")
			}
			
		case 4:
			q.Display()
			
		default:
			fmt.Println("Invalid Choice! Try Again.")
				
		}
		fmt.Println("\nDo you want to continue? (Yes/No): ")
		fmt.Scanln(&yn)
	}

}

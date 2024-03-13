package main

import "fmt"

type Deque struct{
	items []int
}

func (d *Deque) PushFront(i int){
	if len(d.items) == 0{
		d.items = append(d.items, i)
	}else{
		d.items = append([]int {i},d.items...)
	}
}

func(d *Deque) PushBack(i int){
	d.items = append(d.items, i)
}

func(d *Deque) PopFront() (int,bool){
	if len(d.items)==0{
		return  0 , false
	}
	pop := d.items[0]
	d.items  = d.items[1:]
	return pop,true
}

func(d *Deque) PopBack() (int,bool){
	l:=len(d.items)-1
	if l < 0{
		return 0 ,false
	}

	pop :=  d.items[l]
	d.items = d.items[:l]
	return pop, true
}

func(d *Deque) Front()int{
	if len(d.items)==0{
		fmt.Println( "Empty deque")
		return 0
	}
	return d.items[0]
}

func(d *Deque) Rear()int{
	l:=len(d.items)-1
	if l == -1{
		fmt.Println( "Empty deque")
		return 0
	}
	return d.items[len(d.items)-1]
}

func(d *Deque) Display(){
	if len(d.items) == 0{
		fmt.Println("The deque is empty")
	}
	fmt.Println(d.items)
}

func main(){
	deque := &Deque{}
	fmt.Println("Do you wanna access Deque?(Yes/No): ")
	var yn string
	fmt.Scanln(&yn)
	for yn!= "No"{
		fmt.Println("The following are the choices:")
		fmt.Println("1.PushFront")
		fmt.Println("2.PushBack")
		fmt.Println("3.PopFront")
		fmt.Println("4.PopBack")
		fmt.Println("5.Front")
		fmt.Println("6.Rear")
		fmt.Println("7.Display")
		var c int
		fmt.Println("Enter your choice : ")
		fmt.Scanln(&c)
		switch c{
		case 1:
			var num1 int
			fmt.Println("Enter a number to push into Deque at Front: ")
			fmt.Scanln(&num1)
			deque.PushFront(num1)
		case 2:
			var num2 int
			fmt.Println("Enter a number to push into Deque at Back: ")
			fmt.Scanln(&num2)
			deque.PushBack(num2)
		case 3:
			v,b:=deque.PopFront()
			if b==true {
				fmt.Println("Popped element from Front = ", v)
			}else{
				fmt.Println("The Deque is Empty.")
			}
		case 4:
			v,b:=deque.PopBack()
			if b==true {
				fmt.Println("Popped element from Rear = ", v)
			}else{
				fmt.Println( "The Deque is Empty.")
			}
		case 5:
			v :=  deque.Front()
			fmt.Println("Element at Front = ", v)
		case 6:
			v := deque.Rear()
			fmt.Println( "Element at Rear = ", v)
		case 7:
			deque.Display()
		default:
			fmt.Println("Invalid Choice! Try Again.")
				
		}
		fmt.Println("\nDo you want to continue? (Yes/No): ")
		fmt.Scanln(&yn)
	}
}
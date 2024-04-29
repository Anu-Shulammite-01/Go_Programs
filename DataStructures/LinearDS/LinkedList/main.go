package main 

import "fmt"

type node struct{
	data int
	next *node
}

type linkedList struct{
	head *node
	length int
}

func(l *linkedList) Prepend(n *node){
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func(l *linkedList) Postpend(n *node){
	current := l.head
	for current.next != nil{
		current = current.next
	}
	current.next = n
	l.length++
}

func(l *linkedList) AppendAt(n *node, position int){
	if position < 0 || position >l.length{
		fmt.Println("Out of bounds")
	}else if position == 0{
		l.Postpend(n)
		l.length++
	}else if position == l.length{
		l.Prepend(n)
		l.length++
	}else{
		current:=l.head
		for i:=0;i<position-1;i++{
			current = current.next
		}
		n.next = current.next
		current.next = n
		l.length++
	}
	
}

func(l *linkedList) Delete(value int){
	if l.length == 0{
		fmt.Println("List is empty")
		return
	}
	if l.head.data == value{
		l.head = l.head.next
		l.length--
		return
	}
	current := l.head
	for current.next.data != value{
		if current.next.next == nil{
			fmt.Println("Value not found")
			return
		}
		current = current.next
	}
	current.next = current.next.next
}

func(l *linkedList) Sort(){
	if l.length == 0{
		fmt.Println("List is empty")
		return
	}
	for i:=l.head;i!=nil;i=i.next{
		for j:=i.next;j!=nil;j=j.next{
			if i.data > j.data{
				temp := i.data
				i.data = j.data
				j.data = temp
			}
		}
	}
}

func(l *linkedList) PrintData(){
	current:=l.head
	for current != nil{
		fmt.Printf("%d ",current.data)
		current = current.next
	}
	fmt.Println()
}

func main(){
	mlist := linkedList{}
	node1 := &node{data: 12}
	node2 := &node{data: 13}
	node3 := &node{data: 14}
	node4 := &node{data: 11}
	node5 := &node{data: 19}
	mlist.Prepend(node1)
	mlist.Prepend(node2)
	mlist.Prepend(node3)
	mlist.Postpend(node4)
	mlist.AppendAt(node5,2)

	fmt.Println(mlist)
	mlist.PrintData()
	mlist.Delete(100)
	mlist.PrintData()
	mlist.Delete(19)
	mlist.PrintData()
	mlist.Sort()
	mlist.PrintData()
}

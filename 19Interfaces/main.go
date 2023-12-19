package main 

import "fmt"

type Shape interface{
	Area() float64
	Circum() float64
}
type Circle struct{
	Radius float64
}
type Square struct{
	Side float64
}
func (c *Circle) Area()float64{
	return 3.14 * c.Radius * c.Radius
}
func (s *Square) Area()float64{
	return s.Side * s.Side
}
func (s *Square) Circum()float64{
	return 4 * s.Side
}
func (c *Circle) Circum()float64{
	return 2 * 3.14 * c.Radius
}

func PrintShapeInfo(s Shape){
	fmt.Printf("The area of the shape is: %v\n", s.Area())
	fmt.Printf("The circumference of the shape is: %v\n", s.Circum())
}

func main(){
	var circle = &Circle{5} 
	var square = &Square{4.6}
	//Printing information about a circle
	fmt.Println("\nInformation about a circle")
	PrintShapeInfo(circle)
	//Printing information about a square
	fmt.Println("\nInformation about a square")
	PrintShapeInfo(square)

}


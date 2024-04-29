package main

import (
	"fmt"
	"math"
	"strings"
)

func main(){
	fmt.Println("Hey")

	fmt.Println(math.Ceil(56.3456888))
	fmt.Println(math.Floor(56.3456888))
	fmt.Println(math.Round(56.3456888))
	fmt.Println(math.RoundToEven(56.3456888))
	fmt.Println(math.Min(3,4))
	fmt.Println(math.Max(7,9))
	fmt.Println(math.Pow(2,2))
	fmt.Println(math.Sqrt(100))
	fmt.Println(math.Mod(10,2))
	fmt.Println(complex(1,2))

	var sb strings.Builder
	fmt.Println(sb.String())
	sb.WriteString("你好 世界")
	fmt.Println(sb.String())
	fmt.Println(sb.Len())
	fmt.Println(sb.Cap())

	//constants
	const(
		a = 1
		b          // takes the previous constant value
		c = 3
		d
	)
	fmt.Println(a, b, c, d)

	const(
		x = iota     // starts counting from zero by default
		y
		z
	)
	fmt.Println(x, y, z)


	fmt.Println("Label Usage")
	//label in for loop
	L:
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if j == 3 {
				break L
			}
			fmt.Printf("%d %d", i,j)
		}
		fmt.Println()
	}

}



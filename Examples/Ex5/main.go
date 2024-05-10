package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scanln(&n)
	if n > 1{
		for i := 2; i < int(n); i++ {
			for j := -10; j < int(n); j++ {
				temp := math.Pow(float64(i), float64(j))
				if temp == n  {
					fmt.Printf("x^y is %d,%d\n", i, j)
					return
				} else if temp > n {
					break 
				}
			}
		}
	} else {
		for i := 0.1; i < 10; i += 0.1 {
			for j := -10.0; j < 10.0; j += 0.1 {
				temp := math.Pow(i, j)
				if math.Abs(temp-n) < 0.0001 {
					fmt.Printf("x^y is %.1f,%.1f\n", i, j)
					return
				}
			}
		}
	}
	
	fmt.Println("No valid x^y found for the given n.")
}

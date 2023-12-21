package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var count = 0

func Update(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("Update canceled at", id)
	default:
		fmt.Println("Updated at", id)
		count++
	}
}

func main() {
	// Create a new context with cancelation
	ctx := context.Background()
	var wg sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		wg.Add(1)
		go Update(ctx, i, &wg)
		defer cancel()
	}
	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("The total count is:",count)
}

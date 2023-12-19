package main

import (
	"fmt"
	"sync"
)

// var wg sync.WaitGroup
// var mut sync.Mutex
var count = 0

func main(){
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go counter(&wg,&m)
	}
	wg.Wait()
	fmt.Println("Count:",count)
}

//wait groups and mutex are passed as pointers
func counter(wg *sync.WaitGroup,mut *sync.Mutex){
	mut.Lock()
	count++
	mut.Unlock()
	defer wg.Done()
}
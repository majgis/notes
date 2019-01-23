package main

import (
	"fmt"
	"sync"
)

// Example of a data race
func main() {
	var wg sync.WaitGroup
	value := 0
	wg.Add(1)
	go func(){
		defer wg.Done()
		value++
		fmt.Printf("inside value = %v\n", value)
	}()
	// wg.Wait()
	fmt.Printf("outside value = %v\n", value)
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	//create the wait group
	var wg sync.WaitGroup

	wg.Add(2) //increase the counter by 2

	go func() {
		count("sheep")
		wg.Done() //decrease the counter parameter in Add(int) by 1
	}()

	go func() {
		count("cow")
		wg.Done()
	}()
	wg.Wait() //this method blocks the execution of code until the internal counter reduces to a 0 value
}

func count(thing string) {
	for i := 0; i <= 3; i++ {
		fmt.Println(thing)
	}
}

// declare the two goroutines function with the main func

package main

import (
	"fmt"
	"sync"
)

var chanMutex = make(chan int, 5) // Create channel with buffer size == 1

func main() {
	var iCounter int
	var waitGroup sync.WaitGroup
	const iIterations = int(1000)

	waitGroup.Add(iIterations)

	for iIndex := 0; iIndex < iIterations; iIndex++ {
		go func() {
			handler(&iCounter)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Printf("iCounter: %+8v\n", iCounter)
}

func handler(pCounter *int) {
	// Set value in channel. So channel's buffer is full. Other Goroutines will wait when the buffer will free.
	chanMutex <- 1 // Other Goroutines will wait here.
	*pCounter++
	fmt.Println(*pCounter)
	// Reed value from channel. So channel's buffer is free. Other Goroutines can write in buffer new.
	<-chanMutex
}

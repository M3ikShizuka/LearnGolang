package main

import (
	"fmt"
)

const countOfGorutines = 5 // count of call go handler

var chanWaitGroup = make(chan int, countOfGorutines)

func main() {
	go handler(1)
	go handler(2)
	go handler(3)
	go handler(4)
	go handler(5)

	for i := 0; i < countOfGorutines; i++ { // Read from channel data sent from every "go handler"
		<-chanWaitGroup
	}

	fmt.Printf("Done!\n")
}

func handler(iGoroutineIndex int) {
	fmt.Printf("iGoroutineIndex: %v\n", iGoroutineIndex)
	chanWaitGroup <- 1
}

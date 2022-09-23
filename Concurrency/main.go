package main

import (
	"fmt"
	"sync"
)

type Counters struct {
	iVal int
}

var mutex sync.Mutex
var iCounter int
var sliceData []int = []int{1, 2, 3}
var mapData map[string]int = map[string]int{"i": 1}

func main() {
	var pCounter = new(Counters)
	var waitGroup sync.WaitGroup
	const iIterations = int(1000)
	waitGroup.Add(iIterations)
	for iIndex := 0; iIndex < iIterations; iIndex++ {
		go func() {
			ProcessData(pCounter)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Printf("pCounter: %+8v\n", pCounter)
	fmt.Printf("iCounter: %+8v\n", iCounter)
	fmt.Printf("sliceData[0]: %+8v\n", sliceData[0])
}

func ProcessData(pCounter *Counters) {
	// mutex.Lock()
	pCounter.iVal++
	iCounter++
	sliceData[0] = sliceData[0] + sliceData[1] + sliceData[2]
	mapData["i"]++
	// mutex.Unlock()
}

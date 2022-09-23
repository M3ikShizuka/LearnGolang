package main

import "sync"

func main() {
	// Create a large heap allocation of 10 GiB
	const sliceSize = 10 << 30
	ballast := make([]byte, sliceSize)
	_ = ballast
	for iIndex := 0; iIndex < sliceSize; iIndex++ {
		ballast[iIndex] = byte(iIndex)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

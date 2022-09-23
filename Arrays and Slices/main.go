package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func changeSlice(slice []int) {
	slice = append(slice, 31, 32, 33, 34, 35)
	fmt.Printf("changeSlice(): slice (addr: %p): %+8v\n",
		&slice, *(*reflect.SliceHeader)(unsafe.Pointer(&slice)))
}

func changeSlicePointer(pSlice *[]int) {
	*pSlice = append(*pSlice, 31, 32, 33, 34, 35)
	fmt.Printf("changeSlicePointer(): slice (addr: %p): %+8v\n",
		pSlice, *(*reflect.SliceHeader)(unsafe.Pointer(pSlice)))
}

func main() {
	arrayOneDimensional := [5]int{1, 2, 3, 4, 5}
	arrayTwoDimensional := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	sliceArrayOneDimensional := arrayOneDimensional[:]
	sliceArrayTwoDimensional := arrayTwoDimensional[:]
	mySlice := []int{1, 2, 3}

	fmt.Printf(
		"arrayOneDimensional: %v\n"+
			"arrayTwoDimensional: %v\n"+
			"sliceArrayOneDimensional: %v\n"+
			"sliceArrayTwoDimensional: %v\n"+
			"mySlice: %v\n",
		arrayOneDimensional,
		arrayTwoDimensional,
		sliceArrayOneDimensional,
		sliceArrayTwoDimensional,
		mySlice,
	)

	var sliceNil []int            // Empty slice: nil slice (prefer declaration)
	sliceEmpty1 := []int{}        // Empty slice: zero-length slise
	sliceEmpty2 := make([]int, 0) // Empty slice: zero-length slise

	// sliceNil = append(sliceNil, 1, 2, 3)

	fmt.Printf("sliceNil (addr: %p): %+8v\n",
		&sliceNil, *(*reflect.SliceHeader)(unsafe.Pointer(&sliceNil)))
	fmt.Printf("sliceEmpty1 (addr: %p): %+8v\n",
		&sliceEmpty1, *(*reflect.SliceHeader)(unsafe.Pointer(&sliceEmpty1)))
	fmt.Printf("sliceEmpty2 (addr: %p): %+8v\n",
		&sliceEmpty2, *(*reflect.SliceHeader)(unsafe.Pointer(&sliceEmpty2)))

	sliceEmpty1 = append(sliceEmpty1, 1)
	sliceEmpty2 = append(sliceEmpty2, 2, 3, 4, 5)

	fmt.Printf("sliceEmpty1 after append (addr: %p): %+8v\n",
		&sliceEmpty1, *(*reflect.SliceHeader)(unsafe.Pointer(&sliceEmpty1)))
	fmt.Printf("sliceEmpty2 after append (addr: %p): %+8v\n",
		&sliceEmpty2, *(*reflect.SliceHeader)(unsafe.Pointer(&sliceEmpty2)))

	sliceEmpty2 = sliceEmpty2[1:2]

	fmt.Printf("sliceEmpty2 change slice (addr: %p): %+8v\n",
		&sliceEmpty2, *(*reflect.SliceHeader)(unsafe.Pointer(&sliceEmpty2)))

	sliceChange := make([]int, 5)

	fmt.Printf("main(): sliceChange before changeSlice() (addr: %p): %+8v\n",
		&sliceChange, *(*reflect.SliceHeader)(unsafe.Pointer(&sliceChange)))

	changeSlice(sliceChange)

	fmt.Printf("main(): sliceChange after changeSlice() (addr: %p): %+8v\n",
		&sliceChange, *(*reflect.SliceHeader)(unsafe.Pointer(&sliceChange)))

	changeSlicePointer(&sliceChange)

	fmt.Printf("main(): sliceChange after changeSlicePointer() (addr: %p): %+8v\n",
		&sliceChange, *(*reflect.SliceHeader)(unsafe.Pointer(&sliceChange)))
}

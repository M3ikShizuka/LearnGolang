package main

func main() {
	myMap := make(map[string]int)
	myMap["MyNumber"] = 31
	varA := myMap["MyNumber"]
	_ = varA
	// pMyValMap := (*runtime.hmap)(unsafe.Pointer(&myMap))
	// _ = pMyValMap
}

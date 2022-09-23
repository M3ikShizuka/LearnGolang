package main

func foo(i int) int {
	if i < 1e8 {
		return foo(i + 1)
	}
	return -1
}

func main() {
	foo(0)
}

package main

func main() {
	//// A send to a nil channel blocks forever
	// var c chan string = make(chan string)
	// c <- "let's get started"

	//// A receive from a nil channel blocks forever
	// var c chan string = make(chan string)
	// str := <-c
	// _ = str

	//// A send to a closed channel panics
	// var c chan string = make(chan string)
	// close(c)
	// c <- "let's get started"

	//// A receive from a closed channel returns the zero value immediately
	// var c chan string = make(chan string)
	// close(c)
	// byteData, ok := <-c
	// _ = ok
	// _ = byteData

	//// Закрытие закрытого канала приводит к панике
	// var c chan string = make(chan string)
	// close(c)
	// close(c)

	//// Закрытие нулевого канала приводит к панике
	var c chan string
	close(c)
}

package main

import (
	"errors"
	"fmt"
)

type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string {
	return e.Query
}

var ErrExample = errors.New("Error example!")

func main() {
	//// For if errors.Is(err, ErrExample) {
	// err := fmt.Errorf("user %w not found", ErrExample)

	//// For if errors.As(err, &pQueryError) {
	var pQueryError *QueryError = &QueryError{Query: "Error QueryError"}

	err := fmt.Errorf("Error 2: wrape %w", pQueryError)

	if errors.Is(err, pQueryError) {
		fmt.Println("err.Is(ErrExample)!")
	}

	if errors.As(err, &pQueryError) {
		fmt.Println("err.As(pQueryError)!")
		fmt.Printf("%v [%v]\n", err, errors.Unwrap(err))
	}

	fmt.Printf("%+8v, %T\n", err, err)
}

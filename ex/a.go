package main

import "fmt"

func main() {
	MyPrint("Hello binary!")
	MyFunc()

	fmt.Println("We started")
	Block{
		Try: func() {
			fmt.Println("I tried")
			Throw("Oh,...sh...")

			fmt.Println("WTF WHY I'm HERE!?")
		},
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finally...")
		},
	}.Do()
	fmt.Println("We went on")
}

func MyPrint(text string) {
	fmt.Println(text)
}

func MyFunc() {
	var value int = 0
	MyPrint("This is MyFunc!")
	MyRecurs(&value)
}

func MyRecurs(pVal *int) {
	if *pVal < 3 {
		*pVal++
		MyPrint(fmt.Sprintf("Val = %d", *pVal))
		MyRecurs(pVal)
	}
}

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

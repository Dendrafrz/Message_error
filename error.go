package main

import "fmt"

func badCall() {
	panic("Bad call happend!")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panicking %s\n\r", err)
		}
	}()

	badCall()
	fmt.Println("This is never executed!!")
}

func main() {
	fmt.Println("Start testing")
	test()
	fmt.Println("End testing")
}

package main

import "fmt"

func main() {
	fmt.Println("This is poniters")

	//////creation of pointer
	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	////intilisation of pointers

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("The new value of pointer is:", myNumber)
}

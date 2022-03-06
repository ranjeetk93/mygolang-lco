package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions of Golang")
	greeter()
	greeterTwo()
}

func greeter() {
	fmt.Println("Namaste from golang")
}

func greeterTwo() {
	fmt.Println("This is function from greeter")

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proRes, myMessage := proAdder(2, 3, 4, 5, 6)
	fmt.Println("Pro Result is:", proRes)
	fmt.Println("Pro Message is:", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	//if you want to return string as well
	return total, "Hi this is pro result"
}

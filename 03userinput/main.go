package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input exersise"
	fmt.Println(welcome)

	// comma  || error syntax
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza center")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of pizza rating is %T", input)
}

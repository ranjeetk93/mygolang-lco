package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Our Pizza Store")
	fmt.Println("Pleaser Rate Our Pizza Store")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating:", input)
	fmt.Printf("Type of Rating is %T ", input)
}

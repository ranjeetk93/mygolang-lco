package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch Case in Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice Number is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is one now you can open")
	case 2:
		fmt.Println("Dice Value is 2 you can move 2 spot")
	case 3:
		fmt.Println("Dice value is 3 you can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 you can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("Dice Value is 5 you can move 5 spot")
	case 6:
		fmt.Println("Dice value is 6 You can move 6 and roll again")
	default:
		fmt.Println("What was that!")
	}
}

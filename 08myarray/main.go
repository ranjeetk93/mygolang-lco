package main

import "fmt"

func main() {
	fmt.Println("Welcome to array class of golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "orange"

	fmt.Println("fruit list is :", fruitList)
	fmt.Println("fruit list is :", len(fruitList))

	var vegList = [3]string{"potato", "tomato", "beans"}
	fmt.Println("Vegitable list is:", vegList)
	fmt.Println("Length of array is :", len(vegList))
}

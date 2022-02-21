package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slice of Golang")

	var fruitList = []string{"Apple", "Potato", "Beans"}

	// fmt.Println("The type of fruitlist is:", fruitList)
	// fmt.Printf("The type of fruitlist is %T", fruitList)

	fruitList = append(fruitList, "Banana", "Tinda")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 111
	highScores[2] = 222
	highScores[3] = 333

	fmt.Println(highScores)

	highScores = append(highScores, 123, 456, 789)

	fmt.Println(highScores)
	fmt.Println("is array sort:", sort.IntsAreSorted(highScores))
	if (sort.IntsAreSorted(highScores)) == true {
		fmt.Println("This array is sorted")
	} else {
		fmt.Println("This array is not sorted")
	}
	sort.Ints(highScores)
	fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slice based in index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

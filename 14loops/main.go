package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Loop")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	//in place of index you can use (_)
	// 	fmt.Printf("Index is and value is %v\n", day)
	// }

	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}
		if rougeValue == 5 {
			rougeValue++
			continue

		}
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jump into learncode.in")
}

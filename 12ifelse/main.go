package main

import "fmt"

func main() {
	fmt.Println("This is code IF Else of Golang")

	loginCount := 10
	var result string

	//type one
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"

	} else {
		result = "Exctly 10 Login COunt"
	}
	fmt.Println(result)

	//type two
	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is odd")
	}

	//type three

	if num := 3; num < 10 {
		fmt.Println("Number is less than 3")
	} else {
		fmt.Println("Number is not less than 3")
	}

	//type four

	// if err !null {

	// }
}

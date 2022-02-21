package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Glang Maps")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "Ruby"

	fmt.Println("List of All courses: ", languages)
	fmt.Println("JS sort for: ", languages["js"])

	// delete(languages, "rb")
	fmt.Println("List of language: ", languages)

	//introduction of loops in golang

	for key, value := range languages {
		fmt.Printf("For Key %v, For value %v\n", key, value)
	}
}

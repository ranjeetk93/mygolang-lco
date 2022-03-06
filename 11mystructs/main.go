package main

import "fmt"

// func main() {
// 	fmt.Println("Welcome to my Structs Class in golang")

// 	//inheritance ,no super and no parent in golang

// 	ranjeet := User{"Ranjeet", "ranjeet@go.dev", true, 26}
// 	fmt.Println(ranjeet)
// 	//for user of struct on this format
// 	fmt.Printf("Ranjeet Details are %+v\n", ranjeet)
// 	fmt.Printf("Name of User is: %v \nEmail of User is:%v", ranjeet.Name, ranjeet.Email)

// }

// type User struct {
// 	Name   string
// 	Email  string
// 	Status bool
// 	Age    int
// }

func main() {
	fmt.Println("Welcome to Structs in golang")

	ranjeet := User{"Ranjeet", "ranjeet@go.dev", true, 26}
	fmt.Println(ranjeet)

	fmt.Printf("Ranjeet Details are: %v\n", ranjeet)
	fmt.Printf("Name of User is: %v\nEmail of User is:%v", ranjeet.Name, ranjeet.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

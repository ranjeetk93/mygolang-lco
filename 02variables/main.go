package main

import "fmt"

const LoginToken string = "This is public variable" //public var

func main() {
	var username string = "Ranjeet"
	fmt.Println(username)
	fmt.Printf("This is variable of type: %T \n", username)
	fmt.Println(LoginToken)
	fmt.Printf("variable is type: %T \n", LoginToken)
}

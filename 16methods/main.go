package main

import "fmt"

func main() {
	fmt.Println("Welcome to methos of golang and we are using structs")

	ranjeet := User{"Ranjeet", "ranjeet@go.dev", true, 26}
	fmt.Println(ranjeet)

	fmt.Printf("Name of User:%v\nEmail of User is:%v\n", ranjeet.Name, ranjeet.Email)
	ranjeet.GetStatus()
	ranjeet.NewEmail()
	fmt.Printf("Name of User:%v\nEmail of User is:%v\n", ranjeet.Name, ranjeet.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is User Active:", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("New Email is", u.Email)
}

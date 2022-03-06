package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://bizcarver.com"

func main() {
	fmt.Println("Welcome to Golang Web requests")
	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Responce type of is:%T\n", responce)
	defer responce.Body.Close() //caller's responcibility to colse the connection

	datatype, err := ioutil.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}

	content := string(datatype)

	fmt.Println(content)
}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files of Golang")
	content := "This is in the file of Golang"

	file, err := os.Create("./mylcogofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Lenght is :", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data insde the file is:\n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

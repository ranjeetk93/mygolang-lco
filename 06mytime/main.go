package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is Golang Time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.February, 10, 11, 12, 0, 0, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}

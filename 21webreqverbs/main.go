package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to video of webrequests")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:7000/get"
	responce, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()

	fmt.Println("Status code:", responce.StatusCode)
	fmt.Println("COntent Length is:", responce.ContentLength)

	var responceString strings.Builder
	content, _ := ioutil.ReadAll(responce.Body)
	byteCount, _ := responceString.Write(content)

	fmt.Println("Byte Count is:", byteCount)
	fmt.Println(responceString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:7000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"courseName":"Course of Golang",
			"price":0,
			"plateform":"youtube.com"

		}
	`)

	responce, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()

	content, _ := ioutil.ReadAll(responce.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:7000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "ranjeet")
	data.Add("lastname", "yadav")
	data.Add("email", "ranjeet@go.dev")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

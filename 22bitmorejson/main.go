package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to video of  JSON")
	// EncodedJson()
	DecodeJson()
}

func EncodedJson() {
	lcoCourses := []course{
		{"ReachJS Bootcamp", 299, "youtube.com", "abc123", []string{"web-dev", "JS"}},
		{"MERN Bootcamp", 199, "youtube.com", "abc123", []string{"Full-stack", "JS"}},
		{"Golang Bootcamp", 299, "youtube.com", "abc123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDatafromWeb := []byte(`
	{
		"coursename": "ReachJS Bootcamp",
		"Price": 299,
		"website": "youtube.com",        
		"tags": ["web-dev","JS"]
	}
	`)

	var lcoCourses course

	checkValid := json.Valid(jsonDatafromWeb)

	if checkValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonDatafromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some case where you want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDatafromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}
}

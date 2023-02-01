package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func PostFormReq() {
	const MyUrl string = "http://localhost:1111/postform"

	data := url.Values{}

	var first string
	var last string
	var age uint64

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&first)
	fmt.Println("Enter Last Name: ")
	fmt.Scanln(&last)
	fmt.Println("Enter Your Age: ")
	fmt.Scanln(&age)

	data.Add("first", first)
	data.Add("last", last)
	data.Add("age", strconv.FormatUint(age, 10))
	res, err := http.PostForm(MyUrl, data)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

}

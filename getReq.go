package main

import (
	"fmt"
	"net/http"
)

func GetReq() {
	const MyUrl string = "http://localhost:1111/get"

	res, err := http.Get(MyUrl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		fmt.Println("Status code is OKAY ")
		ParserAndValidator(res)
	} else {
		fmt.Println("SERVER IS DOWN!!!")
	}
}

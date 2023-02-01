package main

import (
	"fmt"
	"net/http"
)

func GetJson() {
	const MyUrl string = "http://localhost:1111/getData"

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

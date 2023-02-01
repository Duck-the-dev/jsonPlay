package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PostReq() {
	const MyUrl string = "http://localhost:1111/post"

	reqBody := strings.NewReader(`
		{
			"name":"mohamed",
			"age": 0,
			"uni":"i don't know"
		}
		`)
	res, err := http.Post(MyUrl, "application/json", reqBody)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}

	con, _ := io.ReadAll(res.Body)

	fmt.Println(string(con))
}

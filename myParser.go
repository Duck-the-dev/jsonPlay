package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type person struct {
	name string
	age  uint64
	uni  string
}

func ParserAndValidator(res *http.Response) {
	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)
	resString.Write(content)
	var strRes = resString.String()
	strRes = strings.ReplaceAll(strRes, "[", "")
	strRes = strings.ReplaceAll(strRes, "]", "")
	strRes = strings.ReplaceAll(strRes, "{", "")
	strRes = strings.ReplaceAll(strRes, "}", "")
	strRes = strings.ReplaceAll(strRes, "\"", "")
	resp := strings.Split(strRes, ",")
	var newResp []string
	per := person{}

	for i, _ := range resp {
		newResp = strings.Split(resp[i], ":")
		if newResp[0] == "name" {
			per.name = newResp[1]
			fmt.Println("person name is : ", per.name)
		}
		if newResp[0] == "age" {
			if _, err := strconv.Atoi(newResp[1]); err == nil {
				per.age, _ = strconv.ParseUint(newResp[1], 10, 32)

			} else {
				panic(err)
			}
			per.age, _ = strconv.ParseUint(newResp[1], 10, 32)
			fmt.Println("person age is : ", per.age)

		}
		if newResp[0] == "uni" {
			per.uni = newResp[1]
			fmt.Println("person uni is : ", per.uni)

		}
	}

}

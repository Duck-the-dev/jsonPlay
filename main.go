package main

import "fmt"

func main() {
	var chose string

	for {
		fmt.Println("Choose one of the end points : ")
		fmt.Println("1. /get")
		fmt.Println("2. /getData")
		fmt.Println("3. /post")
		fmt.Println("4. /postform")
		fmt.Println("----------")
		fmt.Println("Enter 'Q' to Quit the App")

		fmt.Scanln(&chose)
		switch chose {
		case "1":
			GetReq()
			continue
		case "2":
			GetJson()
			continue
		case "3":
			PostReq()
			continue
		case "4":
			PostFormReq()
			continue
		case "q", "Q":
			panic("Thank you and see you soon")

		default:
			fmt.Println("Invalid choice, Try again")
			continue

		}
	}
}

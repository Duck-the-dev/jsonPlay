package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprint(randomFormat())
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
func main() {
	var chose string

	for {
		fmt.Println("Choose one of the end point's : ")
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
			panic("Thank you and see you soon, Bye")

		default:
			fmt.Println("Invalid choice, Try again")
			continue

		}
	}
}

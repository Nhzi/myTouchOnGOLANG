package main

import "fmt"

func main() {
	var user string

	user = "Chinaza"

	switch user {
	case "Ved":
		fmt.Println("You are root")
	case "":
		fmt.Println("Who are you?")
		fmt.Println("Please set up your Username for file recognition and access.")
	case "Sheriff":
		fmt.Println("You are not root")
	default:
		fmt.Println("404 Not Found")
		fmt.Println(user, "Not Found in DB")
	}
}

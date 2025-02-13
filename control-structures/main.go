package main

import "fmt"

func main() {
	var num int
	num = 9

	if num == 0 {
		fmt.Println("WTF do you want me to do with", num)
	} else {
		fmt.Println("ouuuuuuu", num, "I can work with that")
	}

	name := "Sheriff"
	if name == "Sheriff" {
		num = 3
	}
	if num == 3 {
		fmt.Println("Welcome AbdulSheriff")
		fmt.Println("The Bomb would be ready soon")
	} else {
		fmt.Println("You aint Sheriff you bitch!")
	}
}

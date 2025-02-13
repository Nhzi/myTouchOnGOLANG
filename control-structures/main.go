package main

import "fmt"

func main() {
	var num int
	num = 0

	if num == 0 {
		fmt.Println("WTF do you want me to do with", num)
	} else if num == 9 {
		fmt.Println("ouuuuuuu", num, "I can work with that, I would jus need '11' at the back")
	} else {
		fmt.Println("Lets see.")
	}

	name := "Sheriff"

	var password string
	password = ""

	authenticator(name, num)
	passwordChange(&password, "nhzi120438")
	fmt.Println(password, "is the new Password Set by User:", name)
}

func authenticator(name string, num int) {
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

func passwordChange(old *string, new string) {
	fmt.Println("It is a Start of the password Change")
	fmt.Println("Get Geeked; Nhzi")
	fmt.Println("0xAddress is:", old)
	*old = new
}

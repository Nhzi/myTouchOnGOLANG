package main

import "fmt"

func main() {
	name := "Naz"

	var Author string
	Author = "Ved ZhaDev"

	pointerAddress(&name, Author)
	// fmt.Println("Switching Places, we get! ", &n)
}

func pointerAddress(n *string, a string) {
	//See only the n is a string!
	fmt.Println("The Address on the memory is ", n, "And the value is; Name:", *n)
	// fmt.Println("The Address on the memory is ", a, "And the value is; Author:", a)
	// now we switch
	*n = a

	fmt.Println("User switched to", *n, "now.")
}

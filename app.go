package main

import "fmt"

func main() {
	name := "Naz"
	author := "Ved ZhaDev"

	pointerAddress(&name, &author)
	// fmt.Println("Switching Places, we get! ", &n)
}

func pointerAddress(n *string, a *string) {
	fmt.Println("The Address on the memory is ", n)
	fmt.Println("The Address on the memory is ", a)
}

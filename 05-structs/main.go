package main

import "fmt"

type Person struct {
	Firstname, Lastname, Country string
	Age                          int
}

func main() {
	Prototype := Person{
		Firstname: "Ved",
		Lastname:  "Zhadev",
		Country:   "Russia"}

	fmt.Println("Structs; My First Creation is", Prototype)
	fmt.Println("My Name surely is King", getProperties(Prototype))
}

//We wanna create a function that returns only properties of a struct.
func getProperties(name Person) string {
	return name.Firstname + " " + name.Lastname
}

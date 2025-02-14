package main

import (
	"fmt"
	"math"
)

func main() {
	//Beginning of FOR loop

	i := 0
	for i <= 12 {
		fmt.Println("The square of i at", i, "is", i*i)
		i++
	}
	//Second for loop
	for x := -1; x >= -12; x-- {
		fmt.Println("The exponetial of value x at", x, "is", math.Pow(float64(x), float64(x)))
	}
}

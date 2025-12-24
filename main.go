package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number = rand.Int31n(10)
	fmt.Println("Random number:", number)

	var a int = 5
	b := "String" // short hand syntax for declaring variables

	// The `:=` operator is distinct from the regular assignment operator `=`, which requires the variable to be already declared.
	fmt.Println("a:", a, "b:", b)

	isBoolean := false
	fmt.Print(isBoolean)
	// A

}

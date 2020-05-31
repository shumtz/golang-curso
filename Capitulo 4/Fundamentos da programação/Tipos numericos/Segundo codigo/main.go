package main

import (
	"fmt"
)

// Default é usado em var
var x = 10
var y = 10.0

func main() {
	// Default tambem é usado em :=
	// x := 10
	// y := 10.0

	// Um tipo só pode ser mudado se houver conversão
	// x = 20.5

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T", y, y)
}

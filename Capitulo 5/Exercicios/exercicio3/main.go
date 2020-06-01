package main

import (
	"fmt"
)

const x = 10
const y int = 10

func main() {
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T", y, y)
}

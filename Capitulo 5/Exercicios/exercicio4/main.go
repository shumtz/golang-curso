package main

import (
	"fmt"
)

func main() {
	var x = 10
	fmt.Printf("%d %#x %b\n", x, x, x)
	var y = x << 1
	fmt.Printf("%d %#x %b", y, y, y)
}

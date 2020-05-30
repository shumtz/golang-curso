package main

import "fmt"

type loxt int

func main() {
	var x loxt
	fmt.Println(x)
	fmt.Printf("%T \n", x)
	x = 42
	fmt.Println(x)
}

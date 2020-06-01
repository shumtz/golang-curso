package main

import "fmt"

const (
	a = iota + 1000000
	_
	c
	x
	_
	z
)

func main() {
	fmt.Println(a, c, x, z)
}

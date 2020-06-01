package main

import "fmt"

const (
	_ = 2000 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)
}

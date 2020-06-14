package main

import "fmt"

func main() {
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{10, 11, 12, 13}

	fmt.Println(umaslice)

	umaslice = append(umaslice, 5, 6, 7, 8, 9)
	fmt.Println(umaslice)

	umaslice = append(umaslice, outraslice...)
	fmt.Println(umaslice)
}

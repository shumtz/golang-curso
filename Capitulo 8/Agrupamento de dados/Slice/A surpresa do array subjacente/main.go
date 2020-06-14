package main

import (
	"fmt"
)

func main() {
	primeiroslice := []int{1, 2, 3, 4, 5}
	fmt.Println(primeiroslice)

	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)
	fmt.Println(segundoslice)

	fmt.Println(primeiroslice)
}

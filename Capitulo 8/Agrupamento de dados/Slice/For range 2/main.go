package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 10, 10, 10}

	total := 0

	for _, valor := range slice {
		total += valor
	}

	fmt.Println("O valor total é:", total)
}

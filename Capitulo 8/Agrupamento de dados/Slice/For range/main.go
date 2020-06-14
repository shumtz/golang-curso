package main

import (
	"fmt"
)

func main() {
	slice := []string{"Ferrari", "Red Bull", "McLaren", "Williams"}

	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor:", valor)
	}
}

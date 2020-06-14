package main

import "fmt"

func main() {
	sabores := []string{"Pepperoni", "Calabresa", "Marguerita", "Quatro queijos"}
	fatia := sabores[1:4]
	fmt.Println(fatia)

	fmt.Println(sabores[0], sabores[1], sabores[2], sabores[3])

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

}

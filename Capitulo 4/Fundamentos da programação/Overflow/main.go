package main

import (
	"fmt"
)

func main() {
	var i uint16
	i = 65535
	// Volume ultrapassado, maximo é 65535
	// i = 65536
	fmt.Println(i)
	// O que acontece se adicionar mais um?
	// i++
	// fmt.Println(i)
	// Ele zera

	// i++
	// fmt.Println(i)
	// Como zerou ele vai pra um e assim adiante
}

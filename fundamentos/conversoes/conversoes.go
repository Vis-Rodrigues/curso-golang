package main

import (
	"fmt"
	str "strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(97)) // Mostra o código correspondente da tabela ASC

	// int para string
	fmt.Println("Teste " + str.Itoa(123))

	// string para int
	num, _ := str.Atoi("123")
	fmt.Println(num - 122)

	b, _ := str.ParseBool("true") // poderia ser 1 = true ou 0 = false
	if b {
		fmt.Println("Verdadeiro")
	}
}

package main

import "fmt"

func main() {

	i := 1

	// Go não tem operações aritméticas de ponteiros

	// ponteiro é uma referência de memória
	var p *int = nil

	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}

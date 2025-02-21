package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64 u=ussined (sem sinal)
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", i2)
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	var z = float64(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo de z é", reflect.TypeOf(z))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // sempre será do tipo float64

	// boolean
	bo := true
	fmt.Println("O tipo de bool é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá, estamos aprendendo Golang! Graças a Deus"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//string com multiplas linhas
	s2 := `Olá,
	estamos
	apredendo
	Golang!
	Graças a Deus!`
	fmt.Println("O tamanho da string é", len(s2))

	// não existe char no Go
	char := 'b'
	fmt.Println("O tipo de char é", reflect.TypeOf(char)) //int32
	fmt.Println(char)
}

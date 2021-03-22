// NUMEROS INTERIOS E NUMEROS REAIS
package main

import (
	"errors"
	"fmt"
)

func main() {
	// NUMEROS INTEIROS
	var num int = 1000000000001
	num2 := 1000000000000000
	num3 := -235000
	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(num3)

	var num4 uint32 = 123456
	fmt.Println(num4)

	// alias
	//INT32 = RUNE
	var num5 rune = 1234567
	fmt.Println(num5)

	// Byte = UINT8
	var num6 = 123
	fmt.Println(num6)

	// FIM DOS NUMEROS INTEIROS

	// NUMEROS REAIS //

	var numreal1 float32 = 897.83
	fmt.Println(numreal1)

	var numreal2 float64 = 123.44
	fmt.Println(numreal2)

	numreal3 := 456.43
	fmt.Println(numreal3)

	// FIM DOS NUMEROS REAIS

	// STRINGS
	var str string = "Texto1"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// FIM STRING

	// VALOR ZERO

	var texto int
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var booleano2 bool = true
	fmt.Println(booleano2)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro Interno")
	fmt.Println(erro2)
}

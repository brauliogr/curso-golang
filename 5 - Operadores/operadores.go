package main

import "fmt"

func main() {
	// ARITIMETICOS

	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 4
	multiplicacao := 1 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	somatorio := numero1 + numero2
	fmt.Println(somatorio)

	// FIM DOS ARITIMETICOS

	// ATRIBUIÇÃO

	var variavel1 string = "string"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)

	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	//OPERADORES RELACIONAIS

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	//FIM DOS RELACIONAIS

	//OPERADORES LOGICOS
	fmt.Println("------------------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// FIM DOS OPERADORES LOGICOS

	// OPERADORES UNARIOS

	numero := 10
	numero++
	fmt.Println(numero)

	numero--
	numero -= 20 //numero = numero -20
	numero *= 3  // numero = numero * 3
	numero /= 10
	numero %= 3

	// FIM DOS OPERADORES UNARIOS

	// OPERADORES TERNARIO

	fmt.Println(numero)

	fmt.Println("----------------------------------")

	var texto string
	if numero > 5 {
		texto = "Maior que 5 "
	} else {
		texto = "Menor que 5 "
	}

	fmt.Println(texto)

}

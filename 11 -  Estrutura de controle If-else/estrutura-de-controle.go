package main

import "fmt"

func main() {

	numero := 50

	if numero > 15 {
		fmt.Println("Maior que 15")

	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if numero2 := numero; numero2 > 0 {
		fmt.Println("Numero maior que 0")
	} else if numero < -10 {
		fmt.Println("Numero é menor que -10")

	} else {
		fmt.Println("Entre 0 e 10")
	}

}

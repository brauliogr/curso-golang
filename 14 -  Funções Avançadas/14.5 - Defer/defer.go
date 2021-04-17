package main

import "fmt"

func alunoaprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado será retornado.")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 7 {

		return true
	}

	return false

}

func main() {
	fmt.Println(alunoaprovado(7, 8))

}

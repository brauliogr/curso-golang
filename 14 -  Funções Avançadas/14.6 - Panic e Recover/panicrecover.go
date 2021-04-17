package main

import "fmt"

func recuperarexecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso.")
	}

}

func calculomedia(n1, n2 float64) bool {
	defer recuperarexecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MEDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(calculomedia(6, 6))
	fmt.Println("Pós execução")

}

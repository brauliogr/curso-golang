package main

import (
	"fmt"
)

func main() {
	//	i := 0
	//	for i < 10 {
	//		i++
	//		fmt.Println("Incrementando i")
	//time.Sleep(time.Second)
	//	}

	//fmt.Println(i)

	//for j := 1; j < 10; j++ {
	//		fmt.Println("Incrementando J")
	//	time.Sleep(time.Second)

	//}

	nomes := []string{"Ceciliane", "Bernardo", "Emanuelle"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
	usuario := map[string]string{
		"nome":      "Braulio",
		"sobrenome": "Gomes",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)

	}

}

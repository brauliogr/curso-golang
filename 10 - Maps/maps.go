package main

import "fmt"

func main() {
	usuario := map[string]string{
		"Nome":      "Pedro",
		"Sobrenome": "Silva",
	}

	fmt.Println(usuario["Nome"], usuario["Sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"Primeiro": "João",
			"Ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus1",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}
	fmt.Println(usuario2)
}

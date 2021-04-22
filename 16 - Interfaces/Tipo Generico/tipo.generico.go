package main

import (
	"fmt"
)

func generica(interf interface{}) {
	fmt.Println(interf)
}

// Forma Generica nem sempre é bom ser usada, por causa da segurança

//Exemplo de uma função utilizando formas genericas ignorando os tipos.

func main() {
	generica("String")
	generica(1)
	generica(true)
	generica(false)
	fmt.Println(1, 2, "Testando String", false, true, float64(12345))
	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"string":     "String",
	}
	fmt.Println(mapa)
}

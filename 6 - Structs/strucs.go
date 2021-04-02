package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint16
	bairro string
}

func main() {
	fmt.Println("Arquivo Structs")

	var user usuario
	user.nome = "Bernardo"
	user.idade = 7
	fmt.Println(user)

	enderecoexemplo := endereco{"Rua Albano Schmidt", 1499, "Boa Vista"}

	user2 := usuario{"Davi", 21, enderecoexemplo}
	fmt.Println(user2)

	user3 := usuario{idade: 21}
	fmt.Println(user3)

}

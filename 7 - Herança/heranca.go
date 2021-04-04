package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	user := pessoa{"Braulio", "Gomes", 42, 176}
	fmt.Println(user)

	user2 := estudante{user, "Engenhaia", "EStacio"}
	fmt.Println(user2.idade)
}

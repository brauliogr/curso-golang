package main

import "fmt"

type usuario struct {
	nome  string
	idade uint16
}

func (user usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco\n", user.nome)
}

func (user usuario) maiordeIdade() bool {
	return user.idade >= 18

}

func (user *usuario) fazeraniversario() {
	user.idade++
}

func main() {
	usuario1 := usuario{"Bernardo", 07}
	usuario1.salvar()

	usuario2 := usuario{"Emanuelle", 11}
	usuario2.salvar()

	maiordeIdade := usuario2.maiordeIdade()
	fmt.Println(maiordeIdade)

	usuario1.fazeraniversario()
	fmt.Println(usuario1.idade)

}

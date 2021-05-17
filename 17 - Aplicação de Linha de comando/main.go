package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Ponto de Partida")

	aplicacao := app.Gerar()
	aplicacao.Run(os.Args)

}

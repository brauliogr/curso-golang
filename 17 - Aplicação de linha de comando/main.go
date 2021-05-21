package main

import (
	"linhadecodigo/app"
	"log"
	"os"
)

func main() {

	var aplicacao = app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}

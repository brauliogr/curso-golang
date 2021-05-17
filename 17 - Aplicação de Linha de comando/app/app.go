package app

import "github.com/urfave/cli"

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de comando"
	app.Usage = "Buca Ips e Nomes na Internet"

	return app
}

package app

//Gerar Vai retornar a aplicação pronta para ser executada

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de comando"
	app.Usage = "Busca Ips e nomes de Servidores"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca Ips de endereços na internet",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: buscaIps,
		},
	}

	return app

}

func buscaIps(c *cli.Context) error {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}

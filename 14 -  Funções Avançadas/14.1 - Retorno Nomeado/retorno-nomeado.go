package main

import "fmt"

func calculos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n2 - n1
	return

}

func main() {

	soma, subtracao := calculos(10, 20)
	fmt.Println(soma, subtracao)

}

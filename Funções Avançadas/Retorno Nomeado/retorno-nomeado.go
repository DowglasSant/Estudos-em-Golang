package main

import "fmt"

func calculosMatematicos(n1, n2 float32) (soma float32, subtracao float32) {
	soma = n1 + n2
	subtracao = n1 - n2

	return
}

func main() {
	fmt.Println("Funções com Retorno Nomeado")

	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}

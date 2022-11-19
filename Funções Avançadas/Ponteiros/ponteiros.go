package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	numeroNovo := 40
	fmt.Println(numeroNovo)

	inverterSinalPonteiro(&numeroNovo)
	fmt.Println(numeroNovo)
}

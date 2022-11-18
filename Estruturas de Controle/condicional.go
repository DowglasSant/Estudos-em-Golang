package main

import "fmt"

func main() {
	fmt.Println("Condicionais - Fluxo de Execução!")

	var numero float32

	fmt.Println("Digite o numero:")
	fmt.Scanln(&numero)

	if numero > 15 {
		fmt.Println("Maior que 15!")
	} else {
		fmt.Println("Menor ou igual a 15!")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que 0!")
	} else {
		fmt.Println("Numero é menor ou igual a 0!")
	}
}

package main

import "fmt"

//Permite não limitar a quantidade de paramentros que vem na chamada da função. Se comporta como um Slice.
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

//Importante: O parametro variático precisa ser o último parametro e vir depois dos parametros fixos.
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	total := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(total)

	escrever("Ola mundo!", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

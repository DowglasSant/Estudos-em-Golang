package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops!")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j <= 10; j++ {
	// 	fmt.Println(j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "Carla"}

	//O primeiro é sempre o indice.
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	fmt.Println("------------------")

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("------------------")

	//Percorre palavra e considera "letra" na tabela ASCII
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	fmt.Println("------------------")

	//Percorre palavra e converte letra para string
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	fmt.Println("------------------")

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	fmt.Println("------------------")

	fmt.Println("Fim!")
}

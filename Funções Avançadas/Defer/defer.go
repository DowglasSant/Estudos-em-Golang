package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média Calculada. Resultado:")
	fmt.Println("Acessando função!")

	media := (n1 + n2) / 2

	if media <= 6 {
		return true
	}

	return false
}

func main() {
	funcao1()
	funcao2()

	defer funcao1()
	funcao2()

	fmt.Println(alunoAprovado(6, 6))
}

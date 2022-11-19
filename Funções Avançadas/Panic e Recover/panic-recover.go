package main

import "fmt"

func recuperarExecucao() {
	//tenta recuperar o fluxo de execução da aplicação
	if r := recover(); r != nil {
		fmt.Println("Recuperado com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	//faz programa entrar em panic
	//encerra a execução da aplicação
	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pos Execução")
}

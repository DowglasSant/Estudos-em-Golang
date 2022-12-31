package main

import "fmt"

func main() {
	//Canal com Buffer, segundo paramentro define a capacidade do canal
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Programming in Golang!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}

package main

import "fmt"

func main() {
	fmt.Println("Ponteiros!")

	var variavel1 int = 12

	//isso cria uma copia do valor de v1 para v2
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//Ponteiro é uma referencia de memoria

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)

	//mudam
	fmt.Println(variavel3, *ponteiro)

}

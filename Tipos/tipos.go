package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 1000000000000000000

	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	//UINT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	//float
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	numeroReal4 := 123.45
	fmt.Println(numeroReal4)

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	//Valor 0

	var texto string
	fmt.Println(texto)

	var numero5 int64
	fmt.Println(numero5)

	//boolean
	var boolean bool = true
	var boolean2 bool = false

	fmt.Println(boolean)
	fmt.Println(boolean2)

	//error
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Internal Error")
	fmt.Println(erro2)

}

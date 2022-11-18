package main

import "fmt"

func somar(num1 int64, num2 int64) int64 {
	return num1 + num2
}

func calculos(num1, num2 float32) (float32, string) {
	resultado := num1 * num2
	return resultado, "Multiplicado."
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função func")
	fmt.Println(resultado)

	multiplicar, mensagem := calculos(30.5, 2.1)
	fmt.Println(multiplicar, mensagem)

	//Quando só um retorno interessa
	multiplicar2, _ := calculos(30.5, 2.1)
	fmt.Println(multiplicar2)
}

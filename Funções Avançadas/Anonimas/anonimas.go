package main

import "fmt"

func main() {

	concat := func(texto string) string {
		return fmt.Sprintf("Recebido: %s", texto)
	}("Passando Parâmetro") //A chamada da função se dá pelos parenteses depois das chaves

	fmt.Println(concat)
}

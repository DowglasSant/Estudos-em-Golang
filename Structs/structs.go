package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 12
	fmt.Println(u)

	enderecoDavi := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Davi", 12, enderecoDavi}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}

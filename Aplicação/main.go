package main

import (
	"aplicacao/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting...")

	aplicacao := app.Gerar()
	if error := aplicacao.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}

package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando Init...")
	n = 10
}

func main() {
	fmt.Println("Executando Main...")
	fmt.Println(n)
}

package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Soul")
	generica(12)
	generica(true)
	generica(3.14)
}

package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2"

	fmt.Println(variavel1 + " " + variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
		variavel5 int    = 5
	)

	variavel6, variavel7 := "varivael6", "variavel7"

	fmt.Println(variavel3 + " " + variavel4)
	fmt.Println(variavel5)

	fmt.Println(variavel6 + " " + variavel7)

	variavel6, variavel7 = variavel7, variavel6
	fmt.Println(variavel6 + " " + variavel7)
}

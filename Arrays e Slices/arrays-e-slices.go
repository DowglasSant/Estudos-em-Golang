package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	//Arrays
	var array1 [5]string

	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6, 10}
	fmt.Println(slice)

	slice = append(slice, 24)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Editada"
	fmt.Println(slice2)
	fmt.Println(slice2[1])

	//Arrays Internos
	slice3 := make([]float32, 10, 11)

	slice3 = append(slice3, 5.2)
	slice3 = append(slice3, 6.9)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 11.1)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}

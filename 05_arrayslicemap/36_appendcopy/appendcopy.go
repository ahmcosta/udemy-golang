package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6) // primeiro argumento deve ser um slice e não um array
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // [4 5]. O copy não aumenta de tamanho, por isso o [5] não é copiado. Ambos os parâmetros devem ser slices
	fmt.Println(slice2)
}

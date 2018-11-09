package main

import "fmt"

func main() {
	s := make([]int, 10) // como não havia um array para referenciar, GO cria um array interno com 10 posições
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // 10 elementos, mas com capacidade de 20 elementos
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) // dobra a capacidade do slice
}

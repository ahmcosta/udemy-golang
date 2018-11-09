package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta! se tirar o "..." seria um slice

	for i, numero := range numeros { // o range retorna o índice (i) e o elemento da iteração (numero)
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // se escrever 	num := range numeros	o "num" conterá os índices
		fmt.Println(num)
	}
}

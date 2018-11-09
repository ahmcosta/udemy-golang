package main

import "fmt"

/*
	Para uma variável receber uma função, esta dever ser ANÔNIMA
*/
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	/*
		Para uma variável receber uma função, esta dever ser ANÔNIMA
	*/
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}

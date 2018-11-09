package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(38004)) // valor do UTF8

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, err := strconv.Atoi("123")
	if err == nil {
		fmt.Println(num - 122)
	}

	// string para int
	num2, _ := strconv.Atoi("1234") // se usar o _, o conteúdo não será tratato.
	fmt.Println(num2 - 123)

	b, _ := strconv.ParseBool("true") // para a função ParseBool, somente a string true ou int 1 são true
	if b {
		fmt.Println("Verdadeiro")
	}
}

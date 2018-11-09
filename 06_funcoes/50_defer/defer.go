package main

import "fmt"

func obterValorAprovado(numero int) int {
	/*
		defer atrasa a execução do comando passado como seu parâmetro até que o momento imediatamente anterior a chamada do return
		Muito usado para fechar um recurso que foi aberto
	*/
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}

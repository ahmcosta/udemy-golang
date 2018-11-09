package main

import "fmt"

func imprimirAprovados(aprovados ...string) { // quantidade de parâmetros variáveis. trabalhados como se fosse um ARRAY
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	/*
		se fosse passado [...], o aprovados seria um ARRAY, pois o compilador contaria a quantidade de elementos. Neste caso é um slice
		Só funciona com slice. Não funciona com ARRAY
	*/
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...) // explode o slide como se fosse parâmetros para a função
}

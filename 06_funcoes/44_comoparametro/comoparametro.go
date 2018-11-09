package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func adicao(a, b int) int {
	return a + b
}

func subtracao(a, b float64) int {
	return int(a + b)
}

/*
	@params
	1º => Espera receber uma função que recebe dois inteiros e retorna um inteiro	==> funcao func(int, int) int
	2º => Inteiro																	==> p1
	3º => Inteiro																	==> p2
*/
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2) // retorna o retorno da função passada como parâmetro
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println("Multiplicação:", resultado)

	resultado = exec(adicao, 3, 4)
	fmt.Println("Adição:", resultado)

	/*
		resultado = exec(subtracao, 3, 4) // erro de compilação. Assinatura da função esperada diverge da passada
		fmt.Println("Subtração:", resultado)
	*/
}

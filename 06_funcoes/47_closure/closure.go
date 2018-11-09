package main

import "fmt"

/*
	https://go-tour-br.appspot.com/moretypes/25

	As funções Go podem ser closures. Um closure é uma função valor que referencia variáveis de
	fora de seu corpo. A função pode acessar e atribuir nas variáveis referenciadas; nesse
	sentido a função é "limitada" às variáveis.

	Por exemplo, a função adder retorna um closure. Cada closure limita sua própria variável sum.

	package main

	import "fmt"

	func adder() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	func main() {
		pos, neg := adder(), adder()
		for i := 0; i < 10; i++ {
			fmt.Println(
				pos(i),
				neg(-2*i),
			)
		}
	}

*/
func closure() func() {
	x := 10 // escopo da função. A função retornada sabe onde foi definida, mesmo que neste exemplo o valor de x tenha sido declarado fora da função anônima
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX() // chama a função devolvida em imprimeX
}

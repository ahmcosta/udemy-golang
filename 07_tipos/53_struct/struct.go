package main

import "fmt"

type produto struct {
	/*
		Não se separa com vírgulas (,).
		nome (atributo)		tipo

	*/
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
// (p produto) é o dono da função
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		// deve-se usar dois pontos (:) as vírgulas (,)
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05, // atenção a esta última vírgula
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1989.90, 0.10} // nesta forma, não é preciso nomear os atributos, mas deve-se colocá-los na ordem
	fmt.Println(produto2.precoComDesconto())
}

package main

import "fmt"

/*
	se a interface tem mais de um método, as structs devem
	implementar TODOS os métodos para serem compatíveis
*/
type imprimivel interface {
	toString() string
	tosco(texto string) // criado por mim
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string { // mesma assinatura da interface
	return p.nome + " " + p.sobrenome
}

// criado por mim
func (p pessoa) tosco(texto string) {
	fmt.Print("pessoa, ", texto)
}

func (p produto) toString() string { // mesma assinatura da interface
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

// criado por mim
func (p produto) tosco(texto2 string) {
	fmt.Print("produto, ", texto2)
}

/*
	não há marcação que pessoa e produto implementam imprimivel, mas o
	compilador infere devido às funções receivers terem a mesma
	assinatura da interface imprimivel
*/
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"} // coisa pode ser pessoa ou produto
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Calça Jeans", 179.90}
	imprimir(p2)
}

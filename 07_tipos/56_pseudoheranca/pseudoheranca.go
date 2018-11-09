package main

import "fmt"

/*
	Não existe herança em GO
*/

type carro struct {
	nome            string
	velocidadeAtual int
}

type caracteristicas struct {
	cor         string
	tamanhoRoda byte
}

type ferrari struct {
	carro // campos anonimos
	caracteristicas
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40" // ferrari não tem o atributo nome, mas como o atributo carro dela é anônimo, dá a impressão que houve herança
	f.velocidadeAtual = 0
	f.turboLigado = true
	f.cor = "vermelha"
	f.tamanhoRoda = 23

	fmt.Printf("A ferrari %s, de cor %s e roda %d'', está com turbo ligado? %v\n", f.nome, f.cor, f.tamanhoRoda, f.turboLigado)
	fmt.Println(f)
}

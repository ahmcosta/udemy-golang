package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() { // efeito colateral!!! Usa ponteiro
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0} // modo tradicional estudado na aula 55_metodos
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0} // quando se usa a interface (esportivo), deve-se passar o endereço da estrutura que implementa a interface (ferrari)
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)

	/*
		Normalmente quando se usa interface, não se está preocupado em alterar valores, mas lê-los de forma padronizada, sem efeito colateral
	*/
}

package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	/*
		É possível usar a func catetos, mesmo sendo privada, pq ambos os arquivos estão dentro do mesmo pacote
	*/
	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}

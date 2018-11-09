package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int, modo string) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("[%s]\t%s: %s (iteração %d)\n", modo, pessoa, texto, i+1)
	}
}

func main() {
	// Forma serial
	fale("Maria", "Pq vc não fala comigo?", 3, "serial")
	fale("João", "Só posso falar depois de vc!", 1, "serial")

	/*
		No trecho a seguir, nada é escrito na sysout, pq o fluxo princial (main) encerrou antes da goroutine acabar
	*/
	// Forma concorrente (Go routine)
	go fale("Maria", "Ei...", 500, "go routine")
	go fale("João", "Opa...", 500, "go routine!")

	// Forma concorrente (Go routine)
	go fale("Maria", "Entendi!!!", 10, "go routine")
	fale("João", "Parabéns!", 5, "serial")
}

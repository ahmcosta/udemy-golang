package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	fmt.Println("[1] Executou!")
	ch <- 2
	fmt.Println("[2] Executou!")
	ch <- 3
	fmt.Println("[3] Executou!")
	ch <- 4
	fmt.Println("[4] Executou!") // Dependendo do time, pode ser que o consumo do canal seja executado antes da inclusão na linha anterior
	ch <- 5
	fmt.Println("[5] Executou!") // Não vai ser escrito. Buffer cheio!!!
	ch <- 6
	fmt.Println("[6] Executou!")
}

func main() {
	ch := make(chan int, 3) // buffer com 3 slots de inteiro
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}

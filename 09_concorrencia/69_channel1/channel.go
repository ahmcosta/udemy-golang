package main

import "fmt"

/*
	Concorrência em GO é inerente da linguagem. Não foi disponibilizada como uma biblioteca!
	Channel é um tipo na linguagem, assim como uint, float64, char, etc
*/
func main() {
	ch := make(chan int, 1) // cria um channel de inteiros com buffer de 1

	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}

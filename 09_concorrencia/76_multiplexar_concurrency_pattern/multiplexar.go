/*
	Google I/O 2012 - Go Concurrency Patterns
	Multiplexing Pattern: encapsules many channels.
	More info in: https://talks.golang.org/2012/concurrency.slide#27
*/

package main

import (
	"fmt"

	"github.com/ahmcosta/html"
)

// origem is a read-only channel. destino is a rw channel
func encaminhar(origem <-chan string, destino chan string) {
	for { // while true
		destino <- <-origem // forward
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Title("https://www.cod3r.com.br", "https://www.google.com"),
		html.Title("https://www.amazon.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}

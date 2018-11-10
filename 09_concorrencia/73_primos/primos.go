package main

import (
	"fmt"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				// time.Sleep(time.Millisecond * 180)
				break
			}
		}
	}
	close(c) // Quando se termina a função é interessante que se feche o canal
}

func main() {
	c := make(chan int, 30)
	go primos(100000, c)
	// go primos(cap(c), c)
	for primo := range c { // laço for sobre um canal
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}

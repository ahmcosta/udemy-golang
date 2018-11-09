package main

import "fmt"

/*
	É chamada antes do método main
*/
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}

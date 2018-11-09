package main

/*
	Se uma variável é declarada, mas não é utilizada, haverá um erro de compilação.
*/

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // se há atribuição o compilador infere o tipo da variável, no caso float64

	// forma reduzida de criar variável. Declara e inicializa a variável com um só comando.
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}

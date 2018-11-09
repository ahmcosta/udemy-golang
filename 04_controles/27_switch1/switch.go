package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		/*
			nas outras linguagens, o próximo case é executado a
			não ser que se encontre a palavra reservada "brake".
			Essa necessidade é interessante em um número muito
			pequeno de casos.
			Em GO a lógica é inversa. O case seguinte NÃO é
			executado a não ser que se use a palavra reservada
			"falltrough"
		*/
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println("9.8 =>", notaParaConceito(9.8))
	fmt.Println("6.9 =>", notaParaConceito(6.9))
	fmt.Println("2.1 =>", notaParaConceito(2.1))
}

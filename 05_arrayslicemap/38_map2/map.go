package main

import "fmt"

func main() {
	/*
		Pode ser inicializado com make(), como no exemplo map1
		Inicialização de forma literal
	*/
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0, // o último elemento deve ter a vírgula (,)
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente") // não dá mensagem de erro!!!

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}

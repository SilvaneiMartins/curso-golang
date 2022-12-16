package main

import "fmt"

var a int = 20
var b string = "Nome"

func main() {
	// Tipo(variavel) - conversão de tipo
	var c float64 = float64(a)

	fmt.Printf("O valor de C é: %g é o valor de B é: %s ", c, b)
}

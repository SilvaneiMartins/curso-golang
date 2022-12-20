package main

import "fmt"

func main() {
	i := 100

	if i%2 == 0 && i%3 <= 0 {
		fmt.Println("X é multiplo de 2 ou 3!")
	} else {
		fmt.Println("Não, X não multiplo é 2 ou 3!")
	}
}

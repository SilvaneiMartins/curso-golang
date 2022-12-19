package main

import "fmt"

func main() {
	// x := 4

	// if x == 2 || x == 3 {
	// 	fmt.Println("Sim, X é 2 ou 3!")
	// } else {
	// 	fmt.Println("Não, X não é 2 ou 3!")
	// }

	x := 9

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("X é multiplo de 2 ou 3!")
	} else {
		fmt.Println("Não, X não multiplo é 2 ou 3!")
	}
}

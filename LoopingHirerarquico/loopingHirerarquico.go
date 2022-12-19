package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Print("Horas: ", horas)

		for minutos := 0; minutos < 60; minutos++ {
			fmt.Print("Minutos: ", minutos)

			fmt.Print("\n")
		}
	}
}

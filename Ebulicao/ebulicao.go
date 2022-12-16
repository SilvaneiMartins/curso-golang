package main

import "fmt"

// Constantes
const ebulicaoF = 212

// Função principal
func main() {
	tempFF := 32
	tempFC := 0

	// Variável do valor da temperatura em graus F
	tempF := ebulicaoF

	// converta °F para °C
	tempC := (tempF - 32) * 5 / 9

	// Apareça na tela o resultado
	fmt.Printf("A temperatura de ebulição do água é °F ou %v (%T), temperatura de ebulição da água em °C = %v (%T).", tempF, tempF, tempC, tempC)
	fmt.Printf("A temperatura de funsão do água em °F ou %v (%T), temperatura de fusão da água em °C é = %v (%T).", tempFF, tempFF, tempFC, tempFC)

	// Esperamos que apareça F = 212 e C = 100
}

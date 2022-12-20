// Interface são coleções
// de métodos.
package main

import (
	"fmt"
	"math" // É utilizado para calcular a raiz quadrada.
)

// Interface "geometria" com dois métodos.
type geometria interface {
	area() float64
	perimetro() float64
}

// Para a aula usaremos a interface dos
// tipos QUADRADO e CIRCULO.
// Área de um quadrado: lado/Elevado ao quadrado ou Lado * Lado
// Perímetro de um quadrado: lado * 4
type quadrado struct {
	lado float64
}

// Área de um círculo: (Pi)*raio/Elevado ao quadrado perimetro de um círculo: 2*r*(pi)
type circulo struct {
	raio float64
}

// Para usar a interface em GO, so precisamos usar todos os metodos
// da interface. Aqui nós usaremos "geometria" em "quadrado" s.
func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

// Implementação da interface do "circulo".
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// Se uma variável tem uma interface como tipo, então podemos chamar
// métodos que estão na interface. Aqui está um exemplo de um "função"
// que aceita qualquer tipo que implemente a interface "geometria".
func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	q := quadrado{lado: 25}
	c := circulo{raio: 100}

	// Os tipos de estrutura "quadrado" e "circulo" implementam
	// a interface "geometria", então podemos usar instâncias
	// desses tipos como argumentos para "medir".

	medir(q)
	medir(c)
}

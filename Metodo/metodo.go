package main

// A struct é uma coleção de campos.
type retangulo struct {
	comprimento, altura int
}

// Essa função "area" tem um "receiver" do tipo retangulo.
func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

// Essa função "perimetro" também tem um "receiver" do tipo retangulo.
func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

// Essa função "main" chama as funções "area" e "perimetro" usando a struct "retangulo".
func main() {
	r := retangulo{comprimento: 50, altura: 25}

	println("Área:", r.area())
	println("Perímetro:", r.perimetro())
}

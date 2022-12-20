package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println(Pessoa{"Silvanei Martins", 44})
	fmt.Println(Pessoa{"João Martins", 72})
}

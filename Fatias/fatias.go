package main

import "fmt"

func main() {
	fatia1 := []string{"Palmeiras", "Conrinthias", "São Paulo", "Flamento", "Fluminense", "Internacional", "Gremio"}
	fatia2 := make([]string, 2)
	copy(fatia2, fatia1)
	fmt.Println(fatia1, fatia2)
}

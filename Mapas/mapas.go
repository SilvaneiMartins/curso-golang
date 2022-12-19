package main

import "fmt"

func main() {

	// x := make(map[string]int)
	// x["Chave"] = 10
	// fmt.Println(x["Chave"])

	// x := make(map[int]int)
	// x[1] = 20
	// x[2] = 30
	// fmt.Println(x[1], x[2])

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	elemento["Be"] = "Berílio"

	fmt.Println(elemento["Li"], elemento["Be"], elemento["H"], elemento["He"])
}

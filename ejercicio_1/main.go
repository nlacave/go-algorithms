package main

import (
	"fmt"

	"github.com/nlacave/go-algorithms/ejercicio_1/anagrama"
)

func main() {
	p := anagrama.EsAnagrama("manzana", "zanmana")
	fmt.Println(p)

	p = anagrama.EsAnagramaEficiente("manzana", "zanmana")
	fmt.Println(p)
}

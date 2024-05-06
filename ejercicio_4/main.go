/*
 * Reto #4
 * ÁREA DE UN POLÍGONO
 * Dificultad: FÁCIL
 *
 * Enunciado: Crea UNA ÚNICA FUNCIÓN (importante que sólo sea una) que sea capaz de calcular y retornar el área de un polígono.
 * - La función recibirá por parámetro sólo UN polígono a la vez.
 * - Los polígonos soportados serán Triángulo, Cuadrado y Rectángulo.
 * - Imprime el cálculo del área de un polígono de cada tipo.
 *
 */

package main

import (
	"fmt"
	"strings"
)

type Poligono struct {
	tipo   string
	base   int
	altura int
}

func main() {
	p := Poligono{"cuadrado", 4, 4}
	area := calcularAreaPoligono(p)
	fmt.Println(area)
	printArea(p)
}

func printArea(p Poligono) {
	switch p.tipo {
	case strings.ToLower("triangulo"), strings.ToLower("triángulo"):
		fmt.Println("El área del triángulo es", calcularAreaPoligono(p))
	case strings.ToLower("rectangulo"), strings.ToLower("rectángulo"):
		fmt.Println("El área del rectángulo es", calcularAreaPoligono(p))
	case strings.ToLower("cuadrado"):
		if p.base != p.altura {
			fmt.Println("El poligono proporcionado no es un cuadrado.")
			return
		}
		fmt.Println("El área del cuadrado es", calcularAreaPoligono(p))
	default:
		fmt.Println("Tipo no reconocido")
	}
}

func calcularAreaPoligono(p Poligono) int {
	if strings.ToLower(p.tipo) == "triangulo" || strings.ToLower(p.tipo) == "triángulo" {
		return (p.altura * p.base) / 2
	} else if strings.ToLower(p.tipo) == "cuadrado" {
		if p.base != p.altura {
			return -1
		}
		return p.altura * p.base
	} else if strings.ToLower(p.tipo) == "rectangulo" || strings.ToLower(p.tipo) == "rectángulo" {
		return p.altura * p.base
	} else {
		return -1
	}
}

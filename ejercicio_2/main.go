package main

import "fmt"

/*
 * Reto #2
 * LA SUCESIÓN DE FIBONACCI
 * Dificultad: DIFÍCIL
 *
 * Enunciado: Escribe un programa que imprima los 50 primeros números de la sucesión de Fibonacci empezando en 0.
 * La serie Fibonacci se compone por una sucesión de números en la que el siguiente siempre es la suma de los dos anteriores.
 * 0, 1, 1, 2, 3, 5, 8, 13...
 */

func main() {
	v := 0
	orig := 0
	fin := 1
	for i := 0; i < 50; i++ {
		if i == 0 {
			fmt.Println(v)
		} else {
			fmt.Println(fin)
			v = orig
			orig = fin
			fin = v + orig
		}
	}
}

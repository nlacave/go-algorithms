package main

import "fmt"

/*
 * Reto #3
 * ¿ES UN NÚMERO PRIMO?
 * Dificultad: MEDIA
 *
 * Enunciado: Escribe un programa que se encargue de comprobar si un número es o no primo.
 * Hecho esto, imprime los números primos entre 1 y 100.
 */
func main() {
	for i := 1; i <= 100; i++ {
		if esPrimo(i) {
			fmt.Println(i)
		}
	}
}

func esPrimo(num int) bool {
	var esPrimo = true
	if num < 2 {
		esPrimo = false
	}
	divisor := 2
	for divisor < num && divisor > 1 {
		if num%divisor == 0 {
			esPrimo = false
			break
		}
		divisor++
	}
	return esPrimo
}

/*
 * Reto #7
 * CONTANDO PALABRAS
 * Fecha publicación enunciado: 14/02/22
 * Fecha publicación resolución: 21/02/22
 * Dificultad: MEDIA
 *
 * Enunciado: Crea un programa que cuente cuantas veces se repite cada palabra y que muestre el recuento final de todas ellas.
 * - Los signos de puntuación no forman parte de la palabra.
 * - Una palabra es la misma aunque aparezca en mayúsculas y minúsculas.
 * - No se pueden utilizar funciones propias del lenguaje que lo resuelvan automáticamente.
 */
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	contarPalabras("Hola amigos, mi nombre es Brais. Si, me llamo brais")
}

func contarPalabras(texto string) {
	textoMinuscula := strings.ToLower(texto)
	letras := strings.Split(textoMinuscula, "")
	var nuevoTexto string
	mapa := make(map[string]int)
	for _, v := range letras {
		if unicode.IsLetter([]rune(v)[0]) {
			nuevoTexto += v
		} else {
			if v != "" {
				_, ok := mapa[nuevoTexto]
				if ok {
					mapa[nuevoTexto]++
				} else {
					mapa[nuevoTexto] = 1
				}
				nuevoTexto = ""
			}
		}
	}
	for k, v := range mapa {
		fmt.Println("La palabra", k, "tiene", v, "repeticiones.")
	}
}

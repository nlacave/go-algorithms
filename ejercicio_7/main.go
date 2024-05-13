/*
 * Reto #7
 * CONTANDO PALABRAS
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
	palabrasRepetidas("Hola amigos, mi nombre es Brais, si, me llamo brais. Eres parte de mis AMigos, verdad?")
}

func palabrasRepetidas(frase string) {
	fraseEnMinuscula := strings.ToLower(frase)
	mapa := make(map[string]int)
	fraseFiltrada := strings.FieldsFunc(fraseEnMinuscula, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	for _, v := range fraseFiltrada {
		mapa[v]++
	}

	for k, v := range mapa {
		fmt.Printf("Para la palabra '%v' tenemos %v repeticiones\n", k, v)
	}
}

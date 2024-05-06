/*
* Reto #6
* INVIRTIENDO CADENAS
* Dificultad: FÁCIL
*
* Enunciado: Crea un programa que invierta el orden de una cadena de texto sin usar funciones propias del lenguaje que lo hagan de forma automática.
* - Si le pasamos "Hola mundo" nos retornaría "odnum aloH"
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	invertirCadena("Fijate pues")
}

func invertirCadena(cadena string) {
	var arr []string
	cadenaNueva := strings.Split(cadena, "")
	for i := len(cadenaNueva) - 1; i >= 0; i-- {
		arr = append(arr, cadenaNueva[i])
	}
	cadenaRetornada := strings.Join(arr, "")
	fmt.Println(cadenaRetornada)
}

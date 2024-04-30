package anagrama

import (
	"sort"
	"strings"
)

/* Reto #1
 * ¿ES UN ANAGRAMA?
 * Dificultad: MEDIA
 *
 * Enunciado: Escribe una función que reciba dos palabras (String) y retorne verdadero o falso (Boolean) según sean o no anagramas.
 * Un Anagrama consiste en formar una palabra reordenando TODAS las letras de otra palabra inicial.
 * NO hace falta comprobar que ambas palabras existan.
 * Dos palabras exactamente iguales no son anagrama.
 */

func EsAnagramaEficiente(str1, str2 string) bool {
	if str1 == str2 || len(str1) != len(str2) {
		return false
	}
	str1Slice := strings.Split(strings.ToLower(str1), "")
	str2Slice := strings.Split(strings.ToLower(str2), "")
	sort.Strings(str1Slice)
	sort.Strings(str2Slice)
	return strings.Join(str1Slice, "") == strings.Join(str2Slice, "")
}

func EsAnagrama(str1 string, str2 string) bool {
	if str1 != str2 && len(str1) == len(str2) {
		sliceStr1 := strings.Split(str1, "")
		sliceStr2 := strings.Split(str2, "")
		cont := 0
		for _, v := range sliceStr2 {
			for i, k := range sliceStr1 {
				if string(v) == k {
					cont++
					sliceStr1 = append(sliceStr1[:i], sliceStr1[i+1:]...)
					break
				}
			}
		}
		if cont == len(sliceStr2) {
			return true
		}
	}
	return false
}

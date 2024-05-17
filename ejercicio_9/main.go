package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	listadoFechas := []string{"1st Mar 1974", "22nd Jan 2013", "7th Apr 1904"}
	f := preprocesarFechas(listadoFechas)
	fmt.Println(f)
}

func preprocesarFechas(fechas []string) []string {
	var fechasFormateadas []string
	for _, fecha := range fechas {
		arrFecha := strings.Fields(fecha)
		var posicion string
		var dia string
		mapaPosicion := map[int]string{
			1:  "st",
			21: "st",
			31: "st",
			2:  "nd",
			22: "nd",
			3:  "rd",
			23: "rd",
		}
		th := "th"
		for _, v := range arrFecha[0] {
			if !unicode.IsLetter(v) {
				dia += string(v)
			} else {
				posicion += string(v)
			}
		}
		diaAEntero, err := strconv.Atoi(dia)
		if err != nil {
			panic(err)
		}
		if diaAEntero >= 1 && diaAEntero <= 9 {
			dia = fmt.Sprintf("0%v", dia)
		}

		mapaMeses := map[string]string{
			"Jan": "01",
			"Feb": "02",
			"Mar": "03",
			"Apr": "04",
			"May": "05",
			"Jun": "06",
			"Jul": "07",
			"Aug": "08",
			"Sep": "09",
			"Oct": "10",
			"Nov": "11",
			"Dec": "12",
		}
		mes := mapaMeses[arrFecha[1]]
		anio := arrFecha[2]
		anioEntero, err := strconv.Atoi(anio)
		if err != nil {
			panic(err)
		}

		if (mapaPosicion[diaAEntero] == posicion || th == posicion) && mes != "" && anioEntero >= 1900 && anioEntero <= 2100 {
			fechasFormateadas = append(fechasFormateadas, fmt.Sprintf("%v-%v-%v", anio, mes, dia))
		} else {
			fechasFormateadas = append(fechasFormateadas, "Fecha Invalida")
		}
	}
	return fechasFormateadas
}

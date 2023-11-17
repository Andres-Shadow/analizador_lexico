package main

import (
	"fmt"
	"proyecto_tlf/automatas"
	"strings"
)

func main() {
	texto := "$nombre mi socio$ .=       ms pr pr pr   aaaa"
	var txtOriginal string
	var recorrido int
	txtOriginal = texto
	fin := len(txtOriginal)

	texto = strings.TrimSpace(texto)
	for i := 0; i < 100; i++ {
		//fmt.Println("evaluando: -", texto, "-")
		tamCadena := esCadena(texto)
		if tamCadena != len(texto) {
			recorrido += tamCadena
			texto = txtOriginal[tamCadena:fin]
		} else {
			tamIgualacion := esIgualacion(texto)
			if tamIgualacion != len(texto) {
				recorrido += tamIgualacion
				texto = txtOriginal[tamIgualacion:fin]
			} else {
				tamAritmetico := esAritmetico(texto)
				if tamAritmetico != len(texto) {
					recorrido += tamAritmetico
					texto = txtOriginal[tamAritmetico:fin]
				} else {
					recorrido++
					texto = txtOriginal[recorrido:fin]
				}

			}
		}

		if recorrido == len(txtOriginal)-1 {
			fmt.Println("no hay más palabras")
			break
		}
	}
}

func esCadena(texto string) int {
	valido, tam, _ := automatas.EvaluarCadena(texto)
	if valido {
		return tam
	} else {
		return len(texto)
	}
}

func esIgualacion(texto string) int {
	valido, tam, _ := automatas.EvaluarIgualacion(texto)
	if valido {
		return tam
	} else {
		return len(texto)
	}
}

func esAritmetico(texto string) int {
	valido, tam := automatas.EvaluarAritmetico(texto)
	if valido {
		return tam
	} else {
		return len(texto)
	}
}

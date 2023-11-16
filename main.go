package main

import (
	"fmt"
	"proyecto_tlf/automatas"
	"strings"
)

func main() {
	var texto string
	var result bool
	//texto = "n22n;r22,0r;mQ;pr;&"
	texto = "msms;mnmn;prpr;dvdv"
	segmentacion := strings.Split(texto, ";")

	for i := 0; i < len(segmentacion); i++ {
		cadena := segmentacion[i]
		fmt.Println("----------------------------")
		fmt.Println("evaluando la palabra: ", cadena)

		result = automatas.EvaluarIncremento(cadena)
		fmt.Println(result)
		/*result = automatas.EvaluarAutomata(cadena)
		fmt.Println("natural: ", result)
		result = automatas.EvaluarReales(cadena)
		fmt.Println("real: ", result)
		result = automatas.EvaluarComparacion(cadena)
		fmt.Println("comparacion: ", result)
		result = automatas.EvaluarAritmetico(cadena)
		fmt.Println("aritmetico: ", result)
		result = automatas.EvaluarLogico(cadena)
		fmt.Println("logico: ", result)*/

	}
	//a := automatas.EvaluarAutomata(verificar)
}

package main

import (
	"fmt"
	"proyecto_tlf/automatas"
)

func main() {
	texto := "$nombre mi socio$ .= "
	var txtOriginal string
	var recorrido int
	txtOriginal = texto

	continuar := true

	for continuar {
		fmt.Println("evaluando: ", texto)
		valido, tam, restante := automatas.EvaluarCadena(texto)
		if valido == true {
			fmt.Println("es una cadena: ", texto[:tam])
			recorrido = tam
			texto = restante
		} else {
			fmt.Println("no es una cadena: ", texto[:tam])
			recorrido ++
			texto = txtOriginal[recorrido:]
			
		}

		if recorrido == len(txtOriginal) {
			continuar = false
		}
	}
}

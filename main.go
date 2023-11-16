package main

import (
	"fmt"
	"proyecto_tlf/automatas"
)

func main() {
	verificar := "$a$"
	//a := automatas.EvaluarAutomata(verificar)
	a := automatas.EvaluarCadena(verificar)
	fmt.Print(a)
}

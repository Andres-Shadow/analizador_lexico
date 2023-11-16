package main

import (
	"fmt"
	"proyecto_tlf/automatas"
)

func main() {
	verificar := "pr"
	//a := automatas.EvaluarAutomata(verificar)
	a := automatas.EvaluarAutomataMultiplicacion(verificar)
	fmt.Print(a)
}

package main

import (
	"fmt"
	"proyecto_tlf/automatas"
)

func main() {
	verificar := ".="
	//a := automatas.EvaluarAutomata(verificar)
	a := automatas.EvaluarIgualacion(verificar)
	fmt.Print(a)
}

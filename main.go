package main

import (
	"fmt"
	"proyecto_tlf/automatas"
)

func main() {
	verificar := "mq"
	//a := automatas.EvaluarAutomata(verificar)
	a := automatas.EvaluarComparacion(verificar)
	fmt.Print(a)
}

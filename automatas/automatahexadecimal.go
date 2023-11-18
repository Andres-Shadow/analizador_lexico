package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoHexadecimal int

const (
	EstadoInicialHexadecimal EstadoHexadecimal = iota
	EstadoIntermedioHexadecimal
	EstadoFinalHexadecimal
	EstadoNoAceptadoHexadecimal
	EstadoErrorHexadecimal
)

type AutomataHexadecimal struct {
	estadoActual EstadoHexadecimal
}

func (a *AutomataHexadecimal) procesarHexadecimal(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialHexadecimal:
		if simbolo == 'h' {
			a.estadoActual = EstadoIntermedioHexadecimal
		} else {
			a.estadoActual = EstadoNoAceptadoHexadecimal
		}
	case EstadoIntermedioHexadecimal:
		if !esMayusculaValida(simbolo) {
			a.estadoActual = EstadoErrorHexadecimal
		} else if simbolo == 'h' {
			a.estadoActual = EstadoFinalHexadecimal
		} else {
			//se mantiene en el estado
		}
	case EstadoFinalHexadecimal:
		a.estadoActual = EstadoErrorHexadecimal

	}
}

func esMayusculaValida(simbolo rune) bool {
	return (simbolo >= 'A' && simbolo <= 'F')
}

func EvalaurHexadecimal(cadena string) (bool, int) {
	automata := &AutomataHexadecimal{}
	var iterator int

	for i, simbolo := range cadena {
		automata.procesarHexadecimal(simbolo)

		if automata.estadoActual == EstadoNoAceptadoHexadecimal {
			break
		}

		if automata.estadoActual == EstadoFinalHexadecimal {
			iterator = iterator + 1
			break
		}

		if i == len(cadena)-1 && i != 'h' {
			automata.estadoActual = EstadoErrorHexadecimal
			iterator = iterator + 1
			break
		}
	}

	if automata.estadoActual == EstadoFinalHexadecimal {
		contenido := cadena[:iterator] + " -> hexadecimal"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un cadena: ", cadena[:iterator])
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorHexadecimal {
		contenido := cadena[:iterator] + " -> error sintactivo hexadecimal"
		fmt.Println("error sintactico -> : ", cadena[:iterator])
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		return true, iterator
	}

	return false, len(cadena)
}

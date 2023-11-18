package automatas

import "fmt"

type EstadoSeparador int

const (
	EstadoInicialSeparador EstadoSeparador = iota
	EstadoFinalSeparador
	EstadoNoAceptadoSeparador
)

type AutomataSeparador struct {
	estadoActual EstadoSeparador
}

func (a *AutomataSeparador) procesarSeparador(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialSeparador:
		if simbolo == ',' {
			a.estadoActual = EstadoFinalSeparador
		} else {
			a.estadoActual = EstadoNoAceptadoSeparador
		}
	case EstadoFinalSeparador:
		a.estadoActual = EstadoNoAceptadoSeparador
	}
}

func EvaluarSeparador(cadena string) (bool, int) {
	automata := &AutomataSeparador{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarSeparador(simbolo)

		if automata.estadoActual == EstadoFinalSeparador {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoNoAceptadoSeparador {
			break
		}
	}

	if automata.estadoActual == EstadoFinalSeparador {
		fmt.Println("es un separador: ", cadena[:iterator])
		return true, iterator
	} else {
		return false, len(cadena)
	}

}

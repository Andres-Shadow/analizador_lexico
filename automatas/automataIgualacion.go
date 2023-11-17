package automatas

import "fmt"

type EstadoIgualacion int

const (
	EstadoIncialIgualacion EstadoIgualacion = iota
	EstadoIntermedioIgual
	EstadoFinalIgualacion
	EstadoErrorIgualacion
)

type AutomataIgualacion struct {
	estadoActual EstadoIgualacion
}

func (a *AutomataIgualacion) procesarIgualacion(simbolo rune) {
	switch a.estadoActual {
	case EstadoIncialIgualacion:
		if simbolo == '.' {
			a.estadoActual = EstadoIntermedioIgual
		} else {
			a.estadoActual = EstadoErrorIgualacion
		}
	case EstadoIntermedioIgual:
		if simbolo == '=' {
			a.estadoActual = EstadoFinalIgualacion
		} else {
			a.estadoActual = EstadoErrorIgualacion
		}
	case EstadoFinalIgualacion:
		a.estadoActual = EstadoErrorIgualacion
	}
}

func EvaluarIgualacion(cadena string) (bool, int, string) {
	automata := &AutomataIgualacion{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarIgualacion(simbolo)

		if automata.estadoActual == EstadoFinalIgualacion {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorIgualacion {
			break
		}
	}

	if automata.estadoActual == EstadoFinalIgualacion {
		fmt.Println("es un igualacion: ", cadena[:iterator])
		return true, iterator, cadena[iterator:]
	} else {
		return false, len(cadena), cadena
	}
}

package automatas
import "fmt"

type EstadoIncrementos int

const (
	EstadoInicialIncremento EstadoIncrementos = iota
	EstadoInicialMsMn
	EstadoIntermedioMs
	EstadoFinalMs
	EstadoIntermedioMn
	EstadoFinalMn
	EstadoInicialPr
	EstadoIntermedioPr
	EstadoFinalPr
	EstadoInicialDv
	EstadoIntermedioDv
	EstadoFinalDv
	EstadoFinalIncremento
	EstadoErrorIncremento
)

type AutomataIncremento struct {
	estadoActual EstadoIncrementos
}

func (a *AutomataIncremento) procesarIncremento(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialIncremento:
		if simbolo == 'm' {
			a.estadoActual = EstadoInicialMsMn
		} else if simbolo == 'p' {
			a.estadoActual = EstadoInicialPr
		} else if simbolo == 'd' {
			a.estadoActual = EstadoInicialDv
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoInicialMsMn:
		if simbolo == 's' {
			a.estadoActual = EstadoIntermedioMs
		} else if simbolo == 'n' {
			a.estadoActual = EstadoIntermedioMn
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoInicialPr:
		if simbolo == 'r' {
			a.estadoActual = EstadoIntermedioPr
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoInicialDv:
		if simbolo == 'v' {
			a.estadoActual = EstadoIntermedioDv
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoIntermedioMs:
		if simbolo == 'm' {
			a.estadoActual = EstadoFinalMs
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoIntermedioMn:
		if simbolo == 'm' {
			a.estadoActual = EstadoFinalMn
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoIntermedioPr:
		if simbolo == 'p' {
			a.estadoActual = EstadoFinalPr
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoIntermedioDv:
		if simbolo == 'd' {
			a.estadoActual = EstadoFinalDv
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoFinalMs:
		if simbolo == 's' {
			a.estadoActual = EstadoFinalIncremento
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoFinalMn:
		if simbolo == 'n' {
			a.estadoActual = EstadoFinalIncremento
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoFinalPr:
		if simbolo == 'r' {
			a.estadoActual = EstadoFinalIncremento
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoFinalDv:
		if simbolo == 'v' {
			a.estadoActual = EstadoFinalIncremento
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoFinalIncremento:
		a.estadoActual = EstadoErrorIncremento
	}
}

func EvaluarIncrementos(cadena string) (bool, int) {
	automata := &AutomataIncremento{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarIncremento(simbolo)

		if automata.estadoActual == EstadoFinalIncremento {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorIncremento {
			break
		}
	}

	if automata.estadoActual == EstadoFinalIncremento {
		fmt.Println("es un operador de incremento", cadena[:iterator])
		return true, iterator
	} else {
		return false, len(cadena)
	}
}

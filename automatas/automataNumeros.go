package automatas
import "fmt"

// Definición del tipo de estado
type EstadoNatural int

const (
	EstadoInicialNatural EstadoNatural = iota
	EstadoNumeroNatural
	EstadoFinalNatural
	EstadoErrorNatural
)

type AutomataNatural struct {
	estadoActual EstadoNatural
}

func (a *AutomataNatural) ProcesarSimbolo(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialNatural:
		if simbolo == 'n' {
			a.estadoActual = EstadoNumeroNatural
		} else {
			a.estadoActual = EstadoErrorNatural
		}
	case EstadoNumeroNatural:
		if esDigito(simbolo) {
			// Permanecer en el estado de número
		} else if simbolo == 'n' {
			a.estadoActual = EstadoFinalNatural
		} else {
			a.estadoActual = EstadoErrorNatural
		}
	case EstadoFinalNatural:
		a.estadoActual = EstadoErrorNatural
	}
}

func esDigitoNatural(simbolo rune) bool {
	return simbolo >= '0' && simbolo <= '9'
}

func EvaluarNatural(cadena string) (bool, int) {
	automata := &AutomataNatural{}
	var iterator int
	for i, simbolo := range cadena {
		automata.ProcesarSimbolo(simbolo)

		if automata.estadoActual == EstadoFinalNatural {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorNatural {
			break
		}
	}

	if automata.estadoActual == EstadoFinalNatural {
		fmt.Println("es un numero natural", cadena[:iterator])
		return true, iterator
	} else {
		return false, len(cadena)
	}
}
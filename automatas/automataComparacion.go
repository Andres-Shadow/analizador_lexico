package automatas
import "fmt"

type EstadoComparacion int

const (
	EstadoInicialComparacion EstadoComparacion = iota
	EstadoIntermedioComparacion
	EstadoIntermedioIgualacion
	EstadoFinalComparacion
	EstadoErrorComparacion
)

type AutomataComparacion struct {
	estadoActual EstadoComparacion
}

func (a *AutomataComparacion) procesarPalabra(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialComparacion:
		if simbolo == 'm' || simbolo == 'e' {
			a.estadoActual = EstadoIntermedioComparacion
		} else if simbolo == 'e' {
			a.estadoActual = EstadoIntermedioIgualacion
		} else {
			a.estadoActual = EstadoErrorComparacion
		}
	case EstadoIntermedioComparacion:
		if simbolo == 'Q' || simbolo == 'N' {
			a.estadoActual = EstadoFinalComparacion
		} else {
			a.estadoActual = EstadoErrorComparacion
		}
	case EstadoIntermedioIgualacion:
		if simbolo == 'Q' {
			a.estadoActual = EstadoFinalComparacion
		} else {
			a.estadoActual = EstadoErrorComparacion
		}
	case EstadoFinalComparacion:
		a.estadoActual = EstadoErrorComparacion
	}

}

func EvaluarComparacion(cadena string) (bool, int) {
	automata := &AutomataComparacion{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarPalabra(simbolo)

		if automata.estadoActual == EstadoFinalComparacion {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorComparacion {
			break
		}
	}

	if automata.estadoActual == EstadoFinalComparacion {
		fmt.Println("es un operador de comparacion", cadena[:iterator])
		return true, iterator
	} else {
		return false, len(cadena)
	}
}


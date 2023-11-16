package automatas

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

func EvaluarComparacion(cadena string) bool {
	automata := &AutomataComparacion{}

	for _, simbolo := range cadena {
		automata.procesarPalabra(simbolo)
	}

	if automata.estadoActual == EstadoFinalComparacion {
		return true
	} else {
		return false
	}
}

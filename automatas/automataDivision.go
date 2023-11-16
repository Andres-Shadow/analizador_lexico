package automatas

// Definición del tipo de estado
type EstadoDivision int

const (
	EstadoInicialDivision EstadoDivision = iota
	EstadoIntermedioDivision
	EstadoFinalDivision
	EstadoErrorDivision
)

type AutomataDivision struct {
	estadoActualDivision EstadoDivision
}

func (a *AutomataDivision) ProcesarSimboloDivision(simbolo rune) {
	switch a.estadoActualDivision {
	case EstadoInicialDivision:
		if simbolo == 'm' {
			a.estadoActualDivision = EstadoIntermedioDivision
		} else {
			a.estadoActualDivision = EstadoErrorDivision
		}
	case EstadoIntermedioDivision:
		if simbolo == 'n' {
			a.estadoActualDivision = EstadoFinalDivision
		} else {
			a.estadoActualDivision = EstadoErrorDivision
		}
	case EstadoFinalDivision:
		a.estadoActualDivision = EstadoErrorDivision
	}
}

func EvaluarAutomataDivision(cadena string) bool {
	automata := &AutomataDivision{}

	for _, simbolo := range cadena {
		automata.ProcesarSimboloDivision(simbolo)
	}

	return automata.estadoActualDivision == EstadoFinalDivision
}

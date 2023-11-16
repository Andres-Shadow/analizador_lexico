package automatas

// Definición del tipo de estado
type EstadoSuma int

const (
	EstadoInicialSuma EstadoSuma = iota
	EstadoIntermedioSuma
	EstadoFinalSuma
	EstadoErrorSuma
)

type AutomataSuma struct {
	estadoActualSuma EstadoSuma
}

func (a *AutomataSuma) ProcesarSimboloSuma(simbolo rune) {
	switch a.estadoActualSuma {
	case EstadoInicialSuma:
		if simbolo == 'm' {
			a.estadoActualSuma = EstadoIntermedioSuma
		} else {
			a.estadoActualSuma = EstadoErrorSuma
		}
	case EstadoIntermedioSuma:
		if simbolo == 's' {
			a.estadoActualSuma = EstadoFinalSuma
		} else {
			a.estadoActualSuma = EstadoErrorSuma
		}
	case EstadoFinalSuma:
		a.estadoActualSuma = EstadoErrorSuma
	}
}

func EvaluarAutomataSuma(cadena string) bool {
	automata := &AutomataSuma{}

	for _, simbolo := range cadena {
		automata.ProcesarSimboloSuma(simbolo)
	}

	return automata.estadoActualSuma == EstadoFinalSuma
}

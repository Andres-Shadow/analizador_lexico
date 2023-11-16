package automatas

// Definición del tipo de estado
type EstadoMultiplicacion int

const (
	EstadoInicialMultiplicacion EstadoMultiplicacion = iota
	EstadoIntermedioMultiplicacion
	EstadoFinalMultiplicacion
	EstadoErrorMultiplicacion
)

type AutomataMultiplicacion struct {
	estadoActualMultiplicacion EstadoMultiplicacion
}

func (a *AutomataMultiplicacion) ProcesarSimboloMultiplicacion(simbolo rune) {
	switch a.estadoActualMultiplicacion {
	case EstadoInicialMultiplicacion:
		if simbolo == 'p' {
			a.estadoActualMultiplicacion = EstadoIntermedioMultiplicacion
		} else {
			a.estadoActualMultiplicacion = EstadoErrorMultiplicacion
		}
	case EstadoIntermedioMultiplicacion:
		if simbolo == 'r' {
			a.estadoActualMultiplicacion = EstadoFinalMultiplicacion
		} else {
			a.estadoActualMultiplicacion = EstadoErrorMultiplicacion
		}
	case EstadoFinalMultiplicacion:
		a.estadoActualMultiplicacion = EstadoErrorMultiplicacion
	}
}

func EvaluarAutomataMultiplicacion(cadena string) bool {
	automata := &AutomataMultiplicacion{}

	for _, simbolo := range cadena {
		automata.ProcesarSimboloMultiplicacion(simbolo)
	}

	return automata.estadoActualMultiplicacion == EstadoFinalMultiplicacion
}

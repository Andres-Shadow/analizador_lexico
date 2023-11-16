package automatas

// Definición del tipo de estado
type EstadoResta int

const (
	EstadoInicialResta EstadoResta = iota
	EstadoIntermedioResta
	EstadoFinalResta
	EstadoErrorResta
)

type AutomataResta struct {
	estadoActualResta EstadoResta
}

func (a *AutomataResta) ProcesarSimboloResta(simbolo rune) {
	switch a.estadoActualResta {
	case EstadoInicialResta:
		if simbolo == 'm' {
			a.estadoActualResta = EstadoIntermedioResta
		} else {
			a.estadoActualResta = EstadoErrorResta
		}
	case EstadoIntermedioResta:
		if simbolo == 'n' {
			a.estadoActualResta = EstadoFinalResta
		} else {
			a.estadoActualResta = EstadoErrorResta
		}
	case EstadoFinalResta:
		a.estadoActualResta = EstadoErrorResta
	}
}

func EvaluarAutomataResta(cadena string) bool {
	automata := &AutomataResta{}

	for _, simbolo := range cadena {
		automata.ProcesarSimboloResta(simbolo)
	}

	return automata.estadoActualResta == EstadoFinalResta
}

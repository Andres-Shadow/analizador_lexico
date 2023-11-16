package automatas

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
		if simbolo == '='{
			a.estadoActual = EstadoFinalIgualacion
		}else{
			a.estadoActual = EstadoErrorIgualacion
		}
	case EstadoFinalIgualacion:
		a.estadoActual = EstadoErrorIgualacion
	}
}

func EvaluarIgualacion(cadena string) bool {
	automata := &AutomataIgualacion{}
	for _, simbolo := range cadena {
		automata.procesarIgualacion(simbolo)
	}

	if automata.estadoActual == EstadoFinalIgualacion {
		return true
	} else {
		return false
	}
}

package automatas

type EstadoLogico int

const (
	EstadoIncialLogico EstadoLogico = iota
	EstadoFinalLogico
	EstadoErrorLogico
)

type AutomataLogico struct {
	estadoActual EstadoLogico
}

func (a *AutomataLogico) procesarLogico(simbolo rune) {
	switch a.estadoActual {
	case EstadoIncialLogico:
		if simbolo == '&' {
			a.estadoActual = EstadoFinalLogico
		} else if simbolo == '|' {
			a.estadoActual = EstadoFinalLogico
		} else {
			a.estadoActual = EstadoErrorLogico
		}
	case EstadoFinalLogico:
		a.estadoActual = EstadoErrorLogico
	}
}

func EvaluarLogico(cadena string) bool {
	automata := &AutomataLogico{}
	for _, simbolo := range cadena {
		automata.procesarLogico(simbolo)
	}

	if automata.estadoActual == EstadoFinalLogico {
		return true
	} else {
		return false
	}
}

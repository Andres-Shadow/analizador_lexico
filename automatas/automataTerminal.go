package automatas

type EstadoTermial int

const (
	EstadoInicialTerminal EstadoTermial = iota
	EstadoIntermedioTerminal
	EstadoFinalTermianl
	EstadoErrorTerminal
)

type AutomataTerminal struct {
	estadoActual EstadoTermial
}

func (a *AutomataTerminal) procesarTerminal(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialTerminal:
		if simbolo == ';' {
			a.estadoActual = EstadoFinalTermianl
		} else if simbolo == '\n' {
			a.estadoActual = EstadoFinalTermianl
		} else {
			a.estadoActual = EstadoErrorTerminal
		}
	case EstadoFinalTermianl:
		a.estadoActual = EstadoErrorTerminal
	}
}

func EvaluarTerminal(cadena string) bool {
	automata := &AutomataTerminal{}
	for _, simbolo := range cadena {
		automata.procesarTerminal(simbolo)
	}

	if automata.estadoActual == EstadoFinalTermianl {
		return true
	} else {
		return false
	}
}

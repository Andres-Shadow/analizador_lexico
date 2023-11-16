package automatas

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

func EvaluarAutomata(cadena string) bool {
	automata := &AutomataNatural{}

	for _, simbolo := range cadena {
		automata.ProcesarSimbolo(simbolo)
	}

	if automata.estadoActual == EstadoFinalNatural {
		return true
	} else {
		return false
	}
}

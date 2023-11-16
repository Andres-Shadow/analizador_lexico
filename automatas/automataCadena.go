package automatas

type EstadoCadena int

const (
	EstadoInicialCadena EstadoCadena = iota
	EstadoIntermedioCadena
	EstadoFinalCadena
	EstadoErrorCadena
)

type AutomataCadena struct {
	estadoActual EstadoCadena
}

func (a *AutomataCadena) procesarPalabra(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialCadena:
		if simbolo == '$' {
			a.estadoActual = EstadoIntermedioCadena
		} else {
			a.estadoActual = EstadoErrorCadena
		}
	case EstadoIntermedioCadena:
		if esCaracterValido(simbolo) {
			//sigue en el mismo estado
		} else if simbolo == '$' {
			a.estadoActual = EstadoFinalCadena
		} else {
			a.estadoActual = EstadoErrorCadena
		}
	case EstadoFinalCadena:
		a.estadoActual = EstadoErrorCadena
	}
}

func EvaluarCadena(cadena string) bool {
	automata := &AutomataCadena{}

	for _, simbolo := range cadena {
		automata.procesarPalabra(simbolo)
	}

	if automata.estadoActual == EstadoFinalCadena {
		return true
	} else {
		return false
	}
}

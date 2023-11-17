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

func EvaluarCadena(cadena string) (bool, int, string) {
	automata := &AutomataCadena{}
	var iterator int

	for i, simbolo := range cadena {
		automata.procesarPalabra(simbolo)

		// Agregar tu condición de parada aquí
		if automata.estadoActual == EstadoFinalCadena {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorCadena {
			break
		}
	}

	if automata.estadoActual == EstadoFinalCadena {
		return true, iterator, cadena[iterator:]
	} else {
		return false, len(cadena), cadena
	}
}

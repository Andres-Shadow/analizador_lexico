package automatas

type EstadoReales int

const (
	EstadoInicial EstadoReales = iota
	EstadoR
	EstadoNumeros
	EstadoComa
	EstadoDespuesComa
	EstadoFinal
	EstadoError
)

type AutomataNumeroReal struct {
	estadoActual EstadoReales
}

func (a *AutomataNumeroReal) ProcesarSimbolo(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicial:
		if simbolo == 'r' {
			a.estadoActual = EstadoR
		} else {
			a.estadoActual = EstadoError
		}
	case EstadoR:
		if esDigito(simbolo) {
			a.estadoActual = EstadoNumeros
		} else {
			a.estadoActual = EstadoError
		}
	case EstadoNumeros:
		if esDigito(simbolo) {
			// Permanecer en el estado de números
		} else if simbolo == ',' {
			a.estadoActual = EstadoComa
		} else {
			a.estadoActual = EstadoError
		}
	case EstadoComa:
		if simbolo == '0' || esDigito(simbolo) {
			a.estadoActual = EstadoDespuesComa
		} else {
			a.estadoActual = EstadoError
		}
	case EstadoDespuesComa:
		if esDigito(simbolo) {
			// Permanecer en el estado después de la coma
		} else if simbolo == 'r' {
			a.estadoActual = EstadoFinal
		} else {
			a.estadoActual = EstadoError
		}
	case EstadoFinal:
		a.estadoActual = EstadoError
	}
}

func esDigito(simbolo rune) bool {
	return simbolo >= '0' && simbolo <= '9'
}

func EvaluarReales(cadena string) bool {

	automata := &AutomataNumeroReal{}

	for _, simbolo := range cadena {
		automata.ProcesarSimbolo(simbolo)
	}

	if automata.estadoActual == EstadoFinal {
		return true
	} else {
		return false
	}
}

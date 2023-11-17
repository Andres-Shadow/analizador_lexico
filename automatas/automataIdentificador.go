package automatas
import "fmt"

// Definición del tipo de estado
type EstadoS int

const (
	EstadoInicialS EstadoS = iota
	EstadoSegundaS
	EstadoIntermedioS
	EstadoPreFinalS
	EstadoFinalS
	EstadoErrorS
)

type AutomataS struct {
	estadoActual EstadoS
	contador     int
}

func (a *AutomataS) ProcesarSimbolo(simbolo rune) {

	switch a.estadoActual {
	case EstadoInicialS:
		if simbolo == '$' {
			a.estadoActual = EstadoSegundaS
		} else {
			a.estadoActual = EstadoErrorS
		}
	case EstadoSegundaS:
		if simbolo == '$' {
			a.estadoActual = EstadoIntermedioS
			a.contador = 1
		} else {
			a.estadoActual = EstadoErrorS
		}
	case EstadoIntermedioS:
		if esCaracterValido(simbolo) {
			if a.contador <= 10 {
				// Permanecer en el estado intermedio mientras no se superen los 10 caracteres
				a.contador++
			} else {
				a.estadoActual = EstadoErrorS
			}
		} else if simbolo == '$' && a.contador <= 10 {
			a.estadoActual = EstadoPreFinalS
		}
	case EstadoPreFinalS:
		if simbolo == '$' {
			a.estadoActual = EstadoFinalS
		} else {
			a.estadoActual = EstadoErrorS
		}
	case EstadoFinalS:
		a.estadoActual = EstadoErrorS
	}
}

func esCaracterValido(simbolo rune) bool {
	// Puedes ajustar la lógica según tus necesidades
	return (simbolo >= 'a' && simbolo <= 'z') || (simbolo >= 'A' && simbolo <= 'Z') || simbolo == ' '
}

func EvaluarIdentificador(cadena string) (bool, int) {
	automata := &AutomataS{}
	var iterator int
	for i, simbolo := range cadena {
		automata.ProcesarSimbolo(simbolo)

		if automata.estadoActual == EstadoFinalS {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorS {
			break
		}
	}

	if automata.estadoActual == EstadoFinalS {
		fmt.Println("es un Identificador", cadena[:iterator])
		return true, iterator
	} else {
		return false, len(cadena)
	}
}
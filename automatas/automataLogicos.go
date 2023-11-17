package automatas
import "fmt"

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

func EvaluarLogico(cadena string) (bool, int) {
	automata := &AutomataLogico{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarLogico(simbolo)

		if automata.estadoActual == EstadoFinalLogico {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorLogico {
			break
		}
	}

	if automata.estadoActual == EstadoFinalLogico {
		fmt.Println("es un operador logico", cadena[:iterator])
		return true, iterator
	} else {
		return false, len(cadena)
	}
}

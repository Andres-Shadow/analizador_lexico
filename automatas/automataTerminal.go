package automatas
import "fmt"

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

func EvaluarTerminal(cadena string) (bool, int) {
	automata := &AutomataTerminal{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarTerminal(simbolo)

		if automata.estadoActual == EstadoFinalTermianl {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorTerminal {
			break
		}
	}

	if automata.estadoActual == EstadoFinalTermianl {
		fmt.Println("es un final de sentencia", cadena[:iterator])
		return true, iterator
	} else {
		return false, len(cadena)
	}
}


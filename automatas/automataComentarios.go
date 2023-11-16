package automatas

type EstadoComentario int

const (
	EstadoInicialComentarios EstadoComentario = iota
	//comentario de linea
	EstadoInicioComLinea
	EstadoIntermedioComLinea

	//comentario de bloque
	EstadoInicioComBlo
	EstadoComBloIntermedio
	EstadoComBloPosIntermedio

	//estado final
	EstadoFinalComentario
	EstadoErrorComentario
)

type AutomataComentario struct {
	estadoActual EstadoComentario
}

func (a *AutomataComentario) procesarComentario(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialComentarios:
		if simbolo == '#' {
			a.estadoActual = EstadoInicioComLinea
		} else if simbolo == '/' {
			a.estadoActual = EstadoInicioComBlo
		} else {
			a.estadoActual = EstadoErrorComentario
		}
	case EstadoInicioComLinea:
		terminal := EvaluarTerminal(string(simbolo))
		if terminal == true {
			a.estadoActual = EstadoFinalComentario
		} else {
			a.estadoActual = EstadoInicioComLinea
		}
	case EstadoInicioComBlo:
		if simbolo == '#' {
			a.estadoActual = EstadoComBloIntermedio
		} else {
			a.estadoActual = EstadoErrorComentario
		}
	case EstadoComBloIntermedio:
		if simbolo != '#' {
			//sigue en el mismo estado
		} else if simbolo == '#' {
			a.estadoActual = EstadoComBloPosIntermedio
		} else {
			a.estadoActual = EstadoErrorComentario
		}
	case EstadoComBloPosIntermedio:
		if simbolo == '/' {
			a.estadoActual = EstadoFinalComentario
		} else {
			a.estadoActual = EstadoErrorComentario
		}
	case EstadoFinalComentario:
		a.estadoActual = EstadoErrorComentario
	}

}

func EvaluarComentario(cadena string) bool {
	automata := &AutomataComentario{}
	for _, simbolo := range cadena {
		automata.procesarComentario(simbolo)
	}

	if automata.estadoActual == EstadoFinalComentario {
		return true
	} else {
		return false
	}
}

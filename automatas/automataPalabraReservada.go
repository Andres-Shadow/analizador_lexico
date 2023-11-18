package automatas

type estadoPalabraReservada int

const (
	EstadoInicialPalabraReservada estadoPalabraReservada = iota
	//si - sino
	EstadoInicioSi
	EstadoIntermedioSi
	EstadoFinalSi
	EstadoPosSi

	//mod max min
	EstadoInicialM
	EstadoInterMod
	EstadoFinalMod
	EstadoInterMax
	EstadoFinalMax
	EstadoInterMin
	EstadoFinalMin

	//print
	EstadoInicioPrint
	EstadoRRrint
	EstadoIPrint
	EstadoNPrint
	EstadoTPrint

	EstadoErrorPalabraReservada
	EstadoNoAceptadoPalabraReservada
)

type AutomataPR struct {
	estadoActual estadoPalabraReservada
}

/*
func (a *AutomataPR) procesarPalabraReservada(simbolo rune){
	switch a.estadoActual{
	case EstadoInicialPalabraReservada:
		if simbolo == 'm'{
			a.estadoActual = EstadoInicialM
		}else if simbolo == 's'{
			a.estadoActual = EstadoInicioSi
		}else if simbolo == 'p'{
			a.estadoActual = EstadoInicioPrint
		}
	case EstadoInicioSi:
		if simbolo == 'i'{
			a.estadoActual = EstadoFinalSi
		}else{
			a.estadoActual =
		}
	}
}

UM
YM
TH
FD
GD
VV*/

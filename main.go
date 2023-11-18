package main

import (
	"fmt"
	"proyecto_tlf/automatas"
)

func main() {
	//texto := "?nombre mi socio? .=       ms pr pr pr   aaaa & n45n"
	texto := "?nombre"
	//texto := utilities.LeerArchivo("./entrada.txt")
	var txtOriginal string
	var recorrido int
	recorrido = 0
	txtOriginal = texto
	fin := len(txtOriginal)

	//for i := 0; recorrido < fin; i++ {
	for i := 0; i < 2; i++ {
		if recorrido < fin {
			texto = txtOriginal[recorrido:fin]

			posicion := esCadena(texto)

			if posicion != -1 {
				recorrido += posicion
			} else {
				posicion = esAsignacion(texto)
				if posicion != -1 {
					recorrido += posicion
				} else {
					posicion = esAritmetico(texto)
					if posicion != -1 {
						recorrido += posicion
					} else {
						posicion = esLogico(texto)
						if posicion != -1 {
							recorrido += posicion
						} else {
							posicion = esNaturla(texto)
							if posicion != -1 {
								recorrido += posicion
							} else {
								//error
								if txtOriginal[recorrido] == ' ' {
									recorrido++
								} else {
									recorrido++
									fmt.Println("es un error: -", txtOriginal[recorrido-1:recorrido], " -")
								}

							}

						}

					}
				}
			}

		}

		//fmt.Println("recorrido = ", recorrido)

		/*fmt.Println("----------------------------------------------------------------------------------")
		fmt.Println("iteracion: ", i, " | palabra restante: -", texto, "- | recorrido va en: ", recorrido)
		fmt.Println("----------------------------------------------------------------------------------")*/
	}

}

func esCadena(texto string) int {
	valido, tam := automatas.EvaluarCadena(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esAsignacion(texto string) int {
	valido, tam := automatas.EvaluarIgualacion(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esAritmetico(texto string) int {
	valido, tam := automatas.EvaluarAritmetico(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esLogico(texto string) int {
	valido, tam := automatas.EvaluarLogico(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esNaturla(texto string) int {
	valido, tam := automatas.EvaluarNatural(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esReal(texto string) int {
	valido, tam := automatas.EvaluarReales(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esComparacion(texto string) int {
	valido, tam := automatas.EvaluarComparacion(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esIncremento(texto string) int {
	valido, tam := automatas.EvaluarIncrementos(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esTerminal(texto string) int {
	valido, tam := automatas.EvaluarTerminal(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esSeparador(texto string) int {
	valido, tam := automatas.EvaluarSeparador(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esHexadecimal(texto string) int {
	valido, tam := automatas.EvalaurHexadecimal(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esComentario(texto string) int {
	valido, tam := automatas.EvaluarComentario(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

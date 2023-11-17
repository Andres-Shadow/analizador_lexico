package main

import (
	"fmt"
	"proyecto_tlf/automatas"
)

func main() {
	texto := "?nombre mi socio? .=       ms pr pr pr   aaaa &"
	var txtOriginal string
	var recorrido int
	recorrido = 0
	txtOriginal = texto
	fin := len(txtOriginal)

	for i := 0; recorrido < fin; i++ {
		//for i := 0; i < 1; i++ {

		if recorrido < fin {
			texto = txtOriginal[recorrido:fin]

			/*fmt.Println("----------------------------------------------------------------------------------")
			fmt.Println("iteracion: ", i, " | palabra restante: -", texto, "- | recorrido va en: ", recorrido)
			fmt.Println("----------------------------------------------------------------------------------")*/

			posicion := esCadena(texto)

			if posicion != -1 {
				recorrido += posicion
			} else {
				posicion = esIgualacion(texto)
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

		//fmt.Println("recorrido = ", recorrido)

		/*fmt.Println("----------------------------------------------------------------------------------")
		fmt.Println("iteracion: ", i, " | palabra restante: -", texto, "- | recorrido va en: ", recorrido)
		fmt.Println("----------------------------------------------------------------------------------")*/
	}

}

/*func main() {
	texto := "$nombre mi socio$ .=       ms pr pr pr   aaaa"
	var txtOriginal string
	var recorrido int
	recorrido = 0
	txtOriginal = texto
	fin := len(txtOriginal)

	texto = strings.TrimSpace(texto)
	for i := 0; i < 100; i++ {
		// fmt.Println("evaluando: -", texto, "-")
		tamCadena := esCadena(texto)
		if tamCadena != -1 {
			recorrido += tamCadena
			texto = txtOriginal[recorrido:fin]
		} else {
			tamIgualacion := esIgualacion(texto)
			if tamIgualacion != len(texto) {
				recorrido += tamIgualacion
				texto = txtOriginal[recorrido:fin]
			} else {
				tamAritmetico := esAritmetico(texto)
				if tamAritmetico != len(texto) {
					recorrido += tamAritmetico
					texto = txtOriginal[recorrido:fin]
				} else {
					fmt.Println("se detectó un error: ", texto)
					recorrido++
					texto = txtOriginal[recorrido:fin]
				}
			}
		}

		if recorrido >= len(txtOriginal)-1 {
			fmt.Println("no hay más palabras")
			break
		}
	}
}*/

func esCadena(texto string) int {
	valido, tam, _ := automatas.EvaluarCadena(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esIgualacion(texto string) int {
	valido, tam, _ := automatas.EvaluarIgualacion(texto)
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

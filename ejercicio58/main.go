/*58. Usted ha sido contratado para formar parte de un equipo de inteligencia mega
secreto, todos los mensajes enviados por los integrantes del equipo deben ser codificados,
pero la comunicación se ha hecho muy lenta pues actualmente no usan ninguna
herramienta informática, por lo que usted tiene la misión de crear un programa que
reciba un String y lo codifique según el método “cenit polar”. Cada aparición de la letra “c”
es cambiada por la letra “p”, la “e” por la “o”, etc. Y luego debe imprimir el mensaje
codificado. De usted depende que su equipo no sea descubierto. Buena suerte.*/

package main

import "fmt"

func codification(phrase string) []byte {
	cenit := "cenit"
	polar := "polar"
	output := []byte{}
	for i := 0; i < len(phrase); i++ {
		output = append(output, phrase[i])
	}
	for i := 0; i < len(output); i++ {
		for n := 0; n < len(cenit); n++ {
			if output[i] == cenit[n] {
				output[i] = polar[n]
			} else if output[i] == polar[n] {
				output[i] = cenit[n]
			}
		}
	}
	return output
}

func main() {
	fmt.Println(string(codification("Juanito quiere comer")))
}

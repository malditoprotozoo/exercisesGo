/*26. Hacer un programa para calcular el máximo común divisor de dos números enteros
positivos N y M siguiendo el algoritmo de Euclídes, que es el siguiente:
1. Se divide N por M, sea R el resto.
2. Si R=0, el máximo común divisor es M y se acaba.
3. Se asigna a N el valor de M y a M el valor de R y volver al paso 1. */

package main

import "fmt"

func minDiv(n, m int) {
	r := n % m
	if r == 0 {
		fmt.Println(m)
	} else {
		minDiv(m, r)
	}
}

func main() {
	minDiv(2, 1231236)
}

/*32. Escribir un programa que sume los elementos de cada fila de una matriz de nxm */

package main

import "fmt"

func main() {
	matrix := [][]int{{3, 5, 1}, {43, 1, 9}, {0, 2, 3}}
	for _, outArr := range matrix {
		sum := 0
		for _, num := range outArr {
			sum += num
		}
		fmt.Println(sum)
	}
}

/*34. Escribir un programa que sume los elementos de la diagonal principal de una matriz de
nxm */

package main

import "fmt"

func main() {
	matrix := [][]int{
		{2, 3, 4, 6},
		{0, 1, 1, 3},
		{5, 1, 9, 0},
		{4, 21, 2, 3}}

	result := 0
	index := 0
	for i := 0; i < len(matrix); i++ {
		result += matrix[i][index]
		index++
	}
	fmt.Println(result)
}

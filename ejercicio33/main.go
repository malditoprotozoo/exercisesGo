/*33. Escribir un programa que sume los elementos de cada columna de una matriz de nxm */

package main

import "fmt"

func main() {
	matrix := [][]int{{3, 5, 1, 3}, {43, 1, 9, 2}, {0, 2, 3, 1}}
	result := []int{}
	for i := 0; i < len(matrix); i++ {
		for n := 0; n < len(matrix[i]); n++ {
			if len(result) < len(matrix[i]) {
				result = append(result, matrix[i][n])
			} else {
				result[n] += matrix[i][n]
			}
		}
	}
	fmt.Println(result)
}

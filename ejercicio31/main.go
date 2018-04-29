/*31. Escribir un programa que encuentre el mayor y menor ingresado dentro de una matriz
de mxn */

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {43, 837, 0}, {-12, 3, 2}}
	min := matrix[0][0]
	max := matrix[0][0]
	for _, outArr := range matrix {
		for _, num := range outArr {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
	}
	fmt.Println(min, max)
}

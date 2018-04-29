/*31. Escribir un programa que encuentre el mayor y menor ingresado dentro de una matriz
de mxn */

package main

import "fmt"

func main() {
	matrix := []int{
		3, 5, 7,
		9, 0, 1,
		43, 76, 2,
		19, -1, 233,
	}
	min := matrix[0]
	max := matrix[0]
	for _, element := range matrix {
		if element < min {
			min = element
		}
		if element > max {
			max = element
		}
	}
	fmt.Println(min, max)
}

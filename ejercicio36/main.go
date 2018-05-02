/*36. Escribnxir un programa que copie los elementos de los bordes en sentido horario de una
matriz de m en un vector resultante llamado v, imprima su contenido. */

package main

import "fmt"

func main() {
	matrix := [][]int{
		{2, 3, 4, 6},
		{1000, 1, 1, 3},
		{5, 3, 9, 0},
		{4, 21, 2, 3},
		{99, 11, 22, 33},
	}
	v := make([]int, (len(matrix)*2)+(len(matrix[0])*2)-4)
	counter := 0
	// RIGHT
	for i := 0; i < len(matrix[0]); i++ {
		counter++
		v[i] = matrix[0][i]
	}
	// DOWN
	for i := 1; i < len(matrix); i++ {
		counter++
		v[len(matrix[0])+i-1] = matrix[i][len(matrix[0])-1]
	}
	// LEFT
	for i := len(matrix[0]) - 2; i >= 0; i-- {
		counter++
		v[len(v)-len(matrix[0])-i] = matrix[len(matrix)-1][i]
	}
	// TOP
	for i := len(matrix) - 2; i >= 1; i-- {
		v[counter] = matrix[i][0]
		counter++
	}
	fmt.Println(v)
}

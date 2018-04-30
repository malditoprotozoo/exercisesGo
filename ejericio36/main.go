/*36. Escribir un programa que copie los elementos de los bordes en sentido horario de una
matriz de nxm en un vector resultante llamado v, imprima su contenido. */

package main

import "fmt"

func main() {
	matrix := [][]int{
		{2, 3, 4, 6},
		{0, 1, 1, 3},
		{5, 3, 9, 0},
		{4, 21, 2, 3},
		{99, 11, 22, 33},
	}
	v := []int{}
	firstArr, rowsLen, lastArr := 0, len(matrix[0])-1, len(matrix)-1
	// RIGHT
	for i := 0; i < len(matrix[firstArr]); i++ {
		v = append(v, matrix[firstArr][i])
	}
	// DOWN
	for i := 1; i < len(matrix); i++ {
		v = append(v, matrix[i][rowsLen])
	}
	// LEFT
	for i := len(matrix[firstArr]) - 2; i >= 0; i-- {
		v = append(v, matrix[lastArr][i])
	}
	// TOP
	for i := lastArr - 1; i >= 0; i-- {
		v = append(v, matrix[i][0])
	}
	fmt.Println(v)
}

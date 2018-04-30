/*37. Escribir un programa que sume e imprima cada elemento en una matriz de mxn, que
forme en ella la letra N */
package main

import "fmt"

func main() {
	matrix := [][]int{
		{2, 3, 4, 6},
		{0, 1, 1, 3},
		{5, 3, 9, 0},
		{4, 21, 2, 3},
	}
	result := []int{}
	first, diagonal, last := 0, 0, len(matrix[0])-1
	for i := 0; i < len(matrix); i++ {
		result = append(result, matrix[i][first])
	}
	for i := 0; i < len(matrix); i++ {
		result = append(result, matrix[i][diagonal])
		diagonal++
	}
	for i := 0; i < len(matrix); i++ {
		result = append(result, matrix[i][last])
	}
	fmt.Println(result, diagonal, last)
}

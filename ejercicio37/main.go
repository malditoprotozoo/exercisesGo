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
	first, diagonal, last, sum := 0, 1, len(matrix[0])-1, 0
	for i := 0; i < len(matrix); i++ {
		result = append(result, matrix[i][first])
		sum += matrix[i][first]
	}
	for i := 1; i < len(matrix); i++ {
		result = append(result, matrix[i][diagonal])
		sum += matrix[i][diagonal]
		diagonal++
	}
	for i := 0; i < len(matrix)-1; i++ {
		result = append(result, matrix[i][last])
		sum += matrix[i][last]
	}
	fmt.Println(result, sum)
}

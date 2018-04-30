/*38. Escribir un programa que al ingresar el usuario una matriz de nxn entregue la
traspuesta.*/

package main

import "fmt"

type row []int
type matrix []row

func printMatrix(m matrix) {
	for _, innerArr := range m {
		fmt.Println(innerArr)
	}
}

func transpose(m matrix) matrix {
	r := make(matrix, len(m[0]))
	for x := range r {
		r[x] = make(row, len(m))
	}
	for y, s := range m {
		for x, e := range s {
			r[x][y] = e
		}
	}
	return r
}

func main() {
	mat := matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	printMatrix(transpose(mat))
}

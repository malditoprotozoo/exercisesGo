/* 39. Escriba un programa que permita multiplicar dos matrices, la primera de nxm y la
segunda de mxr, donde n, m, r pueden o no ser diferentes.*/

package main

import "fmt"

func multiply(arr1, arr2 [][]int) {
	res := []int{}
	for i := 0; i < len(arr1); i++ {
		for n := 0; n < len(arr2[0]); n++ {
			sum := 0
			for j := 0; j < len(arr1[0]); j++ {
				sum += arr1[i][j] * arr2[j][n]
			}
			res = append(res, sum)
		}
	}
	fmt.Println(res)
}

func main() {
	arr1 := [][]int{
		{1, 2},
		{3, 4},
	}
	arr2 := [][]int{
		{1, 2, 4},
		{3, 4, 5},
	}
	multiply(arr1, arr2)
}

/* 63. Se sabe que en cada década existe un número que impera en las mentes de todos los
jóvenes, claro que si están pensando en alguna cosa que se encuentre relacionada con
otro número si les pregunta por un número no dirán el número de moda. Es por esto que
usted realizará un Encuestador, tendrá que permitir el ingreso del resultado de la
encuesta, que serán 25 números y deberá guardarlos en un arreglo, luego creará un
método que permita recibir el arreglo y que encuentre el número de moda, que será el
más repetido dentro del arreglo.*/

package main

import "fmt"

func contains(arr []int, n int) bool {
	for _, dig := range arr {
		if dig == n {
			return true
		}
	}
	return false
}

func findMode(arr []int) int {
	mode, count, singles := 0, 0, []int{}
	for i := 0; i < len(arr); i++ {
		if len(singles) != 0 {
			if contains(singles, arr[i]) {
				count++
				mode = arr[i]
			} else {
				singles = append(singles, arr[i])
			}
		} else {
			singles = append(singles, arr[i])
		}
	}
	return mode
}

func main() {
	arr := []int{2, 2, 2, 3, 4, 5, 6, 1, 2, 8, 1, 1, 1, 1, 1}
	fmt.Println(findMode(arr))
}

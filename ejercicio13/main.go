/* 13.-Hacer un programa que imprima el mayor y el menor de una serie de cinco números que
vamos introduciendo por teclado (si le ocurrió hacer una variable para cada número
ingresado, no hay problema; pero luego inténtelo usando una sola variable para ingresar).
*/

package main

import "fmt"

func main() {
	arr := [5]int{8, 5, 300, -6, 0}
	min := arr[0]
	max := arr[0]
	for _, element := range arr {
		if element < min {
			min = element
		}
		if element > max {
			max = element
		}
	}
	fmt.Println(min, max)
}

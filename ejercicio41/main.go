/*41. Cree un programa que reciba un numero n y retorne la suma de todos los números
menores o iguales que n y mayores que cero, sin usar ciclos ni formulas. */

package main

import "fmt"

func sumNumbers(n int) int {
	if n != 0 {
		return n + sumNumbers(n-1)
	}
	return n
}

func main() {
	fmt.Println("Insert a Number")
	var input int
	fmt.Scanln(&input)
	fmt.Println(sumNumbers(input))
}

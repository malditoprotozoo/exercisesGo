/*42. Implemente de forma recursiva, un programa que imprima el máximo común divisor
de dos números enteros utilizando el algoritmo de Euclides.
Dados dos números enteros positivos m y n, tal que m > n, para encontrar su máximo
común divisor, es decir, el mayor entero positivo que divide a ambos:
 Dividir m por n para obtener el resto r (0 r < n)
 Si r = 0, el MCD es n.
 Si no, el máximo común divisor es MCD(n,r). */

package main

import "fmt"

func mcd(x, y int) int {
	if x != 0 && y != 0 {
		return mcd(y, x%y)
	}
	return x
}

func main() {
	fmt.Println("Insert two numbers")
	var input int
	var input2 int
	fmt.Scanln(&input)
	fmt.Scanln(&input2)
	fmt.Println(mcd(input, input2))
}

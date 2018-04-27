/*27. Dados dos números enteros positivos N y D, se dice que D es un divisor de N si el resto
de dividir N entre D es 0. Se dice que un número N es perfecto si la suma de sus divisores
(excluido el propio N) es N. Por ejemplo 28 es perfecto, pues sus divisores (excluido el 28)
son: 1, 2, 4, 7 y 14 y su suma es 1+2+4+7+14=28. Cree un programa que reciba un número
N e indique si es perfecto. */

package main

import "fmt"

func main() {
	fmt.Println("Pick a number")
	var input int
	fmt.Scanln(&input)
	if input%(input/2) != 0 {
		fmt.Println("The number is not perfect")
	} else {
		res := 0
		for i := 1; i <= input/2; i++ {
			if input%i == 0 {
				res += i
			}
		}
		if res == input {
			fmt.Println("The number is perfect :D")
		} else {
			fmt.Println("The number is not perfect :(")
		}
	}
}

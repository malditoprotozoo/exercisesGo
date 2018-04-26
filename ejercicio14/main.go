/* 14.- Calcular el factorial de un número ingresado por el usuario (Recuerde: factorial(4) =
1*2*3*4, factorial(5) = 1*2*3*4*5, por lo que si un usuario ingresa 6 sería 1*2*3*4*5*6,
es decir, imprime 720) */

package main

import "fmt"

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
func main() {
	fmt.Println(factorial(6))
}

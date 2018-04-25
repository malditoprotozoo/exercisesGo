// 7.-Hacer un programa que imprima todos los números naturales que hay desde la unidad hasta un numero que introducimos por teclado.
package main

import (
	"fmt"
)

// Pendiente: Averiguar cómo ingresar inputs de usuario

func calc(num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}

func main() {
	calc(9990)
}

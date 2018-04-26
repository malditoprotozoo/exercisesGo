// 7.-Hacer un programa que imprima todos los números naturales que hay desde la unidad hasta un numero que introducimos por teclado.
package main

import (
	"fmt"
)

func calc(num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Wite a number")
	var input int
	fmt.Scanln(&input)
	calc(input)
}

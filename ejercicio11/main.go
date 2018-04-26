// 11.-Imprimir y contar los múltiplos de 3 desde la unidad hasta un número que introducimos por teclado.
package main

import "fmt"

func main() {
	fmt.Println("Write a number")
	var input int
	fmt.Scanln(&input)
	for i := 1; i <= input; i += 3 {
		fmt.Println(i)
	}
}

// 9.-Introducir un numero por teclado. Que nos diga si es positivo o negativo (considere el cero como positivo para este ejercicio).
package main

import "fmt"

func main() {
	fmt.Println("Write a Number")
	var input int
	fmt.Scanln(&input)
	if input >= 0 {
		fmt.Println("Your number is positive")
	} else if input < 0 {
		fmt.Println("Your number is negative")
	}
}

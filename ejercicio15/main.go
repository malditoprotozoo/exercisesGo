/* 15.-Introducir dos números por teclado. Imprimir los números naturales que hay entre
ambos números empezando por el más pequeño, contar cuantos hay y cuántos de ellos
son pares. Calcular la suma de los impares. */

package main

import "fmt"

func main() {
	fmt.Println("Pick Two Numbers")
	var input1, input2 int
	fmt.Scanln(&input1)
	fmt.Scanln(&input2)
	var min, max int
	if input1 > input2 {
		max = input1
		min = input2
	} else {
		max = input2
		min = input1
	}
	sum := 0
	even := 0
	odd := 0
	for i := min + 1; i < max; i++ {
		fmt.Println(i)
		sum++
		if i%2 == 0 {
			even++
		} else {
			odd += i
		}
	}
	fmt.Printf("There are %d numbers among %d and %d \n", sum, min, max)
	fmt.Printf("There are %d even numbers among %d and %d \n", even, min, max)
	fmt.Printf("The sum of the odd numbers among %d and %d is %d \n", min, max, odd)
}

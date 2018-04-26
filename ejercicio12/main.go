// 12.-Imprimir y contar los números que son múltiplos de 2 o de 3 que hay entre 1 y N.

package main

import "fmt"

func mult(n int) {
	sum := 0
	for i := 1; i < n; i++ {
		if i%2 == 0 || i%3 == 0 {
			sum++
			fmt.Println(i)
		}
	}
	fmt.Printf("There are %d multiples of 2 or 3 \n", sum)
}
func main() {
	mult(25)
}

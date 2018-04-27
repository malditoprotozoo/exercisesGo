/*21. Imprimir los primeros N números de la serie fibonacci, la serie de fibonacci se define
como:
F(0) = 1
F(1) = 1
F(n) = F(n-1) + F(n-2) */

package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	n := 20
	for i := 0; i < n; i++ {
		fmt.Println(fibonacci(i))
	}
}

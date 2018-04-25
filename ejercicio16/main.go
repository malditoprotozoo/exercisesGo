// 16.-Imprimir diez veces la serie de números del 1 al 10.
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(j)
		}
	}
}

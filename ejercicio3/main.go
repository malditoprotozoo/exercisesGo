// 3.-Hacer un programa que imprima los números pares entre 0 y 100 tomando a 0 como par.
package main

import "fmt"

func main() {
	for i := 0; i <= 100; i += 2 {
		fmt.Println(i)
	}
}

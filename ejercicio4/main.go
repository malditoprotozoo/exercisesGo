// 4. Modifique el ejercicio anterior para imprimir los impares entre 0 y 100 tomando a 0 como par
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}
}

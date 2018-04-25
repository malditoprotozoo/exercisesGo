// 6.-Hacer un programa que imprima los números impares hasta el 100 y que imprima cuantos impares hay.
package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
		sum++
	}
	fmt.Println(sum)
}

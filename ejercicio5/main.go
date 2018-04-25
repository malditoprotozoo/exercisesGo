// 5.-Hacer un programa que imprima la suma de los 100 primeros números.
package main

import "fmt"

func main() {
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

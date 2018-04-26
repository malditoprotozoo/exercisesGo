/*18.-Realizar un programa que reciba un número N y genere una tabla decreciente, por
ejemplo, si se ingresa un 5, entonces imprime:
12345
1234
123
12
1 */
package main

import "fmt"

func printNum(n int) {
	for i := 1; i <= n; i++ {
		if i != n {
			fmt.Print(i)
		} else {
			fmt.Printf("%d \n", i)
		}
	}
}
func main() {
	num := 9
	for i := num; i >= 0; i-- {
		printNum(i)
	}
}

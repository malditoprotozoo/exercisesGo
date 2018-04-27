/*24. Cree un programa que reciba un número n e imprima n "a" y n "b", es decir, si ingresa
el usuario un 5, imprime aaaaabbbbb, luego intente realizarlo sin usar ciclos (for, while o
do..while) */

package main

import "fmt"

func printLetters(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("a")
	}
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Print("b \n")
		} else {
			fmt.Print("b")
		}
	}
}

func main() {
	fmt.Println("Pick a number")
	var input int
	fmt.Scanln(&input)
	printLetters(input)
}

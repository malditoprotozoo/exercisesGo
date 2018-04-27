/*24. Cree un programa que reciba un número n e imprima n "a" y n "b", es decir, si ingresa
el usuario un 5, imprime aaaaabbbbb, luego intente realizarlo sin usar ciclos (for, while o
do..while) NOTA: Este es sin ciclos */

package main

import "fmt"

func printALetters(n int) {
	if n > 0 {
		printALetters(n - 1)
		fmt.Print("a")
	}
}

func printBLetters(n int) {
	if n > 0 {
		printBLetters(n - 1)
		fmt.Print("b")
	}
}

func main() {
	n := 9
	printALetters(n)
	printBLetters(n)
}

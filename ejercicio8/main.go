// 8.-Hacer un programa que solo nos permita introducir S o N, es decir, le pedirá al usuario una letra, hasta que sea S o N.

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Type S or N")
	var input string
	fmt.Scanln(&input)
	for {
		if strings.Compare("S", input) == 0 {
			fmt.Println("Yay!")
			break
		} else if strings.Compare("N", input) == 0 {
			fmt.Println("Yay!")
			break
		} else {
			fmt.Println("Only S or N allowed")
			fmt.Scanln(&input)
		}
	}
}

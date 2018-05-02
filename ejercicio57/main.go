/*57. Debe recibir una palabra e imprimir si tenía el teclado pegado o no, esto ocurre
cuando cada letra esta repetida inmediatamente, por ejemplo, el usuario quería escribir
“hola”, pero como su teclado estaba pegado escribió “hhoollaa”.*/

package main

import "fmt"

func checkKeyboard(str string) bool {
	res := false
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			res = true
		} else {
			res = false
		}
	}
	return res
}

func main() {
	fmt.Println("Write something")
	var input string
	fmt.Scanln(&input)
	fmt.Println(checkKeyboard(input))
}

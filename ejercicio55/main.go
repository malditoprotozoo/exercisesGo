/*55. Genere una clase (si corresponde) que contenga un método que reciba un arreglo de
palabras y verifique si el arreglo es palíndromo, esto es cuando al leer el arreglo palabra
por palabra desde el final al inicio se lee lo mismo que al leerlo palabra por palabra desde
el inicio al final. Por ejemplo:

{ “esto”, “es”, “un”, “ejemplo”, “ejemplo”, “un”, “es”, “esto”}
Ejemplo de lo que no es un arreglo palíndromo:
{“anita”, “lava”, “la”, “tina”} */

package main

import "fmt"

func checkPalindrome(phrase []string) bool {
	for i := 0; i < len(phrase)/2; i++ {
		if phrase[i] != phrase[len(phrase)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	example := []string{"esto", "es", "un", "ejemplo", "ejemplo", "un", "es", "esto"}
	anita := []string{"anita", "lava", "la", "tina"}
	fmt.Println(checkPalindrome(example))
	fmt.Println(checkPalindrome(anita))
}

/*56. Crear un programa que reciba una palabra y luego imprima si la palabra es un
palíndromo o no. */

package main

import "fmt"

func checkPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Insert a word")
	var input string
	fmt.Scanln(&input)
	if checkPalindrome(input) {
		fmt.Println("Is palindrome")
	} else {
		fmt.Println("Is not palindrome")
	}
}

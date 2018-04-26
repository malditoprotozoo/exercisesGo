// 10.-Introducir un numero por teclado. Que nos diga si es par o impar.
package main

import "fmt"

func even(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Enter a number to find odd or even")
	var input int
	fmt.Scanln(&input)
	if even(input) == true {
		fmt.Printf("%d is an even number \n", input)
	} else if even(input) == false {
		fmt.Printf("%d is an odd number \n", input)
	}
}

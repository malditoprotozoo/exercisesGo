/*17. Ingrese un número y luego realice un menú que tenga las siguientes opciones:
1. imprimir la suma de los números desde el 1 al N
2. imprimir los números desde el 1 al N
3. imprimir el factorial de N
4. salir */

package main

import "fmt"

func sumOfN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func numberSeq(n int) []int {
	sequence := []int{}
	for i := 1; i <= n; i++ {
		sequence = append(sequence, i)
	}
	return sequence
}

func factorial(num int) int {
	if num < 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	fmt.Print("Choose one option:\n 1. Print sum of numbers from 1 to n \n 2. Print numbers from 1 to n \n 3. Print N factorial \n 4. Finish \n")
	var input int
	fmt.Scanln(&input)
	if input == 1 {
		var n int
		fmt.Println("Choose a number")
		fmt.Scanln(&n)
		fmt.Printf("The sum of numbers from 1 to %d is %d \n", n, sumOfN(n))
	} else if input == 2 {
		var n int
		fmt.Println("Choose a number")
		fmt.Scanln(&n)
		fmt.Printf("The sequence from 1 to %d is %v \n", n, numberSeq(n))
	} else if input == 3 {
		var n int
		fmt.Println("Choose a number")
		fmt.Scanln(&n)
		fmt.Printf("The factorial of %d is %d \n", n, factorial(n))
	} else if input == 4 {
		fmt.Println("Byeee!")
	} else {
		fmt.Println("Input not valid")
	}
}

package main

import "fmt"

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		n := 1
		for i := 1; i <= num; i++ {
			n = n * i
		}
		return n
	}
}

func main() {
	fmt.Println(factorial(6))
}

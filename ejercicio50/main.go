/*50. El usuario podrá ingresar cuantos números desee mientras diga “continuar”, si dice
“basta” ya no podrá ingresar más, el programa deberá entregar cuantos de los números
que fue ingresando son primos (un número primo es aquel que sólo es divisible por 1 y
por si mismo). */

package main

import (
	"fmt"
	"math"
)

func checkPrime(n int) bool {
	maxCheck := int(math.Floor(float64(n / 2)))
	for i := 2; i <= maxCheck; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{3, 4, 5, 6, 7, 32, 45, 86, 12, 98, 72}
	result := []int{}
	for _, num := range arr {
		if checkPrime(num) {
			result = append(result, num)
		}
	}
	fmt.Println(result)
}

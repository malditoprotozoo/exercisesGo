/*19.-Comprobar si un número mayor o igual que la unidad es primo. Un número primo es
aquel que es divisible por 1 y por si mismo, pero por ningún otro número, por ejemplo el
11 o el 7, pero el 6 no, pues es divisible por 1, 2, 3 y 6 */

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
	if checkPrime(78) == false {
		fmt.Println("not prime")
	} else {
		fmt.Println("prime")
	}
}

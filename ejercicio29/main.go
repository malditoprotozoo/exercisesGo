/*29. Los conejos se reproducen rápidamente, a los 2 meses de nacidos ya pueden
reproducirse y lo hacen cada mes, cada pareja de conejos da a luz 2 conejos, estos
extraños conejos son inmortales. Cree un programa que retorne el número de conejos
existentes en una cantidad n de meses sabiendo que se inicia con una pareja de conejos
recién nacidos.*/

package main

import "fmt"

func rabbits(n int) int {
	total := 0
	for i := 2; i < n; i++ {
		total++
	}
	return total
}

func main() {
	months := 15
	total := 0
	for i := 0; i < months; i++ {
		if i%2 == 0 {
			total += rabbits(months)
		}
	}
	fmt.Println(total)
}

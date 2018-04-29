/*29. Los conejos se reproducen rápidamente, a los 2 meses de nacidos ya pueden
reproducirse y lo hacen cada mes, cada pareja de conejos da a luz 2 conejos, estos
extraños conejos son inmortales. Cree un programa que retorne el número de conejos
existentes en una cantidad n de meses sabiendo que se inicia con una pareja de conejos
recién nacidos. */

package main

import "fmt"

func rabbits(months int) int {
	if months == 1 || months == 2 {
		return 2
	}
	return rabbits(months-1) + rabbits(months-2)
}

func main() {
	months := 5
	fmt.Println(rabbits(months))
}

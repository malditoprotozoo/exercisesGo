/*45. Ahora ranas inmortales. Una familia de ranas africanas se compone inicialmente de
padre y madre, luego de 3 meses de vida se reproducen teniendo 100 hijos en cada
camada, cada par de hijos forma una familia, la que también se reproduce a los tres
meses. Las ranas se reproducen cada tres meses. Cada 6 meses, en los días posteriores a
la reproducción, se realiza una comilona caníbal donde cada una de las ranas
reproductivas se come a 20 parejas de ranas no reproductivas. Realice un programa que
dado un número de meses ingresado por el usuario y sabiendo que se inicia con una
pareja de ranas, entregue el número de ranas al finalizar el periodo de tiempo señalado.

CORREGIR FUNCIÓN RECURSIVA*/

package main

import "fmt"

/*func rabbits(months int) int {
	if months == 1 || months == 2 {
		return 2
	}
	return rabbits(months-1) + rabbits(months-2)
}

func main() {
	months := 5
	fmt.Println(rabbits(months))
}*/

func frogs(months int) int {
	bornedFrogs := 100
	if months == 1 || months == 2 {
		return 2
	}
	if months%6 == 0 {
		bornedFrogs -= 60
	}
	return (frogs(months-1) + frogs(months-2)) + bornedFrogs
}

func main() {
	fmt.Println(frogs(12))
}

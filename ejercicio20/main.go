/*20.-Ingresar N notas y calcular su promedio, si este es menor a 4,0 imprimir “reprobado”,
si es 4,0 imprimir “justo justo”, si es no mayor que 6,0, pero mayor que 4,0, imprimir
“buena nota” y si es mayor o igual a 6,0 imprimir “felicitaciones!” */
package main

import "fmt"

func averageGrade(arr ...float64) float64 {
	sum := 0.0
	for _, element := range arr {
		sum += element
	}
	return sum / float64(len(arr))
}

func main() {
	arr := []float64{2.0, 3.3, 4.0, 7.0, 6.3, 3.2, 6.8}
	finalGrade := averageGrade(arr...)
	switch {
	case finalGrade < 4.0:
		fmt.Println("Reprobado")
	case finalGrade == 4.0:
		fmt.Println("Justo justo")
	case finalGrade > 4.0 && finalGrade < 6.0:
		fmt.Println("Buena nota")
	case finalGrade >= 6.0:
		fmt.Println("¡Felicitaciones :D!")
	}
}

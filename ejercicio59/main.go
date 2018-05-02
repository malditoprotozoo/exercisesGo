/*59. Debe hacer un programa para Don Memento, el siempre escribe algo y luego se
acuerda de algo más que quería escribir, entonces usted hará un programa que sea capaz
de recibir un String una y otra vez hasta que Don Memento escriba “stop”, estos Strings
deberá concatenarlos uno tras otro e imprimirlo todo junto luego de ingresada la palabra
“stop”. Por ejemplo:
Quiero
Sacarme
Un
Siete
stop
Debe imprimir: QuieroSacarmeUnSiete */

package main

import "fmt"

func main() {
	fmt.Println("Write here")
	var input string
	notes := []string{}
	for input != "stop" {
		fmt.Scanln(&input)
		if input != "stop" {
			notes = append(notes, input)
		}
	}
	for i := 0; i < len(notes); i++ {
		fmt.Print(notes[i])
	}
}

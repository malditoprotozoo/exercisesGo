/*49. El método ruso de multiplicación de dos números A*B realizado “a mano” consiste en:
 Escribir los números (A y B) que se desea multiplicar en la parte superior de sendas
columnas.
 Dividir A entre 2, sucesivamente, ignorando el resto, hasta llegar a la unidad.
Escribir los resultados en la columna A.
 Multiplicar B por 2 tantas veces como veces se ha dividido A entre 2. Escribir los
resultados sucesivos en la columna B.
 Sumar todos los números de la columna B que estén al lado de un número impar
de la columna A. Éste es el resultado.
 Ejemplo: 27 × 82
A B Sumandos
27 82 82
13 164 164
6 328
3 656 656
1 1312 1312
Resultado 2214
Usted debe realizar un programa sin ciclos que reciba dos números a multiplicar y seguir el
método ruso para multiplicar estos dos números, se solicita imprimir el resultado, lo
importante es que realice los procedimientos solicitados.*/

package main

import "fmt"

func russianMethod(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	} else if b > 0 {
		return a + russianMethod(a, b-1)
	}
	return -russianMethod(a, -b)
}

func main() {
	fmt.Println(russianMethod(32, 10))
}

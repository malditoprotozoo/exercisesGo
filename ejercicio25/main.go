/*25. Cree un programa, que reciba un "reloj" es decir, la hora, los minutos y los segundos
para luego recibir una cantidad de segundos en que debe adelantarse el reloj, mostrando
la nueva hora. */

package main

import "fmt"

func main() {
	fmt.Println("Please insert the chosen hour, minutes and seconds. Press enter after each number")
	var hours, minutes, seconds int
	fmt.Scanln(&hours)
	fmt.Scanln(&minutes)
	fmt.Scanln(&seconds)
	if hours > 59 || minutes > 59 || seconds > 59 {
		fmt.Println("Please insert a valid format")
	}
	fmt.Println("Now insert how many seconds you want to add")
	var remainderSec int
	fmt.Scanln(&remainderSec)
	allSeconds := (hours * 3600) + (minutes * 60) + seconds + remainderSec
	hours = allSeconds / 3600
	allSeconds %= 3600
	minutes = allSeconds / 60
	allSeconds %= 60
	seconds = allSeconds
	fmt.Printf("The new time is %d:%d:%d \n", hours, minutes, seconds)
}

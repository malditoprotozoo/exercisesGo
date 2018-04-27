/*23. Resuelva el problema de las torres de hanoi */

package main

import "fmt"

func hanoi(disk int, source, spare, dest string) {
	if disk > 0 {
		hanoi(disk-1, source, dest, spare)
		fmt.Printf("Move disk %d from %s to %s \n", disk, source, dest)
		hanoi(disk-1, spare, source, dest)
	}
}

func main() {
	hanoi(3, "A", "B", "C")
}

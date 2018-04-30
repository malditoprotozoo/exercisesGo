/*22. Realice un programa que imprima todas las combinaciones de 1,2,3,4 */

package main

import "fmt"

func permutation(arr []int, leng int) {
	if leng == 1 {
		fmt.Println(arr)
	}
	for i := 0; i < leng; i++ {
		permutation(arr, leng-1)
		if leng%2 == 1 {
			arr[0], arr[leng-1] = arr[leng-1], arr[0]
		} else {
			arr[i], arr[leng-1] = arr[leng-1], arr[i]
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4}
	permutation(arr, len(arr))
}

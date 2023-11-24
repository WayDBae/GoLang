package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Введите значение N: ")
	fmt.Scan(&N)

	array := make([]int, N)
	array[0] = 1 // F1
	array[1] = 1 // F2

	for i := 2; i < N; i++ {
		array[i] = array[i-2] + array[i-1]
	}

	fmt.Println(array)
}
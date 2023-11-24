package main

import (
	"fmt"
)

func main() {
	var N, A, D int
	fmt.Print("Введите значение N: ")
	fmt.Scan(&N)
	fmt.Print("Введите первый член A: ")
	fmt.Scan(&A)
	fmt.Print("Введите разность D: ")
	fmt.Scan(&D)

	array := make([]int, N)

	for i := 0; i < N; i++ {
		array[i] = A + i*D
	}

	fmt.Println(array)
}
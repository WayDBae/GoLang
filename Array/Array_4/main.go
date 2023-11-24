package main

import (
	"fmt"
	"math"
)

func main() {
	var N, A, D int
	fmt.Print("Введите значение N: ")
	fmt.Scan(&N)
	fmt.Print("Введите первый член A: ")
	fmt.Scan(&A)
	fmt.Print("Введите разность D: ")
	fmt.Scan(&D)

	array := make([]float64, N)

	for i := 0; i < N; i++ {
		array[i] = float64(A) * math.Pow(float64(D), float64(i))
	}

	fmt.Println(array)
}

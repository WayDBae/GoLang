package main

import (
	"fmt"
	"math"
)

func main() {
	var A float64
	var N int

	fmt.Print("Введите вещественное число A: ")
	fmt.Scan(&A)

	fmt.Print("Введите целое число N (> 0): ")
	fmt.Scan(&N)

	result := math.Pow(A, float64(N))

	fmt.Printf("%.2f в степени %d = %.2f\n", A, N, result)
}
